# infra/httpserver/middlewares

HTTP middleware for logging, recovery, identity, and request tracing.

## Files

| File | Purpose |
|---|---|
| `chain.go` | `Middleware` type, `Chain`, `WithMiddleware` |
| `request-id.go` | `RequestID` — reads or generates `x-request-id` |
| `recovery.go` | `Recovery` — panic → 500 |
| `logger.go` | `HttpAccessLog` — request/response logging |
| `identity.go` | `UserHeader`, `AccessShips` + context accessors |
| `training-slash.go` | `StripTrailingSlash` — normalize trailing `/` |

## Chain order

`Chain(h, a, b, c)` → `a(b(c(h)))` — **first listed is outermost**.

Current global chain in `router.go`:

```
RequestID → Recovery → HttpAccessLog → StripTrailingSlash → mux
```

## Global vs per-route

- **Global** — wrap the entire router in `router.go`
- **Per-route** — pass to `routes.Register(..., middlewares.UserHeader)`

Health endpoints (`/livez`, `/readyz`) must not require auth headers.

## Access logs

`HttpAccessLog` emits at **debug** level. Set `LOG_LEVEL=debug` to see request logs.

## Auth headers

| Header | Middleware | Format |
|---|---|---|
| `x-user` | `UserHeader` | JSON → `OwnUser` |
| `x-access-ship` | `AccessShips` | JSON array of ship IDs |

Missing or invalid headers → 401.
