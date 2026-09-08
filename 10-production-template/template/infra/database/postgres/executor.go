package postgres

import (
	"context"
	"database/sql"
)

type DBTX interface {
	ExecContext(ctx context.Context, query string, args ...any) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...any) *sql.Row
	QueryContext(ctx context.Context, query string, args ...any) (*sql.Rows, error)
}

type txContextKey struct{}

func Executor(ctx context.Context, db *sql.DB) DBTX {
	if tx, ok := ctx.Value(txContextKey{}).(*sql.Tx); ok {
		return tx
	}
	return db
}
