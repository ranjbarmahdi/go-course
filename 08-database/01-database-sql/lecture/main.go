package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

/*
============================================================
48 — DATABASE/SQL
============================================================

Topics
------------------------------------------------------------
 1. Why Databases
 2. database/sql Overview
 3. pgx Driver
 4. Open, Ping, Close
 5. Connection Pool
 6. CREATE TABLE
 7. INSERT — Exec
 8. SELECT One Row — QueryRow
 9. sql.ErrNoRows
10. SELECT Many Rows — Query
11. Context With Queries
12. Auth + Database
13. Complete Demo
14. database/sql Mental Model
============================================================
*/

// ============================================================
// 1. WHY DATABASES
// ============================================================

/*
In 47-authentication you used an in-memory store:

    map[string]User


Problems:

    Data is lost when the server restarts
    Not shared across multiple server instances
    No safe concurrent access at scale


Solution:

    PostgreSQL (persistent database)


Flow:

    HTTP Handler
         ↓
    Service
         ↓
    Repository
         ↓
    database/sql
         ↓
    pgx driver
         ↓
    PostgreSQL
*/

// ============================================================
// 2. DATABASE/SQL OVERVIEW
// ============================================================

/*
Package: database/sql

It is NOT a database driver.
It is a standard interface for SQL databases.


Core types:

    *sql.DB     → connection pool
    *sql.Row     → one row result
    *sql.Rows    → many rows result
    *sql.Tx      → transaction (topic 49)


Core methods:

    db.Exec(...)       → INSERT, UPDATE, DELETE, CREATE
    db.QueryRow(...)   → SELECT one row
    db.Query(...)      → SELECT many rows
    db.Ping()          → verify connection


Architecture:

    Your Go code
         ↓
    database/sql
         ↓
    driver (pgx for PostgreSQL)
         ↓
    PostgreSQL server
*/

// ============================================================
// 3. PGX DRIVER
// ============================================================

/*
PostgreSQL driver for Go:

    github.com/jackc/pgx/v5/stdlib


Register with blank import:

    _ "github.com/jackc/pgx/v5/stdlib"


Then open with driver name "pgx":

    sql.Open("pgx", dsn)


DSN = Data Source Name (connection string):

    postgres://USER:PASSWORD@HOST:PORT/DATABASE?sslmode=disable


Docker example (this lecture):

    postgres://postgres:secret@localhost:5432/postgres?sslmode=disable
*/

const dsn = "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"

// ============================================================
// 4. OPEN, PING, CLOSE
// ============================================================

/*
Open a pool:

    db, err := sql.Open("pgx", dsn)


Important:

    sql.Open does NOT connect immediately.
    It prepares a connection pool.


Verify with Ping:

    err = db.Ping()


Always close when done:

    defer db.Close()
*/

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open db: %w", err)
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}

	return db, nil
}

// ============================================================
// 5. CONNECTION POOL
// ============================================================

/*
*sql.DB is a pool of connections, NOT a single connection.


Why pool?

    Opening a connection is expensive.
    Many HTTP requests reuse connections from the pool.


Common settings:

    SetMaxOpenConns(n)       → max open connections
    SetMaxIdleConns(n)       → max idle connections in pool
    SetConnMaxLifetime(d)    → recycle old connections
    SetConnMaxIdleTime(d)    → close idle connections after duration
*/

func configurePool(db *sql.DB) {
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)
	db.SetConnMaxIdleTime(1 * time.Minute)
}

// ============================================================
// 6. CREATE TABLE
// ============================================================

/*
Use Exec for DDL (schema) statements.


PostgreSQL types:

    TEXT         → string
    INTEGER      → int
    BOOLEAN      → bool
    TIMESTAMPTZ  → time.Time (with timezone)


Example users table for auth:

    id
    email
    password_hash
    created_at
*/

func createUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id            TEXT PRIMARY KEY,
			email         TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)

	if err != nil {
		return fmt.Errorf("create users table: %w", err)
	}

	return nil
}

// ============================================================
// 7. INSERT — EXEC
// ============================================================

/*
Exec runs a query that does not return rows.

Used for:

    INSERT
    UPDATE
    DELETE
    CREATE TABLE


PostgreSQL placeholders:

    $1, $2, $3   (NOT ? like MySQL)


Example:

    INSERT INTO users (id, email, password_hash)
    VALUES ($1, $2, $3)
*/

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

func insertUser(db *sql.DB, user User) error {
	_, err := db.Exec(
		`INSERT INTO users (id, email, password_hash)
		 VALUES ($1, $2, $3)`,
		user.ID,
		user.Email,
		user.PasswordHash,
	)

	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil
}

// ============================================================
// 8. SELECT ONE ROW — QUERYROW
// ============================================================

/*
QueryRow returns exactly one row.

Use Scan to read columns into Go variables.


Example:

    err := db.QueryRow(
        "SELECT id, email FROM users WHERE email = $1",
        email,
    ).Scan(&id, &email)
*/

