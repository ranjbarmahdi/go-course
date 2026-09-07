package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

/*
============================================================
50 — REPOSITORY PATTERN
============================================================

Topics
------------------------------------------------------------
 1. Why Repository Pattern
 2. Layered Architecture
 3. Domain Layer
 4. Application Layer — Ports
 5. Application Layer — Use Cases
 6. Infrastructure Layer — PostgreSQL Repository
 7. Presentation Layer — HTTP Handler
 8. Composition Root — Wiring
 9. Error Mapping
10. Complete Demo
11. Repository Mental Model
============================================================

For transactions across multiple repositories, see:

    08-database/04-handling-transactions/lecture/main.go
*/

// ============================================================
// 1. WHY REPOSITORY PATTERN
// ============================================================

/*
In 48 and 49 you wrote SQL directly in main().

Problems:

    SQL mixed with HTTP logic
    Hard to test without a real database
    Hard to swap PostgreSQL for another storage


Solution — Repository Pattern:

    Handler   → HTTP
    Use Case  → business logic
    Repository → SQL only


The service/use case depends on an INTERFACE.

PostgreSQL is one implementation.
A mock is another (for tests).
*/

// ============================================================
// 2. LAYERED ARCHITECTURE
// ============================================================

/*
                    HTTP Request
                         ↓
              ┌──────────────────────┐
              │  Presentation        │  handlers, JSON, status codes
              └──────────┬───────────┘
                         ↓
              ┌──────────────────────┐
              │  Application         │  use cases, port interfaces
              └──────────┬───────────┘
                         ↓
              ┌──────────────────────┐
              │  Domain              │  entities, domain errors
              └──────────┬───────────┘
                         ↓
              ┌──────────────────────┐
              │  Infrastructure      │  PostgreSQL, bcrypt impl
              └──────────────────────┘


Dependency direction:

    outer layers depend on inner abstractions

    Handler  → UseCase  → Repository interface
                              ↑
                    PostgresUserRepository implements it
*/

const dsn = "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"

// ============================================================
// 3. DOMAIN LAYER
// ============================================================

/*
Domain = pure business concepts.

No HTTP.
No SQL.
No JSON tags required (optional for DTOs in presentation).


Domain errors are also defined here.
*/

type User struct {
	ID           string
	Email        string
	PasswordHash string
	CreatedAt    time.Time
}

var (
	ErrUserNotFound       = errors.New("user not found")
	ErrEmailAlreadyExists = errors.New("email already exists")
	ErrInvalidInput       = errors.New("invalid input")
)

// ============================================================
// 4. APPLICATION LAYER — PORTS (INTERFACES)
// ============================================================

/*
Ports = interfaces the application needs.

Defined in application layer.

Infrastructure implements them.


The application does NOT import database/sql.
*/

type UserRepository interface {
	Create(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
	FindByID(ctx context.Context, id string) (User, error)
}

// ============================================================
// 5. APPLICATION LAYER — USE CASES
// ============================================================

/*
Use Case = one business operation.

Examples:

    RegisterUser
    LoginUser
    GetUserProfile


Use case rules:

    - no SQL strings
    - no ResponseWriter
    - depends on repository interface
    - returns domain errors
*/

type RegisterUserUseCase struct {
	users UserRepository
}

func NewRegisterUserUseCase(users UserRepository) *RegisterUserUseCase {
	return &RegisterUserUseCase{users: users}
}

type RegisterUserInput struct {
	Email    string
	Password string
}

func (uc *RegisterUserUseCase) Execute(ctx context.Context, input RegisterUserInput) error {
	if input.Email == "" || input.Password == "" {
		return ErrInvalidInput
	}

	if len(input.Password) < 8 {
		return ErrInvalidInput
	}

	_, err := uc.users.FindByEmail(ctx, input.Email)

	if err == nil {
		return ErrEmailAlreadyExists
	}

	if !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("find user: %w", err)
	}

	hash, err := hashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	user := User{
		ID:           fmt.Sprintf("user-%d", time.Now().UnixNano()),
		Email:        input.Email,
		PasswordHash: hash,
	}

	return uc.users.Create(ctx, user)
}

type GetUserUseCase struct {
	users UserRepository
}

func NewGetUserUseCase(users UserRepository) *GetUserUseCase {
	return &GetUserUseCase{users: users}
}

func (uc *GetUserUseCase) Execute(ctx context.Context, email string) (User, error) {
	user, err := uc.users.FindByEmail(ctx, email)

	if errors.Is(err, sql.ErrNoRows) {
		return User{}, ErrUserNotFound
	}

	if err != nil {
		return User{}, fmt.Errorf("find user: %w", err)
	}

	return user, nil
}

