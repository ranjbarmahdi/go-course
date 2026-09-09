# Go Backend Template

Production-ready Go API skeleton with clean architecture, Wire DI, PostgreSQL, Redis, and goose migrations.

## Quick start

```bash
cp .env.example .env
# Edit .env — set DATABASE_URL, DB_SCHEMA, REDIS_URL

# Create schema in PostgreSQL (once)
# CREATE SCHEMA routes;

go run ./cmd/migrate up
go run ./cmd/api
curl http://localhost:8080/livez
```

## What works today

- Health: `GET /livez`, `GET /readyz`
- Sample API: `POST /api/v2/samples`, `GET /api/v2/samples/{id}`
- Gateway auth via `x-user` header (see HANDOFF.md for curl examples)

## Structure

```
cmd/              entry points (api, migrate)
domain/           entities, value objects, repository interfaces
application/      use cases, app errors, contracts
infra/            HTTP, DB, config, Wire bootstrap, runtime
migrations/       goose SQL files
```

## Documentation

| File | Purpose |
|---|---|
| [ARCHITECTURE.md](./ARCHITECTURE.md) | Layer rules, where to put new code |
| [HANDOFF.md](./HANDOFF.md) | Full session context, decisions, next steps |

## Module

```go
module template
```

Import example: `template/domain/sample`, `template/application/usecase/sample/create-sample`

## Commands

```bash
go build ./...
go run ./cmd/api
go run ./cmd/migrate up
cd infra/bootstrap && go generate   # regenerate Wire
```
