# Architecture

Layer rules and "where do I put X?" for this template.

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
infra/adapters/            PostgreSQL, ID generator
```

`infra/bootstrap/` wires concrete types (Wire). `cmd/api/main.go` only loads config and starts the app.

**Golden rule:** dependencies point inward. Domain never imports infra.

## Layer jobs

| Layer | Job |
|---|---|
| `domain/` | Entities, domain errors, repository **interfaces**. No HTTP, no SQL. |
| `application/` | Use cases, business rules, app errors, shared contracts. |
| `infra/` | HTTP, PostgreSQL, Redis, config, bootstrap, runtime. |
| `cmd/` | Entry points — thin; no business logic. |

## Import rules

| Package | Can import |
|---|---|
| `domain/` | stdlib only |
| `application/` | `domain/` |
| `infra/adapters/` | `domain/`, `database/sql` |
| `infra/httpserver/` | `application/`, `domain/`, `net/http` |
| `infra/bootstrap/`, `cmd/` | everything (wiring only) |

**Never:** handler → `database/sql`, use case → `net/http`, domain → `infra/`.

## Three types of validation

| Type | When | Where | Example |
|---|---|---|---|
| Config | Startup | `infra/config/` | `DATABASE_URL` required |
| Input | HTTP request | `routes/*/requests.go` | `validate:"required"` → 400 |
| Business | Use case | `application/usecase/` | not found → 404 |

Handler order: **decode → validate tags → use case → write response**.

Use `utils.DecodeAndValidate` + `utils.WriteValidationError` — never expose raw JSON parser errors.

## Where do I put X?

| I need to… | Go to… |
|---|---|
| Start app | `cmd/api/main.go` |
| Run migrations | `cmd/migrate/main.go`, `migrations/` |
| Env config | `infra/config/` |
| DB connection | `infra/database/` |
| SQL queries | `infra/adapters/repository/` |
| Entity + repo interface | `domain/<feature>/` |
| Business logic | `application/usecase/<feature>/` |
| Shared ports (Tx, ID) | `application/contracts/` |
| App error kinds | `application/errors/` |
| Routes, handler, JSON DTOs | `infra/httpserver/routes/<feature>/` |
| Shared route helpers | `infra/httpserver/routes/register.go` |
| Middleware | `infra/httpserver/middlewares/` |
| Health checks | `infra/httpserver/health.go` |
| Wire / DI | `infra/bootstrap/` |
| Graceful shutdown | `infra/runtime/` |

## Add a new endpoint (checklist)

Example: `POST /products`

1. `domain/product/` — entity, errors, `repository.go`
2. `application/usecase/create-product/` — use case + request struct (no json tags)
3. `infra/adapters/repository/` — PostgreSQL implementation
4. `infra/httpserver/routes/product/` — routes, handler, requests, responses
5. `infra/httpserver/router.go` — call `product.RegisterRoutes(mux, APIV2, handler)`
6. `infra/bootstrap/sets.go` — Wire sets for repo, use case, handler
7. `migrations/` — SQL schema
8. Test use case with fake repository

## Request flow (sample create)

```
Client → global middlewares → UserHeader (per-route)
      → handler.DecodeAndValidate → create-sample use case
      → sample repository (SQL) → JSON response
```

## Three struct types (do not merge)

| Type | Location | json/validate tags? |
|---|---|---|
| HTTP request DTO | `infra/httpserver/routes/sample/requests.go` | Yes |
| Use case input | `application/usecase/sample/create-sample/request.go` | No |
| Domain entity | `domain/sample/sample.go` | No |

Flow: JSON → HTTP DTO → use case request → domain entity → SQL

## Reference implementation

Copy the **sample** feature when adding new endpoints:

- `domain/sample/`
- `application/usecase/sample/`
- `infra/httpserver/routes/sample/`
- `infra/adapters/repository/sample-repository.go`
