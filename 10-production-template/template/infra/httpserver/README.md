# infra/httpserver

Purpose: HTTP layer — router, health, shared utils, feature handlers.

Contains:
- `router.go` — register all routes
- `health.go` — GET /livez, GET /readyz
- `validator.go` — shared validator instance
- `middlewares/` — logging, recovery, cors, auth
- `utils/` — WriteSuccess, WriteAppError
- `user/` — user routes, handler, requests, responses

Rules:
- Handlers decode JSON, validate, call use case, return JSON
- No SQL in handlers
