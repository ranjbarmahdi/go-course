# infra/database

Purpose: PostgreSQL connection pool and health indicator.

Rules:
- Open pool at startup, close on shutdown
- Ping for `/readyz` health check

Files (Topic 06):
- `postgres.go` — Connect, Ping, Close
- `indicator.go` — PostgresIndicator for readiness probe
