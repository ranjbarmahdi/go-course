# Session Handoff — Production Go Backend Template

> **Last updated:** 2026-09-08
> **Module:** `template`
> **Location:** `10-production-template/template/`
> **Reference project:** `conning-configuration` (CTO's Go microservice)

This file captures everything decided and built across the chat sessions that produced this
template. Read it top to bottom before continuing work in a new session.

---

## How to resume in a new Cursor session

1. Open the `template/` folder in Cursor (it can be opened standalone — the module is self-contained).
2. Start a chat with:

   > Read `@HANDOFF.md` fully, then continue from **Next steps → Phase C**.
   > Start with the C0 quick fixes and guide me step by step.
   > I learn by doing — explain and hand me the code, don't implement the whole phase unless I ask.

---

## 1. What this project is

A **production-ready Go backend skeleton** built with clean architecture, Wire dependency
injection, PostgreSQL + Redis, structured logging, and a layered error model.

Two things to know about its current shape:

- The **infra layer is complete and runnable** — `go run ./cmd/api` boots, connects both
  databases, and serves health endpoints.
- The **`domain/route/` aggregate is temporary**. It exists so the domain patterns (value
  objects, aggregates, domain errors) have something concrete to demonstrate. It will be
  **removed later** to leave a pure, entity-free template.

---

## 2. Architecture

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

`infra/bootstrap/` is the composition root (Wire). `cmd/api/main.go` only loads config,
sets up logging, and hands off to Wire.

**Golden rule:** dependencies point inward. The domain never imports infra.

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

## 3. Current file layout

```
template/
├── cmd/api/main.go                     entry point
├── domain/
│   ├── domainerror/error.go            Kind + KindOf (domain error base)
│   ├── route/                          TEMPORARY aggregate (remove later)
│   │   ├── route.go                    Route entity + NewRoute/ReconstructRoute/AddWaypoint
│   │   ├── waypoint.go                 Waypoint entity
│   │   ├── errors.go                   domain errors (all use domainerror.New)
│   │   ├── validation.go               entity-specific validators
│   │   └── repository.go               RouteRepository interface
│   ├── shared/                         optional.go (clone helpers), geo.go
│   └── value-objects/
│       ├── id.go                       parseUUID (stdlib regex — no 3rd party in domain)
│       ├── route-id.go                 RouteID
│       └── waypoint-id.go              WaypointID
├── application/
│   ├── contracts/
│   │   ├── id-generator.go             UUID type + IDGenerator port
│   │   └── tx_runner.go                RunInTx[T] port
│   ├── errors/error.go                 apperrors Kind + Error + KindOf
│   ├── utils/error.go                  TranslateDomainError (domain Kind → app Kind)
│   └── usecase/                        EMPTY (README only)
├── infra/
│   ├── adapters/
│   │   ├── id-generator.go             UUIDGenerator (google/uuid)
│   │   └── repository/                 EMPTY (README only)
│   ├── bootstrap/
│   │   ├── wire.go                     wire.Build injector (build tag: wireinject)
│   │   ├── wire_gen.go                 generated
│   │   ├── sets.go                     provider sets
│   │   └── providers.go                small provider funcs
│   ├── config/config.go                env loading + validation
│   ├── database/
│   │   ├── postgres/
│   │   │   ├── postgres.go             Open/Close/NewPostgres (+ connect logs)
│   │   │   ├── executor.go             DBTX, txContextKey, Executor()
│   │   │   ├── transaction_manager.go  TransactionManager + NewRunInTx
│   │   │   └── indicator.go            readiness probe
│   │   └── redis/
│   │       ├── redis.go                direct + sentinel client, NewRedis
│   │       └── indicator.go            readiness probe
│   ├── httpserver/
│   │   ├── server.go                   http.Server + timeouts + graceful shutdown
│   │   ├── router.go                   mux + global middleware chain
│   │   ├── health.go                   GET /livez, GET /readyz
│   │   ├── middlewares/
│   │   │   ├── chain.go                Middleware type, Chain, WithMiddleware
│   │   │   ├── identity.go             UserHeader, AccessShips + context accessors
│   │   │   ├── logger.go               HttpAccessLog (debug-level access log)
│   │   │   ├── recovery.go             panic → 500
│   │   │   └── request-id.go           RequestID + context accessors
│   │   └── utils/
│   │       ├── response.go             response envelope, WriteSuccess, WriteError
│   │       ├── app-error.go            WriteAppError (Kind → HTTP status)
│   │       └── validator.go            shared validator + DecodeAndValidate
│   └── runtime/
│       ├── runner.go                   Runner + Indicator interfaces
│       └── app.go                      App.Run (errgroup)
├── migrations/                         EMPTY (README only — no SQL yet)
├── ARCHITECTURE.md
├── README.md
├── HANDOFF.md                          this file
├── go.mod
├── .env / .env.example
└── .gitignore
```

---

## 4. What is DONE

### Entry point and runtime

`cmd/api/main.go`: `main()` → `run()` → load config → configure `slog` → `signal.NotifyContext`
(SIGINT/SIGTERM) → `bootstrap.Wire(ctx, cfg)` → `defer cleanup()` → `app.Run(ctx)`.

`infra/runtime/`: `Runner` and `Indicator` interfaces, and `App` which starts all runners
with `errgroup`. Shutdown happens by context cancellation — each runner handles its own
graceful stop.

### Configuration

`infra/config/config.go` loads from environment via `caarlos0/env` with `godotenv` for local
`.env`, then validates. Fails fast on bad config.

| Variable | Notes |
|---|---|
| `PORT` | default 8080 |
| `ENV` | default `development` |
| `DATABASE_URL` | **required** |
| `JWT_SECRET` | currently required, min 32 chars — see Known issues |
| `LOG_LEVEL` | default `info` |
| `REDIS_URL` | required in direct mode |
| `REDIS_SENTINEL_HOST/PORT/USER/MASTER/PASS` | required in sentinel mode |

**Redis mode selection** mirrors the CTO's NestJS `RedisDB` service:

```go
func (c *Config) RedisSelfHost() bool {
    return c.Env == "self-host" || c.Env == "development"
}
```

- `development` / `self-host` → direct connection via `REDIS_URL`
- anything else (production, staging) → Sentinel failover client

Redis validation is **mode-aware**: sentinel fields are only required when not in self-host
mode, which is why none of the Redis env tags use `,required`.

### Databases

**PostgreSQL** (`infra/database/postgres/`):
- `Open` configures the pool (25 max open, 5 idle, 30 min lifetime) and pings with a 5s timeout.
- `NewPostgres` is the Wire provider returning `(*sql.DB, cleanup, error)` and logs
  `connected` / `disconnected` with `component=postgres`.
- `executor.go` holds the transaction-context plumbing: `DBTX` interface, unexported
  `txContextKey`, and `Executor(ctx, db)` which returns the ambient `*sql.Tx` when present,
  otherwise the pool. **Repositories must call `Executor`, never `db` directly** — that is
  what makes them transparently transaction-aware.
- `transaction_manager.go` implements `RunInTx` and exposes `NewRunInTx(db) contracts.RunInTx[error]`.

**Redis** (`infra/database/redis/`):
- `newClient` branches on `cfg.RedisSelfHost()` — `ParseURL` + `NewClient` for direct,
  `NewFailoverClient` for sentinel (passing `SentinelPass` as both `Password` and
  `SentinelPassword`, matching ioredis behavior in the Nest service).
- `NewRedis` pings before returning, so startup fails fast if Redis is unreachable.
- Import is aliased as `goredis` to avoid clashing with the local package name.

Both databases expose an `Indicator` used by `/readyz`.

### Dependency injection (Wire)

`infra/bootstrap/` holds the composition root.

| Set | Contents |
|---|---|
| `DatabaseSet` | `postgres.NewPostgres`, `redis.NewRedis` |
| `IndicatorsSet` | `postgres.NewIndicator`, `redis.NewIndicator` |
| `HttpSet` | `provideIndicators`, `provideAddr`, `NewRouter`, `NewServer` |
| `RuntimeSet` | `provideRunners`, `newApp` |
| `AdapterSet` | `postgres.NewRunInTx`, `provideIDGenerator` — **defined but not yet in `wire.Build`** |

`newApp` wraps the variadic `runtime.NewApp` because Wire handles `[]Runner` better through
a non-variadic function.

Regenerate with:

```bash
cd infra/bootstrap
wire
```

### HTTP layer

`server.go` sets full timeouts (added deliberately — the earlier version had none):

| Timeout | Value | Why |
|---|---|---|
| `ReadHeaderTimeout` | 5s | Slowloris defense — the critical one |
| `ReadTimeout` | 15s | |
| `WriteTimeout` | 15s | |
| `IdleTimeout` | 60s | keep-alive reaping |

Graceful shutdown: `Start` races `ListenAndServe` against `ctx.Done()`; on cancellation it
calls `Shutdown` with a 5s timeout and drains the serve error.

`router.go` builds the mux, registers health routes, and wraps everything in the global chain.

`health.go` serves `GET /livez` (always 200) and `GET /readyz` (200 only if every `Indicator`
reports ready).

### Middleware

`chain.go` defines the core types. `Chain(h, a, b)` produces `a(b(h))` — **the first middleware
listed is the outermost**. `WithMiddleware` does the same for a single `http.HandlerFunc`,
used for per-route middleware.

| Middleware | Scope | Behavior |
|---|---|---|
| `Recovery` | global | recovers panics, logs, returns 500 |
| `RequestID` | global | reuses incoming `x-request-id` or generates one; echoes it in the response header and stores it in context |
| `HttpAccessLog(logger)` | global | logs method, path, status, bytes, duration, client IP, user agent, referer, request ID |
| `UserHeader` | **per-route** | parses `x-user` JSON into `OwnUser`, stores in context, 401 if missing/invalid |
| `AccessShips` | **per-route** | parses `x-access-ship` JSON array of ship IDs, stores in context, 401 if missing/invalid |

Identity middleware is deliberately **not global** so that `/livez` and `/readyz` stay
reachable without gateway headers.

`HttpAccessLog` only emits at **debug** level — set `LOG_LEVEL=debug` to see HTTP request logs.
`statusCapturingWriter` wraps the `ResponseWriter` to capture status and byte count.

### HTTP utilities

`response.go` defines the shared envelope used across all responses (copied from
conning-configuration):

```json
{ "status": true, "message": "...", "data": {} }
```

- `WriteSuccess(w, status, message, data)`
- `WriteError(w, status, message, optionalError...)`
- `writeJSON` centralizes the `Content-Type` header and encoding.

`app-error.go` maps application error kinds to HTTP statuses:

| `apperrors.Kind` | HTTP |
|---|---|
| `InvalidInput` | 400 |
| `Unauthorized` | 401 |
| `Domain` | 400 |
| `NotFound` | 404 |
| `Conflict` | 409 |
| `Internal` | 500 — logs the real error, returns a generic message (no internals leaked) |

`validator.go` holds a single shared `validator.New(...)` instance (reflection caching makes
per-request construction wasteful) plus `DecodeAndValidate(r, dst)` which combines JSON
decode and struct validation in one call.

### Error model (three layers)

This is the piece that took the most iteration. Errors flow through three distinct layers:

```
domain/domainerror       →  application/errors      →  HTTP status
(Kind: InvalidInput,        (Kind: InvalidInput,       (400, 401, 404,
 NotFound, Conflict,         Unauthorized, Domain,      409, 500)
 Domain, Internal)           NotFound, Conflict,
                             Internal)
```

**Domain layer** — every domain error carries a `Kind`, so callers can classify without
matching individual sentinels:

```go
var ErrNotFound = domainerror.New(domainerror.NotFound, "route not found")
var ErrDuplicateOrderNumber = domainerror.New(domainerror.Conflict, "duplicate order number")
```

`errors.Is(err, route.ErrNotFound)` still works, because `*domainerror.Error` implements `error`.

**Application layer** — `TranslateDomainError` switches on `domainerror.KindOf(err)` and
produces an `apperrors.Error`. This mirrors the CTO's `TranslateDomainError` in
conning-configuration. Use cases call it at every boundary where a domain error can surface.

**HTTP layer** — `WriteAppError` switches on `apperrors.KindOf(err)` and writes the response.

The reason for `Kind`-based dispatch instead of long `errors.Is` chains: adding a new domain
error requires no changes to the translator or the HTTP writer.

### Value objects

`RouteID` and `WaypointID` are distinct struct types wrapping a validated UUID string, so a
waypoint ID can never be passed where a route ID is expected. UUID validation uses a stdlib
regex in `id.go` — **no third-party dependency in the domain**.

Each ID type has two constructors:

```go
// Untrusted input (HTTP path, request body, query string)
func NewRouteID(value string) (RouteID, error)

// Trusted input (idGen.NewUUID() output, or a DB row being reconstructed)
func RouteIDFromTrusted(value string) RouteID
```

---

## 5. Key decisions (already settled — don't relitigate)

| Topic | Decision | Reasoning |
|---|---|---|
| Optional fields | `*string`, `*float64`; normalize in constructors; clone in getters | prevents callers mutating internal state through a returned pointer |
| ID generation | `IDGenerator` port in `application/contracts`, `UUIDGenerator` adapter in infra | keeps `google/uuid` out of the domain |
| ID parsing location | value object (`NewRouteID`) for format; use case orchestrates; entity only checks `IsZero()` | domain owns its own ID rules |
| Parser injection | **Never** pass a parser function or `idGen.ParseUUID` into `NewRoute` | the constructor must receive already-valid domain types |
| DB connection | opened in Wire/bootstrap, not in `main` | main stays thin; cleanup is composed by Wire |
| Transactions | context propagation + `RunInTx[error]` port; repositories use `Executor(ctx, db)` | repositories work identically inside and outside a transaction |
| Multi-database naming | subpackages `database/postgres/`, `database/redis/` | room for `database/mongo/` later |
| Authentication | gateway injects `x-user` and `x-access-ship` headers; no JWT verification in-service | matches the existing microservice fleet |
| Identity middleware scope | per-route, never global | health probes must not require headers |
| CORS | intentionally **omitted** — the API gateway handles it | duplicating it produces conflicting headers |
| Access log level | debug only | avoids noise at info in production |
| `Recovery` in `Chain` | pass `middlewares.Recovery`, not `middlewares.Recovery()` | it *is* the middleware; `HttpAccessLog(logger)` is a factory and does need calling |
| Wire and unused sets | a set listed in `wire.Build` with no consumer is a hard error — leave it commented until something needs it | why `AdapterSet` is currently commented out |
| Internal (500) responses | log the real error, return a generic message | avoid leaking internals to clients |

---

## 6. Known issues / open items

These are real and should be handled when work resumes.

**1. `provideIndicators` silently drops Redis.** The signature accepts the Redis indicator but
the body ignores it, so `/readyz` never checks Redis:

```go
// infra/bootstrap/providers.go — current (wrong)
func provideIndicators(postgres *postgres.Indicator, rd *redis.Indicator) []runtime.Indicator {
    return []runtime.Indicator{postgres}
}

// should be
return []runtime.Indicator{postgres, rd}
```

It compiles because unused *parameters* are legal in Go — only unused locals and imports are errors.

**2. `AdapterSet` is commented out** in `wire.go`. This is currently correct: nothing consumes
`RunInTx` or `IDGenerator` yet, and Wire rejects unused provider sets. Uncomment it in the same
change that introduces the first use case.

**3. `JWT_SECRET` is required but never read.** Authentication comes from gateway headers.
Either drop the field or make it optional (remove `,required` and only length-check when
non-empty), otherwise the template won't boot without a secret nothing uses.

**4. `Runner.Stop` is dead API.** `App.Run` never calls `Stop` or `Name` — the HTTP server
handles its own shutdown on `ctx.Done()`. Either shrink the interface to `Start` + `Name` and
log the name when starting each runner, or have `App` orchestrate `Stop`. Avoid the latter
without removing the server's internal shutdown, or you'll double-shutdown.

**5. Middleware chain order.** Currently `Recovery` is outermost, then `RequestID`. Consider
putting `RequestID` first so panics recovered by `Recovery` can be logged with their request ID.

**6. No vertical slice.** `application/usecase/`, `infra/adapters/repository/`,
`infra/httpserver/<feature>/`, and `migrations/` are all README-only. This is the main gap.

**7. `.env.example` has empty values.** It lists the right keys but no placeholder values or
comments explaining the two Redis profiles.

**8. Documentation drift.** Several READMEs still reference `user` / `register-user` from the
original plan while the code uses `route`. `ARCHITECTURE.md` says `cmd/api/main.go` wires
everything, but Wire in `infra/bootstrap/` actually does. `infra/database/README.md` and
`infra/config/README.md` predate Redis.

---

## 7. Next steps — Phase C: first vertical slice

The goal is one complete `POST /routes` flow that exercises every layer. This is also what
unblocks `AdapterSet` in Wire.

### C0 — Quick fixes (do first)

- [ ] Fix `provideIndicators` to return both indicators
- [ ] Decide on `JWT_SECRET` (optional or removed)
- [ ] Decide on `Runner.Stop` (recommend shrinking the interface)
- [ ] Consider reordering `RequestID` before `Recovery`

### C1 — Migration

- [ ] `migrations/001_create_routes.sql`

`from` and `to` are reserved words in SQL — use `from_port` / `to_port`:

```sql
CREATE TABLE IF NOT EXISTS routes (
    id            UUID PRIMARY KEY,
    vessel_id     TEXT NOT NULL,
    name          TEXT NOT NULL,
    from_port     TEXT,
    to_port       TEXT,
    creator_type  TEXT NOT NULL,
    creator_id    TEXT,
    file_type     TEXT,
    created_at    TIMESTAMPTZ NOT NULL
);

CREATE TABLE IF NOT EXISTS waypoints (
    id                  UUID PRIMARY KEY,
    route_id            UUID NOT NULL REFERENCES routes(id) ON DELETE CASCADE,
    latitude            DOUBLE PRECISION NOT NULL,
    longitude           DOUBLE PRECISION NOT NULL,
    source_ref          TEXT,
    order_number        INTEGER NOT NULL,
    name                TEXT,
    waypoint_type       TEXT NOT NULL,
    planned_speed_knots DOUBLE PRECISION,
    UNIQUE (route_id, order_number)
);
```

The `UNIQUE (route_id, order_number)` constraint enforces the `ErrDuplicateOrderNumber`
invariant at the database level too.

Migration tool: **goose** is the lightest fit (single binary, plain SQL files with
`-- +goose Up` / `-- +goose Down` markers).

### C2 — Repository adapter

- [ ] `infra/adapters/repository/route.go`

```go
type RouteRepository struct {
    db *sql.DB
}

func NewRouteRepository(db *sql.DB) *RouteRepository {
    return &RouteRepository{db: db}
}

// compile-time check — catches interface drift immediately
var _ route.RouteRepository = (*RouteRepository)(nil)

func (r *RouteRepository) CreateRoute(ctx context.Context, rt route.Route) error {
    exec := postgres.Executor(ctx, r.db) // tx when present, pool otherwise
    // INSERT route, then loop waypoints using the same exec
}
```

In `FindByID`, translate the driver error at the boundary so no `database/sql` type leaks
back into the application:

```go
if errors.Is(err, sql.ErrNoRows) {
    return nil, route.ErrNotFound
}
```

### C3 — Use case

- [ ] `application/usecase/create-route/contract.go` — plain input struct, **no** json or
      validate tags (those belong to the HTTP DTO)
- [ ] `application/usecase/create-route/usecase.go`

```go
type UseCase struct {
    repo    route.RouteRepository
    idGen   contracts.IDGenerator
    runInTx contracts.RunInTx[error]
}

func (u *UseCase) Exec(ctx context.Context, in Input) (valueobjects.RouteID, error) {
    routeID := valueobjects.RouteIDFromTrusted(string(u.idGen.NewUUID()))

    rt, err := route.NewRoute(routeID, in.VesselID, in.Name, /* ... */)
    if err != nil {
        return valueobjects.RouteID{}, apputils.TranslateDomainError(err)
    }

    err = u.runInTx(ctx, func(ctx context.Context) error {
        return u.repo.CreateRoute(ctx, *rt)
    })
    if err != nil {
        return valueobjects.RouteID{}, apputils.TranslateDomainError(err)
    }

    return routeID, nil
}
```

### C4 — HTTP feature package

- [ ] `infra/httpserver/route/requests.go` — DTOs with json + validate tags
- [ ] `infra/httpserver/route/responses.go`
- [ ] `infra/httpserver/route/handler.go`
- [ ] `infra/httpserver/route/routes.go`

Handler order is decode → validate → exec → write:

```go
var req CreateRouteRequest
if err := utils.DecodeAndValidate(r, &req); err != nil {
    utils.WriteError(w, http.StatusBadRequest, "invalid request", err)
    return
}

id, err := h.createRoute.Exec(r.Context(), req.ToInput())
if err != nil {
    utils.WriteAppError(w, err)
    return
}

utils.WriteSuccess(w, http.StatusCreated, "route created", CreateRouteResponse{ID: id.String()})
```

Routes apply per-route middleware:

```go
func RegisterRoutes(mux *http.ServeMux, handler *Handler) {
    mux.Handle("POST /api/v1/routes",
        middlewares.WithMiddleware(handler.Create, middlewares.UserHeader))
}
```

### C5 — Wire it up

- [ ] Add to `sets.go`:

```go
var RepositorySet = wire.NewSet(
    repository.NewRouteRepository,
    wire.Bind(new(route.RouteRepository), new(*repository.RouteRepository)),
)

var UseCaseSet = wire.NewSet(createroute.New)
var HandlerSet  = wire.NewSet(routehttp.NewHandler)
```

`wire.Bind` tells Wire "when something needs the `route.RouteRepository` interface, supply
`*repository.RouteRepository`."

- [ ] Uncomment `AdapterSet` in `wire.go` and add the new sets
- [ ] Update `NewRouter` to accept the handler and call `RegisterRoutes`
- [ ] Run `wire` in `infra/bootstrap`

### C6 — Tests

- [ ] In-memory fake implementing `route.RouteRepository`
- [ ] One use case test using the fake

There are currently **zero** `_test.go` files, and `ARCHITECTURE.md` tells readers to "test the
use case with a fake repository" — so the template should ship an example of exactly that.

---

## 8. Phase D — turning this into a pure template

Once the slice works end to end:

- [ ] Remove `domain/route/` (or reduce it to a minimal illustrative example)
- [ ] Remove the corresponding use case, repository, and HTTP feature package
- [ ] Sync all READMEs and `ARCHITECTURE.md` with the final structure
- [ ] Update `.env.example` with placeholder values and both Redis profiles

Deliberately deferred (user's call): `Makefile`, `docker-compose.yml`, `Dockerfile`,
`.dockerignore`, CI workflow.

---

## 9. Commands

```bash
# from template/
go build ./...
go run ./cmd/api

# regenerate Wire (must be run from the bootstrap package)
cd infra/bootstrap && wire
```

Note: `go build .` at the template root fails because there are no `.go` files there — use
`./...` or `./cmd/api`.

### Manual verification

```bash
curl http://localhost:8080/livez     # 200 always
curl http://localhost:8080/readyz    # 200 only when postgres + redis are reachable

# once feature routes exist:
curl -X POST http://localhost:8080/api/v1/routes                 # 401 (no x-user)
curl -X POST -H 'x-user: {"id":"1"}' http://localhost:8080/...   # passes identity middleware
```

Set `LOG_LEVEL=debug` to see HTTP access logs.

---

## 10. Dependencies

Direct:

```
github.com/caarlos0/env/v10           config from environment
github.com/joho/godotenv              local .env loading
github.com/google/wire                compile-time DI
github.com/google/uuid                UUID generation (infra only)
github.com/jackc/pgx/v5               PostgreSQL driver (via database/sql stdlib interface)
github.com/redis/go-redis/v9          Redis client (direct + sentinel)
github.com/go-playground/validator/v10 request DTO validation
golang.org/x/sync                     errgroup for App.Run
```

The domain layer imports **none** of these — only stdlib.
