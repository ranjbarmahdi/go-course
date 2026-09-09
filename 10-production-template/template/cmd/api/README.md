# cmd/api

Application entry point.

## What main.go does

1. Load config from environment (`.env` via godotenv)
2. Configure `slog` from `LOG_LEVEL`
3. Listen for SIGINT/SIGTERM
4. Optionally auto-migrate when `AUTO_MIGRATE=true`
5. Call `bootstrap.Wire(ctx, cfg)` to build the app
6. Run `app.Run(ctx)` until shutdown

## Rules

- No business logic here
- No manual wiring — Wire in `infra/bootstrap/` owns the graph
- Keep `main` thin

## Run

```bash
go run ./cmd/api
```
