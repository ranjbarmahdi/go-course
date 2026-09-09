# infra/database

Database connections, migrations, and health indicators.

## Layout

```
database/
├── migrate/migrate.go    goose wrapper (Up, Down, Status)
├── postgres/
│   ├── postgres.go       Open, NewPostgres, schema search_path
│   ├── executor.go       Executor(ctx, db) — tx-aware queries
│   ├── transaction_manager.go   RunInTx
│   └── indicator.go      readiness probe
└── redis/
    ├── redis.go          direct + sentinel client
    └── indicator.go      readiness probe
```

## PostgreSQL

- Pool opened at startup via Wire
- `DB_SCHEMA` sets `search_path` on the connection
- Operator must create the schema: `CREATE SCHEMA <name>;`
- Repositories call `postgres.Executor(ctx, db)` so they work inside and outside transactions

## Redis

- `development` / `self-host` → direct connection via `REDIS_URL`
- Other environments → Sentinel failover client

## Migrations

Use `cmd/migrate` — not the API — for manual migration runs:

```bash
go run ./cmd/migrate status
go run ./cmd/migrate up
go run ./cmd/migrate down
```

Or set `AUTO_MIGRATE=true` in `.env` to migrate on API startup.

## Health

Both databases expose `Indicator` implementations used by `GET /readyz`.
