# infra/adapters

Purpose: Implementations of domain and application interfaces.

Subfolders:
- `repository/` — PostgreSQL repos (implements `domain/*/Repository`)

Future:
- `kafka-producer/` — implements `application/contracts.MessagePublisher`

Rules:
- SQL and external SDK calls live here only
- Map DB rows ↔ domain entities
