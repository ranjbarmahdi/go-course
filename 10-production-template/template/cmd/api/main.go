package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"template/infra/bootstrap"
	"template/infra/config"
	"template/infra/database/migrate"
	"template/infra/database/postgres"
)

func main() {
	if err := run(); err != nil {
		slog.Error("application stopped with error", "error", err)
		os.Exit(1)
	}
}

func run() error {
	cfg, err := config.Load()
	if err != nil {
		return err
	}

	level := config.ParseSlogLevel(cfg.LogLevel)
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	slog.SetDefault(logger)

	slog.Info("logger configured",
		slog.String("configured_level", cfg.LogLevel),
		slog.String("env", cfg.Env),
		slog.Int("port", cfg.Port),
	)

	ctx, stop := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)
	defer stop()

	if cfg.AutoMigrate {
		db, err := postgres.Open(ctx, cfg.DatabaseURL, cfg.DBSchema)
		if err != nil {
			return err
		}
		if err := migrate.Up(db); err != nil {
			_ = db.Close()
			return fmt.Errorf("auto migrate: %w", err)
		}
		_ = db.Close()
		slog.Info("migrations applied", "mode", "auto", "schema", cfg.DBSchema)
	}

	app, cleanup, err := bootstrap.Wire(ctx, cfg)
	if err != nil {
		return err
	}
	defer cleanup()

	slog.Info("application started")

	if err := app.Run(ctx); err != nil {
		return err
	}

	slog.Info("application stopped gracefully")
	return nil
}
