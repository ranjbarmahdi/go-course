package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"
	"time"

	"template/infra/config"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func Open(ctx context.Context, databaseURL string) (*sql.DB, error) {
	db, err := sql.Open("pgx", databaseURL)
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
	db, err := Open(ctx, cfg.DatabaseURL)
	if err != nil {
		return nil, nil, err
	}
	log.Info("connected")
	cleanup := func() {
		if err := Close(db); err != nil {
			log.Warn("disconnect failed", "err", err)
			return
		}
		log.Info("disconnected")
	}
	return db, cleanup, nil
}
