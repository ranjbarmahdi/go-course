# application

Purpose: Use cases, application errors, shared contracts (ports).

Contains:
- `usecase/` — one folder per action (register-user, login-user, …)
- `errors/` — typed errors with Kind (maps to HTTP status)
- `contracts/` — shared ports (MessagePublisher, RunInTx)

Rules:
- Imports `domain/` only (not infra, not net/http)
- Business rules live in use cases, not in validator tags
