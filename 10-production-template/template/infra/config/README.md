# infra/config

Load and validate configuration from environment at startup.

## File

- `config.go` — `Config`, `Load()`, `Validate()`, `ParseSlogLevel()`

## Required variables

| Variable | Notes |
|---|---|
| `DATABASE_URL` | PostgreSQL connection string |
| `DB_SCHEMA` | Schema name; must exist before boot |

## Common variables

| Variable | Default | Notes |
|---|---|---|
| `PORT` | `8080` | HTTP listen port |
| `ENV` | `development` | Controls Redis mode |
| `LOG_LEVEL` | `info` | debug, info, warn, error |
| `AUTO_MIGRATE` | `false` | Run goose on startup when true |

## Redis

**Development / self-host** (`ENV=development` or `self-host`):

- Requires `REDIS_URL`

**Production / staging** (sentinel mode):

- Requires `REDIS_SENTINEL_HOST`, `REDIS_SENTINEL_MASTER`, `REDIS_SENTINEL_PASS`
- Optional: `REDIS_SENTINEL_PORT` (default 26379), `REDIS_SENTINEL_USER`

## Rules

- Fail fast on missing or invalid config
- Use `.env` locally via godotenv
- Never log secrets

See `.env.example` for all keys.
