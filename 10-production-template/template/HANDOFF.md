# Session Handoff — Production Go Backend Template

> **Last updated:** 2026-09-09  
> **Module:** `template`  
> **Location:** `10-production-template/template/`  
> **Reference:** conning-configuration (CTO Go microservice), organization (NestJS), geotrace (Express-style routing)

Read this file before continuing work in a new session.

---

## How to resume

1. Open the `template/` folder in Cursor.
2. Start a chat with:

   > Read `@HANDOFF.md` fully, then continue from **Next steps**.

3. Default learning mode: explain and hand code — implement only when asked.

---

## What this project is

A **production-ready Go backend skeleton** with clean architecture, Wire DI, PostgreSQL + Redis, structured logging, goose migrations, and a layered error model.

**Current state:** the template boots end-to-end. The **sample** vertical slice (`POST` + `GET` sample) exercises every layer from HTTP to SQL.

---

## Architecture

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
infra/adapters/            PostgreSQL repos, ID generator
```

`infra/bootstrap/` is the composition root (Wire). `cmd/api/main.go` loads config, configures logging, optionally auto-migrates, then calls Wire.

**Golden rule:** dependencies point inward. Domain never imports infra.

### Import rules

| Package | May import |
|---|---|
| `domain/` | stdlib only |
| `application/` | `domain/` |
| `infra/adapters/` | `domain/`, `database/sql` |
| `infra/httpserver/` | `application/`, `domain/`, `net/http` |
| `infra/bootstrap/`, `cmd/` | everything (wiring only) |

**Never:** handler → `database/sql`, use case → `net/http`, domain → `infra/`.

---

## File layout

```
template/
├── cmd/
│   ├── api/main.go                 entry point
│   └── migrate/main.go             goose CLI (up/down/status)
├── domain/
│   ├── domainerror/                Kind + KindOf
│   ├── sample/                     Sample entity, errors, repository interface
│   ├── shared/                     optional field helpers
│   └── value-objects/              SampleID, OwnUser, parseUUID
├── application/
│   ├── contracts/                  IDGenerator, RunInTx
│   ├── errors/                     apperrors Kind + KindOf
│   ├── shared/                     pagination types
│   ├── utils/error.go              TranslateDomainError
│   └── usecase/sample/
│       ├── response.go, mapper.go  app-layer results
│       ├── create-sample/          POST use case
│       └── get-sample/             GET by ID use case
├── infra/
│   ├── adapters/
│   │   ├── id-generator.go         UUIDGenerator
│   │   └── repository/
│   │       └── sample-repository.go
│   ├── bootstrap/                  Wire sets + providers
│   ├── config/config.go            env loading + validation
│   ├── database/
│   │   ├── migrate/migrate.go      goose wrapper
│   │   ├── postgres/               pool, executor, RunInTx, indicator
│   │   └── redis/                  direct + sentinel, indicator
│   ├── httpserver/
│   │   ├── server.go, router.go, health.go
│   │   ├── routes/
│   │   │   ├── register.go         Method enum, Join, Register
│   │   │   └── sample/             handler, requests, responses, routes
│   │   ├── middlewares/            chain, identity, logging, recovery, …
│   │   └── utils/                  response, validation, pagination, app-error
│   └── runtime/                    App, Runner, graceful shutdown
├── migrations/
│   └── 00001_sample_create.sql
├── ARCHITECTURE.md
├── README.md
├── HANDOFF.md
├── go.mod
└── .env / .env.example
```

---

## What is done

### Runtime and entry point

- `cmd/api/main.go`: config → slog → signal context → optional auto-migrate → `bootstrap.Wire` → `app.Run`
- `infra/runtime/`: `Runner`, `Indicator`, `App` with errgroup
- HTTP server graceful shutdown on context cancel

### Configuration

| Variable | Notes |
|---|---|
| `PORT` | default `8080` |
| `ENV` | default `development` |
| `DATABASE_URL` | required |
| `DB_SCHEMA` | required — operator creates schema; app sets `search_path` |
| `LOG_LEVEL` | default `info` |
| `AUTO_MIGRATE` | default `false`; when `true`, runs goose before Wire |
| `REDIS_URL` | required in development/self-host mode |
| `REDIS_SENTINEL_*` | required in production/staging (sentinel mode) |

Redis mode: `development` / `self-host` → direct URL; otherwise → Sentinel.

JWT was removed — auth comes from gateway headers (`x-user`, `x-access-ship`).

### Databases

- PostgreSQL: pool, schema via `search_path`, `Executor(ctx, db)` for tx-aware repos, `NewRunInTx`
- Redis: direct + sentinel, ping on startup
- Both expose readiness indicators for `/readyz`

### Migrations (goose)

- `migrations/00001_sample_create.sql` — `sample_table`
- `cmd/migrate` — `up`, `down`, `status`
- Schema must exist before boot: `CREATE SCHEMA <DB_SCHEMA>;`

### Wire (bootstrap)

| Set | Contents |
|---|---|
| `DatabaseSet` | postgres, redis |
| `IndicatorsSet` | postgres + redis indicators |
| `AdapterSet` | RunInTx, IDGenerator, SampleRepository + bind |
| `SampleUseCaseSet` | create-sample, get-sample |
| `SampleHandlerSet` | sample HTTP handler |
| `HttpSet` | indicators, addr, router, server |
| `RuntimeSet` | runners, app |

Regenerate: `cd infra/bootstrap && go generate`

### HTTP layer

**Global middleware** (outer → inner): `RequestID` → `Recovery` → `HttpAccessLog` → `StripTrailingSlash` → mux

**Per-route:** `UserHeader` on sample routes

**API prefix:** `APIV2 = /api/v2` in `router.go`

**Sample routes:**

| Method | Path | Handler |
|---|---|---|
| POST | `/api/v2/samples` | Create |
| GET | `/api/v2/samples/{id}` | GetByID |

**Route registration pattern:**

- Version prefix in `router.go`
- Resource mount in feature `routes.go` (`Mount = "samples"`)
- Shared `routes.Register(mux, method, base, path, handler, mws...)`
- Typed HTTP methods: `routes.POST`, `routes.GET`, …

### Request validation

- HTTP DTOs in `routes/sample/requests.go` with `json` + `validate` tags
- `utils.DecodeAndValidate` — JSON decode + validator/v10
- `utils.WriteValidationError` — structured field errors for clients
- Unknown JSON fields → `"field is not allowed"` (via `DisallowUnknownFields`)
- Business validation stays in use cases / domain

### Error model (three layers)

```
domain/domainerror  →  application/errors  →  HTTP status (WriteAppError)
```

Use cases call `TranslateDomainError` at domain boundaries.

### Sample vertical slice

- Domain: `Sample` entity with optional `*string`/`*int`, `SampleRepository` interface
- Use cases: `create-sample` (RunInTx), `get-sample`
- Repository: `sample-repository.go` with `Executor`, maps `sql.ErrNoRows` → `ErrNotFound`
- HTTP: handler maps HTTP DTO → use case → `ToApiResponse`

---

## Key decisions (settled)

| Topic | Decision |
|---|---|
| Optional fields | `*string`, `*int`; clone in getters |
| ID generation | `IDGenerator` port in application; `UUIDGenerator` in infra |
| ID parsing | `NewSampleID` in value objects for untrusted input |
| Transactions | `RunInTx` + `Executor(ctx, db)` in repositories |
| Auth | Gateway headers; no in-service JWT |
| Identity middleware | per-route only (health stays open) |
| CORS | omitted — gateway handles it |
| Route prefixes | Express-style: global API version + feature mount |
| Validation tags | HTTP DTOs only — never on use case or domain structs |
| Schema | operator creates; app verifies via `search_path` |

---

## Known issues / follow-ups

1. **HttpAccessLog is debug-only** — set `LOG_LEVEL=debug` to see access logs, or change logger to `InfoContext`.
2. **get-sample** — prefer `valueobjects.NewSampleID` over `idGenerator.ParseUUID`; use `TranslateDomainError` for repo errors instead of mapping everything to NotFound.
3. **Trailing slash on POST /samples/** — `StripTrailingSlash` may not fully normalize; consider registering both paths in `routes.Register`.
4. **No tests yet** — zero `_test.go` files; add use case test with fake repository.
5. **`.env.example`** — empty placeholders; could add sample values and comments.
6. **Package logger timing** — some infra loggers capture `slog.Default()` before `main` sets level (nested "INFO" in messages).

---

## Next steps

### Phase D — polish and template cleanup

- [ ] Add use case tests with in-memory fake repository
- [ ] Fix get-sample error handling and ID parsing
- [ ] Decide access log level (info vs debug)
- [ ] Register trailing-slash variants in `routes.Register` if needed
- [ ] Improve `.env.example` with documented placeholders
- [ ] Optional: remove sample slice to leave pure skeleton (or keep as reference)

### Phase E — deferred (user's call)

- Makefile, docker-compose, Dockerfile, CI workflow

---

## Commands

```bash
# from template/
go build ./...
go run ./cmd/api

