# infra/httpserver/user

Purpose: User HTTP API — routes, handler, request/response DTOs.

Rules:
- JSON + validate tags on requests.go only
- Handler calls `registeruser.UseCase` interface
- Always `return` after writing error response

Files (Topic 05):
- `routes.go` — mux.Handle + middleware
- `handler.go` — Register, Login handlers
- `requests.go` — JSON input + validator tags
- `responses.go` — JSON output + mappers from domain

Routes (planned):
- POST /api/v1/register
- POST /api/v1/login
