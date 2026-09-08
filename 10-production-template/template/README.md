# Go Backend Template

Production-ready Go API skeleton with clean architecture.

## Quick start

```bash
cp .env.example .env
# start PostgreSQL (docker compose — Topic 13)
go run ./cmd/api/
curl http://localhost:8080/livez
```

## Structure

```
cmd/api/              entry point
domain/               entities + repository interfaces
application/          use cases + app errors + contracts
infra/                HTTP, DB, config, bootstrap, runtime
migrations/           SQL migrations
```

Read [ARCHITECTURE.md](./ARCHITECTURE.md) for layer rules and how to add endpoints.

## Learning path (Phase 10)

| Topic | Folder |
|-------|--------|
| 03 Domain | `domain/user/` |
| 04 Application | `application/usecase/` |
| 05 HTTP | `infra/httpserver/` |
| 06 Adapters | `infra/adapters/repository/` |
| 07+ Wiring, Wire, Docker | `cmd/`, `infra/bootstrap/` |

## Module

```go
module template
```

Import paths: `template/domain/user`, `template/application/usecase/register-user`, etc.