# migrations
go run ./cmd/migrate status
go run ./cmd/migrate up
go run ./cmd/migrate down

# wire
cd infra/bootstrap && go generate
```

Note: use `go build ./...` — not `go build .` at repo root.

### Manual verification

```bash
# health (no auth)
curl http://localhost:8080/livez
curl http://localhost:8080/readyz

# create sample (requires x-user with SUPER_ADMIN)
curl -X POST http://localhost:8080/api/v2/samples \
  -H "Content-Type: application/json" \
  -H 'x-user: {"username":"admin","first_name":"Test","last_name":"User","id":"1","admin":true,"user_lvl":"SUPER_ADMIN","organization_id":"org-1"}' \
  -d '{"name":"test","number":42}'

# get by id (use uuid from create response)
curl http://localhost:8080/api/v2/samples/{uuid} \
  -H 'x-user: {"username":"admin","first_name":"Test","last_name":"User","id":"1","admin":true,"user_lvl":"SUPER_ADMIN","organization_id":"org-1"}'
```

Set `LOG_LEVEL=debug` for HTTP access logs.

---

## Dependencies

```
github.com/caarlos0/env/v10           config
github.com/joho/godotenv              local .env
github.com/google/wire                DI
github.com/google/uuid                UUID (infra only)
github.com/jackc/pgx/v5               PostgreSQL driver
github.com/redis/go-redis/v9          Redis
github.com/go-playground/validator/v10 HTTP DTO validation
github.com/pressly/goose/v3           migrations
golang.org/x/sync                     errgroup
```

Domain imports stdlib only.
