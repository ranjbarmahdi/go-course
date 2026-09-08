# infra/config

Purpose: Load configuration from environment at startup.

Rules:
- Fail fast if required vars missing
- Use `.env` locally via godotenv (Topic 13)
- Never log secrets

Files:
- `config.go` — flat Config struct, Load(), Validate(), ParseSlogLevel()

Required env vars (see `.env.example`):
- `DATABASE_URL`, `JWT_SECRET`
- Optional defaults: `PORT`, `ENV`, `LOG_LEVEL`
