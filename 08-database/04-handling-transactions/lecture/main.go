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
51 — HANDLING TRANSACTIONS IN GO
============================================================

Topics
------------------------------------------------------------
 1. Why Transactions in Use Cases
 2. Context-Propagation Pattern
 3. Layer Overview
 4. Domain Layer
 5. Application Layer — Ports
 6. Application Layer — Use Case with RunInTx
 7. Infrastructure — DBTX and getExecutor
 8. Infrastructure — Repositories
 9. Infrastructure — TransactionManager
10. Presentation — Handler
11. Composition Root
12. Complete Demo
13. Transaction Mental Model
============================================================

Pattern:

    use case calls txManager.RunInTx(ctx, func(ctx) {
        userRepo.Create(ctx, ...)
        profileRepo.Create(ctx, ...)
    })

    infrastructure puts *sql.Tx in context
    repositories read tx from context automatically

Application and domain NEVER import database/sql.
Application NEVER calls BeginTx or Commit.
*/

const dsn = "postgres://postgres:secret@localhost:5432/postgres?sslmode=disable"

// ============================================================
// LAYER: DOMAIN
// ============================================================

/*
Pure business entities and errors.

No context keys.
No SQL.
No transactions.
*/

type User struct {
	ID           string
	Email        string
	PasswordHash string
}

type Profile struct {
	ID     string
	UserID string
	Name   string
}

var (
	ErrInvalidInput       = errors.New("invalid input")
	ErrEmailAlreadyExists = errors.New("email already exists")
)

// ============================================================
// LAYER: APPLICATION — PORTS (INTERFACES)
// ============================================================

/*
Interfaces the application needs.

TransactionManager is the abstraction for RunInTx.

Use case depends on TransactionManager — NOT on *sql.Tx.
*/

type UserRepository interface {
	Create(ctx context.Context, user User) error
	FindByEmail(ctx context.Context, email string) (User, error)
}

type ProfileRepository interface {
	Create(ctx context.Context, profile Profile) error
}

type TransactionManager interface {
	RunInTx(ctx context.Context, fn func(ctx context.Context) error) error
}

// ============================================================
// LAYER: APPLICATION — USE CASE
// ============================================================

/*
Use case has injected repos + TransactionManager.

Inside RunInTx:

    return uc.txManager.RunInTx(ctx, func(ctx context.Context) error {
        uc.userRepo.Create(ctx, ...)
        uc.profileRepo.Create(ctx, ...)
    })

Repos stay the same struct — ctx carries the transaction.
*/

type CreateUserWithProfileUseCase struct {
	txManager   TransactionManager
	userRepo    UserRepository
	profileRepo ProfileRepository
}

func NewCreateUserWithProfileUseCase(
	txManager TransactionManager,
	userRepo UserRepository,
	profileRepo ProfileRepository,
) *CreateUserWithProfileUseCase {
	return &CreateUserWithProfileUseCase{
		txManager:   txManager,
		userRepo:    userRepo,
		profileRepo: profileRepo,
	}
}

type CreateUserWithProfileInput struct {
	Email    string
	Password string
	Name     string
}

func (uc *CreateUserWithProfileUseCase) Execute(
	ctx context.Context,
	input CreateUserWithProfileInput,
) error {
	if input.Email == "" || input.Password == "" || input.Name == "" {
		return ErrInvalidInput
	}

	/*
		RunInTx wraps both Create calls in ONE transaction.

		If profileRepo.Create fails → userRepo.Create is rolled back.
	*/
	return uc.txManager.RunInTx(ctx, func(ctx context.Context) error {
		_, err := uc.userRepo.FindByEmail(ctx, input.Email)

		if err == nil {
			return ErrEmailAlreadyExists
		}

		if !errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("find user: %w", err)
		}

		userID := fmt.Sprintf("user-%d", time.Now().UnixNano())

		user := User{
			ID:           userID,
			Email:        input.Email,
			PasswordHash: "hashed:" + input.Password,
		}

		err = uc.userRepo.Create(ctx, user)
		if err != nil {
			return err
		}

		profile := Profile{
			ID:     fmt.Sprintf("profile-%d", time.Now().UnixNano()),
			UserID: userID,
			Name:   input.Name,
		}

		return uc.profileRepo.Create(ctx, profile)
	})
}

