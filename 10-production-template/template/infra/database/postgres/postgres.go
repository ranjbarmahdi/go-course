package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"net/url"
	"strings"
	"time"

	"template/infra/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

// dsnWithSearchPath puts search_path in the connection string so Postgres
// applies it to every connection the pool opens, including ones created later
// as the pool grows or recycles. A SET statement would only affect one.
func dsnWithSearchPath(databaseURL, schema string) (string, error) {
	u, err := url.Parse(databaseURL)
	if err != nil {
		return "", fmt.Errorf("parse database url: %w", err)
	}

	q := u.Query()
	q.Set("search_path", schema)
	u.RawQuery = q.Encode()

	return u.String(), nil
}

// verifySchema fails fast when the schema is missing. Creating it is an
// operator task, so the app only reports the problem instead of papering over
// it with an empty schema. pg_namespace is used rather than
// information_schema.schemata because the latter hides schemas the current
// user does not own.
func verifySchema(ctx context.Context, db *sql.DB, schema string) error {
	const query = `SELECT EXISTS (SELECT 1 FROM pg_catalog.pg_namespace WHERE nspname = $1)`

	var exists bool
	if err := db.QueryRowContext(ctx, query, schema).Scan(&exists); err != nil {
		return fmt.Errorf("check schema %q: %w", schema, err)
	}

	if !exists {
		return fmt.Errorf(
			"schema %q does not exist; create it first: CREATE SCHEMA %s",
			schema, schema,
		)
	}

	return nil
}

func Open(ctx context.Context, databaseURL, schema string) (*sql.DB, error) {
	if strings.TrimSpace(schema) == "" {
		return nil, fmt.Errorf("db schema is required")
	}

	dsn, err := dsnWithSearchPath(databaseURL, schema)
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, fmt.Errorf("open database: %w", err)
	}

	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(30 * time.Minute)

	pingCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if err := db.PingContext(pingCtx); err != nil {
		_ = db.Close()
		return nil, fmt.Errorf("ping database: %w", err)
	}

	if err := verifySchema(pingCtx, db, schema); err != nil {
		_ = db.Close()
		return nil, err
	}

	return db, nil
}

func Close(db *sql.DB) error {
	if db == nil {
		return nil
	}
	return db.Close()
}

var log = slog.Default().With("component", "postgres")

func NewPostgres(ctx context.Context, cfg *config.Config) (*sql.DB, func(), error) {
	db, err := Open(ctx, cfg.DatabaseURL, cfg.DBSchema)
	if err != nil {
		return nil, nil, err
	}
	log.Info("connected", "schema", cfg.DBSchema)
	cleanup := func() {
		if err := Close(db); err != nil {
			log.Warn("disconnect failed", "err", err)
			return
		}
		log.Info("disconnected")
	}
	return db, cleanup, nil
}