// ============================================================
// 6. INFRASTRUCTURE LAYER — POSTGRESQL REPOSITORY
// ============================================================

/*
Infrastructure implements port interfaces.

This is the ONLY place that knows:

    SQL queries
    PostgreSQL types
    database/sql
*/

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user User) error {
	_, err := r.db.ExecContext(ctx, `
		INSERT INTO users (id, email, password_hash)
		VALUES ($1, $2, $3)
	`, user.ID, user.Email, user.PasswordHash)

	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (User, error) {
	var user User

	err := r.db.QueryRowContext(ctx, `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE email = $1
	`, email).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (r *PostgresUserRepository) FindByID(ctx context.Context, id string) (User, error) {
	var user User

	err := r.db.QueryRowContext(ctx, `
		SELECT id, email, password_hash, created_at
		FROM users
		WHERE id = $1
	`, id).Scan(&user.ID, &user.Email, &user.PasswordHash, &user.CreatedAt)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

func hashPassword(password string) (string, error) {
	// simplified for lecture — use bcrypt in production (topic 47)
	return "hashed:" + password, nil
}

// ============================================================
// 7. PRESENTATION LAYER — HTTP HANDLER
// ============================================================

/*
Handler = HTTP adapter.

Responsibilities:

    decode JSON
    call use case
    map errors to HTTP status
    encode JSON response


Handler does NOT contain SQL or business rules.
*/

func registerHandler(uc *RegisterUserUseCase) httpHandler {
	return func(ctx context.Context) error {
		input := RegisterUserInput{
			Email:    "newuser@example.com",
			Password: "secret123",
		}

		err := uc.Execute(ctx, input)

		if errors.Is(err, ErrInvalidInput) {
			return fmt.Errorf("HTTP 400: %w", err)
		}

		if errors.Is(err, ErrEmailAlreadyExists) {
			return fmt.Errorf("HTTP 409: %w", err)
		}

		if err != nil {
			return fmt.Errorf("HTTP 500: %w", err)
		}

		fmt.Println("HTTP 201: user registered")
		return nil
	}
}

type httpHandler func(ctx context.Context) error

// ============================================================
// 8. COMPOSITION ROOT — WIRING (main)
// ============================================================

/*
Composition root = where dependencies are assembled.

This is the only place that knows concrete types:

    PostgresUserRepository
    RegisterUserUseCase
    *sql.DB
*/

func openDB(dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

func ensureUsersTable(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id            TEXT PRIMARY KEY,
			email         TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
		)
	`)

	return err
}

func main() {
	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = ensureUsersTable(db)
	if err != nil {
		log.Fatal(err)
	}

	/*
		Wiring — concrete → interface
	*/
	userRepo := NewPostgresUserRepository(db)
	registerUC := NewRegisterUserUseCase(userRepo)
	getUserUC := NewGetUserUseCase(userRepo)

	ctx := context.Background()

	/*
		Simulate HTTP register
	*/
	err = registerHandler(registerUC)(ctx)
	if err != nil {
		log.Println("register:", err)
	}

	/*
		Simulate get user
	*/
	user, err := getUserUC.Execute(ctx, "newuser@example.com")
	if err != nil {
		log.Fatal("get user:", err)
	}

	fmt.Println()
	fmt.Println("Found user:", user.ID, user.Email)
}

/*
============================================================
9. ERROR MAPPING
============================================================

Domain / use case          HTTP
-------------------------  -----
ErrInvalidInput            400 Bad Request
ErrEmailAlreadyExists      409 Conflict
ErrUserNotFound            404 Not Found
sql.ErrNoRows (in repo)    map to ErrUserNotFound in use case
other errors               500 Internal Server Error


Mapping happens in the HANDLER, not in the repository.
*/

/*
============================================================
10. REPOSITORY MENTAL MODEL
============================================================

Question 1:

    Where does SQL go?

Answer:

    Repository (infrastructure)


Question 2:

    Where does business logic go?

Answer:

    Use case (application)


Question 3:

    Where are interfaces defined?

Answer:

    Application layer (ports)


Question 4:

    Where are concrete types wired?

Answer:

    main() — composition root


Question 5:

    Single table or multiple tables atomically?

Answer:

    Single table     → use case + one repository
    Multiple tables  → see 04-handling-transactions


============================================================
END OF 50 — REPOSITORY PATTERN
============================================================
*/
