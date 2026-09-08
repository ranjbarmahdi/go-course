# Architecture

Production Go backend template — layer rules and "where do I put X?"

## Four layers

```
HTTP Request
     │
     ▼
infra/httpserver/          JSON, routes, middleware
     │
     ▼
application/usecase/       business flow
     │
     ▼
domain/                    entities, repository interfaces
     ▲
     │ implements
infra/adapters/            PostgreSQL, Kafka, JWT
```

`cmd/api/main.go` wires everything (composition root).

**Golden rule:** dependencies point inward. Domain never imports infra.

## Layer jobs

| Layer | Job |
|-------|-----|
| `domain/` | Entities, domain errors, repository **interfaces**. No HTTP, no SQL. |
| `application/` | Use cases, business rules, app errors, shared contracts. |
| `infra/` | HTTP, PostgreSQL, config, Wire bootstrap, runtime, health. |
| `cmd/` | Entry point — only place that knows all concrete types. |

## Import rules

| Package | Can import |
|---------|------------|
| `domain/` | stdlib only |
| `application/` | `domain/` |
| `infra/adapters/` | `domain/`, `database/sql` |
| `infra/httpserver/` | `application/`, `domain/`, `net/http` |
| `infra/bootstrap/`, `cmd/` | everything (wiring only) |

**Never:** handler → `database/sql`, use case → `net/http`, domain → `infra/`.

## Three types of validation

| Type | When | Where | Example |
|------|------|-------|---------|
| Config | Startup | `infra/config/` | `DATABASE_URL` required |
| Input | HTTP request | `httpserver/.../requests.go` | `validate:"email"` → 400 |
| Business | Use case | `application/usecase/` | email exists → 409 |

Handler order: decode → validate tags → `useCase.Exec()`.

## Where do I put X?

| I need to… | Go to… |
|------------|--------|
| Start app | `cmd/api/main.go` |
| Env config | `infra/config/` |
| DB connection | `infra/database/` |
| SQL queries | `infra/adapters/repository/` |
| Entity + repo interface | `domain/<feature>/` |
| Business logic | `application/usecase/<feature>/` |
| Shared ports (Kafka, Tx) | `application/contracts/` |
| App error kinds | `application/errors/` |
| Routes, handler, JSON | `infra/httpserver/<feature>/` |
| Middleware | `infra/httpserver/middlewares/` |
| Health checks | `infra/httpserver/health.go` |
| Wire / DI | `infra/bootstrap/` |
| Graceful shutdown | `infra/runtime/` |
| Migrations | `migrations/` |

## Add a new endpoint (checklist)

Example: `POST /products`

1. `domain/product/` — entity, errors, `repository.go` (interface)
2. `application/usecase/create-product/` — contract, `UseCase`, `Exec()`
3. `infra/adapters/repository/` — PostgreSQL implementation
4. `infra/httpserver/product/` — routes, handler, requests, responses
5. `infra/httpserver/router.go` — register routes
6. `cmd/api/main.go` or `infra/bootstrap/` — wire repo → use case → handler
7. `migrations/` — SQL schema
8. Test use case with fake repository

## Request flow: POST /register

```
Client → middlewares → handler (decode, validate)
      → use case (business rules, repo.Create)
      → repository (SQL) → handler → JSON response
```

## Three struct types (do not merge)

| Type | Location | Has json/validate tags? |
|------|----------|-------------------------|
| HTTP request DTO | `infra/httpserver/user/requests.go` | Yes |
| Use case contract | `application/usecase/register-user/contract.go` | No |
| Domain entity | `domain/user/entity.go` | No |

Flow: JSON → request DTO → contract → entity → SQL
