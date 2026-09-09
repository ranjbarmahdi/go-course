package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"

	"template/infra/config"
	"template/infra/database/migrate"
	"template/infra/database/postgres"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("usage: migrate [up|down|status]")
		os.Exit(1)
	}

	cfg, err := config.Load()
	if err != nil {
		slog.Error("config", "err", err)
		os.Exit(1)
	}

	ctx := context.Background()

	db, err := postgres.Open(ctx, cfg.DatabaseURL, cfg.DBSchema)
	if err != nil {
		slog.Error("database", "err", err)
		os.Exit(1)
	}
	defer db.Close()

	switch os.Args[1] {
	case "up":
		err = migrate.Up(db)
	case "down":
		err = migrate.Down(db)
	case "status":
		err = migrate.Status(db)
	default:
		fmt.Println("unknown command:", os.Args[1])
		os.Exit(1)
	}

	if err != nil {
		slog.Error("migrate failed", "cmd", os.Args[1], "err", err)
		os.Exit(1)
	}

	slog.Info("migrate ok", "cmd", os.Args[1])
}
