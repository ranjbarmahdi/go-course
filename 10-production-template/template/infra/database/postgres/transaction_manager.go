package postgres

import (
	"context"
	"database/sql"
	"fmt"

	"template/application/contracts"
)

type TransactionManager struct {
	db *sql.DB
}

func NewTransactionManager(db *sql.DB) *TransactionManager {
	return &TransactionManager{db: db}
}

func (tm *TransactionManager) RunInTx(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	tx, err := tm.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("begin tx: %w", err)
	}

	defer tx.Rollback()

	txCtx := context.WithValue(ctx, txContextKey{}, tx)

	if err := fn(txCtx); err != nil {
		return err
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("commit tx: %w", err)
	}

	return nil
}

func NewRunInTx(db *sql.DB) contracts.RunInTx[error] {
	tm := NewTransactionManager(db)
	return tm.RunInTx
}