/*
Single-table read — no RunInTx:

    profileRepo.FindByUserID(ctx, userID)

Repo receives ctx without tx → getExecutor uses *sql.DB.
*/

// ============================================================
// LAYER: INFRASTRUCTURE — DBTX + CONTEXT TX KEY
// ============================================================

/*
DBTX = interface both *sql.DB and *sql.Tx implement.

getExecutor reads *sql.Tx from context.

If tx in context  → use transaction
If no tx          → use connection pool


THIS is where database/sql lives — infrastructure only.
*/

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type txContextKey struct{}

func getExecutor(ctx context.Context, db *sql.DB) DBTX {
	tx, ok := ctx.Value(txContextKey{}).(*sql.Tx)

	if ok {
		return tx
	}

	return db
}

// ============================================================
// LAYER: INFRASTRUCTURE — TRANSACTION MANAGER
// ============================================================

/*
PostgresTransactionManager implements TransactionManager.

RunInTx:

    1. BeginTx
    2. ctx = context.WithValue(ctx, txKey, tx)
    3. fn(txCtx)
    4. Commit or Rollback


Use case never sees BeginTx / Commit.
*/

type PostgresTransactionManager struct {
	db *sql.DB
}

func NewPostgresTransactionManager(db *sql.DB) *PostgresTransactionManager {
	return &PostgresTransactionManager{db: db}
}

