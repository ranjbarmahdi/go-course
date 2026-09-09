package migrate

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/pressly/goose/v3"
)

const Dir = "migrations"

// hasMigrations reports whether Dir holds any .sql files. goose rejects an
// empty directory, but a template that has not written its first migration yet
// is a legitimate state, so callers treat this as "nothing to do".
func hasMigrations() (bool, error) {
	entries, err := os.ReadDir(Dir)
	if os.IsNotExist(err) {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("read migrations dir: %w", err)
	}

	for _, entry := range entries {
		if !entry.IsDir() && strings.EqualFold(filepath.Ext(entry.Name()), ".sql") {
			return true, nil
		}
	}

	return false, nil
}

func Up(db *sql.DB) error {
	ok, err := hasMigrations()
	if err != nil {
		return err
	}
	if !ok {
		slog.Info("no migrations to apply", "dir", Dir)
		return nil
	}

	return goose.Up(db, Dir)
}

func Down(db *sql.DB) error {
	ok, err := hasMigrations()
	if err != nil {
		return err
	}
	if !ok {
		slog.Info("no migrations to roll back", "dir", Dir)
		return nil
	}

	return goose.Down(db, Dir)
}

func Status(db *sql.DB) error {
	ok, err := hasMigrations()
	if err != nil {
		return err
	}
	if !ok {
		slog.Info("no migrations found", "dir", Dir)
		return nil
	}

	return goose.Status(db, Dir)
}