func findUserByEmail(db *sql.DB, email string) (User, error) {
	var user User

	err := db.QueryRow(
		`SELECT id, email, password_hash, created_at
		 FROM users
		 WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// ============================================================
// 9. SQL.ERRNOROWS
// ============================================================

/*
When QueryRow finds no row, Scan returns:

    sql.ErrNoRows


Always handle it separately from real database errors.


Not found:

    err == sql.ErrNoRows


Real DB error:

    err != nil && err != sql.ErrNoRows


Login example:

    user, err := findUserByEmail(db, email)

    if err == sql.ErrNoRows {
        // unknown email → 401
    }

    if err != nil {
        // database failure → 500
    }

    // verify password
*/

// ============================================================
// 10. SELECT MANY ROWS — QUERY
// ============================================================

/*
Query returns multiple rows.


Rules:

    1. defer rows.Close()
    2. loop rows.Next()
    3. Scan inside loop
    4. check rows.Err() after loop
*/

func listUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query(
		`SELECT id, email, password_hash, created_at
		 FROM users
		 ORDER BY created_at`,
	)

	if err != nil {
		return nil, fmt.Errorf("list users: %w", err)
	}

	defer rows.Close()

	users := make([]User, 0)

	for rows.Next() {
		var user User

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.PasswordHash,
			&user.CreatedAt,
		)

		if err != nil {
			return nil, fmt.Errorf("scan user: %w", err)
		}

		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return users, nil
}

// ============================================================
// 11. CONTEXT WITH QUERIES
// ============================================================

/*
Use context for timeouts and cancellation.


Methods:

    ExecContext
    QueryContext
    QueryRowContext


Example:

    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    err := db.QueryRowContext(ctx, query, args...).Scan(...)
*/

func findUserByEmailWithContext(
	ctx context.Context,
	db *sql.DB,
	email string,
) (User, error) {
	var user User

	err := db.QueryRowContext(
		ctx,
		`SELECT id, email, password_hash, created_at
		 FROM users
		 WHERE email = $1`,
		email,
	).Scan(
		&user.ID,
		&user.Email,
		&user.PasswordHash,
		&user.CreatedAt,
	)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// ============================================================
// 12. AUTH + DATABASE
// ============================================================

/*
Replace in-memory auth store with PostgreSQL.


REGISTER:

    hash password
         ↓
    INSERT INTO users (...)


LOGIN:

    SELECT user WHERE email = $1
         ↓
    if sql.ErrNoRows → 401
         ↓
    verify password with bcrypt
         ↓
    return JWT


Same handler logic.

Only persistence layer changes.
*/

// ============================================================
// 13. COMPLETE DEMO
// ============================================================

func main() {
	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	configurePool(db)

	fmt.Println("Connected to PostgreSQL")

	err = createUsersTable(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Table users ready")

	/*
		Insert demo user (ignore duplicate email error for re-runs)
	*/
	demoUser := User{
		ID:           "user-1",
		Email:        "mahdi@example.com",
		PasswordHash: "$2a$10$hashedpasswordhere",
	}

	err = insertUser(db, demoUser)
	if err != nil {
		fmt.Println("Insert skipped (user may already exist):", err)
	} else {
		fmt.Println("User inserted:", demoUser.Email)
	}

	/*
		SELECT one row
	*/
	user, err := findUserByEmail(db, "mahdi@example.com")
	if err == sql.ErrNoRows {
		fmt.Println("User not found")
		return
	}
	if err != nil {
		log.Fatal("find user:", err)
	}

	fmt.Println("Found user:", user.ID, user.Email)

	/*
		SELECT many rows
	*/
	users, err := listUsers(db)
	if err != nil {
		log.Fatal("list users:", err)
	}

	fmt.Println()
	fmt.Println("All users:")

	for _, u := range users {
		fmt.Printf("  %s  %s  %s\n", u.ID, u.Email, u.CreatedAt.Format(time.RFC3339))
	}

	/*
		Context query
	*/
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	user, err = findUserByEmailWithContext(ctx, db, user.Email)
	if err != nil {
		log.Fatal("context query:", err)
	}

	fmt.Println()
	fmt.Println("Context query OK:", user.Email)
}

/*
============================================================
DATABASE/SQL MENTAL MODEL
============================================================

Question 1:

    What is *sql.DB?

Answer:

    A connection pool, not one connection


Question 2:

    When do I use Exec vs QueryRow vs Query?

Answer:

    Exec       → no rows returned (INSERT, UPDATE, DELETE)
    QueryRow   → exactly one row (SELECT by unique key)
    Query      → many rows (SELECT list)


Question 3:

    How do I detect "not found"?

Answer:

    err == sql.ErrNoRows


Question 4:

    PostgreSQL placeholders?

Answer:

    $1, $2, $3


Question 5:

    How does auth connect to DB?

Answer:

    Register → INSERT
    Login    → SELECT + bcrypt verify


============================================================
DOCKER POSTGRES (THIS LECTURE)
============================================================

Start container:

    docker run -d --name postgres \
      -e POSTGRES_PASSWORD=secret \
      -p 5432:5432 \
      postgres


Connect with psql:

    docker exec -it postgres psql -U postgres


Run lecture:

    go run main.go


============================================================
END OF 48 — DATABASE/SQL
============================================================
*/
