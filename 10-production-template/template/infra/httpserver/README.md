# infra/httpserver

HTTP layer — router, health, shared utilities, feature routes.

## Layout

```
httpserver/
├── server.go           http.Server, timeouts, graceful shutdown
├── router.go           mux, global middleware, feature registration
├── health.go           GET /livez, GET /readyz
├── routes/
│   ├── register.go     Method enum, Join, Register
│   └── sample/         reference feature (handler, requests, responses, routes)
├── middlewares/        chain, identity, logging, recovery, …
└── utils/
    ├── response.go     WriteSuccess, WriteError, WriteValidationError
    ├── validator.go    DecodeAndValidate
    ├── app-error.go    WriteAppError
    └── pagination.go   page helpers
```

## Route pattern

1. API version constant in `router.go` (e.g. `APIV2 = "/api/v2"`)
2. Feature mount in `routes/<feature>/routes.go` (e.g. `Mount = "samples"`)
3. Endpoints via `routes.Register(mux, routes.POST, base, "", handler.Create, mws...)`

## Handler rules

1. Decode + validate: `utils.DecodeAndValidate` → `utils.WriteValidationError`
2. Call use case
3. Map result: `ToApiResponse`
4. Write: `WriteSuccess` or `WriteAppError`
5. No SQL in handlers

## Middleware

| Scope | Middleware |
|---|---|
| Global | RequestID, Recovery, HttpAccessLog, StripTrailingSlash |
| Per-route | UserHeader, AccessShips |

Identity middleware is **per-route** so health checks stay open.

## Adding a feature

Copy `routes/sample/` structure:

```
routes/myfeature/
├── routes.go
├── handler.go
├── requests.go
└── responses.go
```

Register in `router.go` and Wire in `infra/bootstrap/sets.go`.
