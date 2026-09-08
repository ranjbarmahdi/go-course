# infra/adapters/repository

Purpose: PostgreSQL implementations of domain repository interfaces.

Rules:
- Only place for SQL queries
- Implements `domain/user.Repository`, etc.

Files (Topic 06):
- `user_repo.go` — PostgresUserRepository