func (tm *PostgresTransactionManager) RunInTx(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer tx.Rollback()

	txCtx := context.WithValue(ctx, txContextKey{}, tx)

	err = fn(txCtx)
	if err != nil {
		return err
	}

	err = tx.Commit()
	if err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

// ============================================================
// LAYER: INFRASTRUCTURE — USER REPOSITORY
// ============================================================

/*
Every method:

    exec := getExecutor(ctx, r.db)
    exec.QueryRowContext(ctx, ...)


Same code works inside and outside RunInTx.
*/

type PostgresUserRepository struct {
	db *sql.DB
}

func NewPostgresUserRepository(db *sql.DB) *PostgresUserRepository {
	return &PostgresUserRepository{db: db}
}

func (r *PostgresUserRepository) Create(ctx context.Context, user User) error {
	exec := getExecutor(ctx, r.db)

	_, err := exec.ExecContext(ctx, `
		INSERT INTO users (id, email, password_hash)
		VALUES ($1, $2, $3)
	`, user.ID, user.Email, user.PasswordHash)

	if err != nil {
		return fmt.Errorf("insert user: %w", err)
	}

	return nil
}

func (r *PostgresUserRepository) FindByEmail(ctx context.Context, email string) (User, error) {
	exec := getExecutor(ctx, r.db)

	var user User

	err := exec.QueryRowContext(ctx, `
		SELECT id, email, password_hash
		FROM users
		WHERE email = $1
	`, email).Scan(&user.ID, &user.Email, &user.PasswordHash)

	if err != nil {
		return User{}, err
	}

	return user, nil
}

// ============================================================
// LAYER: INFRASTRUCTURE — PROFILE REPOSITORY
// ============================================================

type PostgresProfileRepository struct {
	db *sql.DB
}

func NewPostgresProfileRepository(db *sql.DB) *PostgresProfileRepository {
	return &PostgresProfileRepository{db: db}
}

func (r *PostgresProfileRepository) Create(ctx context.Context, profile Profile) error {
	exec := getExecutor(ctx, r.db)

	_, err := exec.ExecContext(ctx, `
		INSERT INTO profiles (id, user_id, name)
		VALUES ($1, $2, $3)
	`, profile.ID, profile.UserID, profile.Name)

	if err != nil {
		return fmt.Errorf("insert profile: %w", err)
	}

	return nil
}

func (r *PostgresProfileRepository) FindByUserID(ctx context.Context, userID string) (Profile, error) {
	exec := getExecutor(ctx, r.db)

	var profile Profile

	err := exec.QueryRowContext(ctx, `
		SELECT id, user_id, name
		FROM profiles
		WHERE user_id = $1
	`, userID).Scan(&profile.ID, &profile.UserID, &profile.Name)

	if err != nil {
		return Profile{}, err
	}

	return profile, nil
}

// ============================================================
// LAYER: PRESENTATION — HANDLER (SIMULATED)
// ============================================================

/*
HTTP handler calls use case.

Handler does not know about transactions.
*/

func createUserWithProfileHandler(
	uc *CreateUserWithProfileUseCase,
) func(ctx context.Context, input CreateUserWithProfileInput) error {
	return func(ctx context.Context, input CreateUserWithProfileInput) error {
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

		fmt.Println("HTTP 201: user + profile created in one transaction")
		return nil
	}
}

// ============================================================
// LAYER: COMPOSITION ROOT — main
// ============================================================

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

func ensureSchema(db *sql.DB) error {
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id            TEXT PRIMARY KEY,
			email         TEXT NOT NULL UNIQUE,
			password_hash TEXT NOT NULL,
			created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);

		CREATE TABLE IF NOT EXISTS profiles (
			id       TEXT PRIMARY KEY,
			user_id  TEXT NOT NULL UNIQUE REFERENCES users(id),
			name     TEXT NOT NULL,
			created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
		);
	`)

	return err
}

func main() {
	db, err := openDB(dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = ensureSchema(db)
	if err != nil {
		log.Fatal(err)
	}

	/*
		Composition root — wire concrete implementations
	*/
	txManager := NewPostgresTransactionManager(db)
	userRepo := NewPostgresUserRepository(db)
	profileRepo := NewPostgresProfileRepository(db)

	createUserUC := NewCreateUserWithProfileUseCase(
		txManager,
		userRepo,
		profileRepo,
	)

	handler := createUserWithProfileHandler(createUserUC)

	ctx := context.Background()

	input := CreateUserWithProfileInput{
		Email:    "mahdi@example.com",
		Password: "secret123",
		Name:     "Mahdi",
	}

	err = handler(ctx, input)
	if err != nil {
		log.Fatal(err)
	}

	/*
		Verify both rows exist
	*/
	var userCount, profileCount int

	_ = db.QueryRow(`SELECT COUNT(*) FROM users WHERE email = $1`, input.Email).Scan(&userCount)
	_ = db.QueryRow(`SELECT COUNT(*) FROM profiles p JOIN users u ON u.id = p.user_id WHERE u.email = $1`, input.Email).Scan(&profileCount)

	fmt.Println()
	fmt.Printf("Users: %d  Profiles: %d\n", userCount, profileCount)
}

/*
============================================================
REQUEST FLOW
============================================================

    Handler
       ↓
    CreateUserWithProfileUseCase.Execute(ctx, input)
       ↓
    txManager.RunInTx(ctx, fn)
       ↓
    BeginTx → ctx = WithValue(tx)
       ↓
    userRepo.FindByEmail(txCtx)     ─┐
    userRepo.Create(txCtx)          ├─ same *sql.Tx via context
    profileRepo.Create(txCtx)      ─┘
       ↓
    Commit (or Rollback on error)


============================================================
TRANSACTION MENTAL MODEL
============================================================

Question 1:

    Who starts the transaction?

Answer:

    TransactionManager (infrastructure)


Question 2:

    How do repos join the transaction?

Answer:

    getExecutor reads *sql.Tx from context


Question 3:

    What does use case call?

Answer:

    txManager.RunInTx(ctx, func(ctx) { repo methods })


Question 4:

    Single table operation?

Answer:

    Skip RunInTx — call repo directly


Question 5:

    Does use case import database/sql?

Answer:

    No — only infrastructure does


============================================================
WHEN TO USE RunInTx
============================================================

Use RunInTx:

    user + profile create
    order + order_items create
    transfer money between accounts


Skip RunInTx:

    find user by email
    list products
    single INSERT into one table


============================================================
END OF 51 — HANDLING TRANSACTIONS IN GO
============================================================
*/
