# infra

Purpose: Technical implementations — HTTP, database, config, wiring, runtime.

Contains:
- `config/` — load and validate environment variables
- `database/` — PostgreSQL connection, ping, close
- `adapters/` — repository implementations, Kafka, JWT helpers
- `httpserver/` — routes, handlers, middleware, health
- `bootstrap/` — Google Wire dependency injection (Topic 11)
- `runtime/` — App, Runner, graceful shutdown

Rules:
- Implements interfaces from `domain/` and `application/contracts/`
- Only `bootstrap/` and `cmd/` wire concrete types together
