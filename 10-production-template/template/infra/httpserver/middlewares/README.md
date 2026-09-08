# infra/httpserver/middlewares

Purpose: HTTP middleware — logging, recovery, CORS, authentication.

Rules:
- Global middleware wraps entire router
- Per-route middleware via `WithMiddleware()` in routes.go

Files (Topic 05):
- `chain.go` — Chain(), WithMiddleware()
- `logging.go`, `recovery.go`, `cors.go`, `auth.go`

Chain order (global): cors → logging → recovery → mux
