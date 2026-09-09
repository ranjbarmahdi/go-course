# infra

Technical implementations — HTTP, databases, config, wiring, runtime.

## Contains

```
infra/
├── adapters/       repository implementations, UUID generator
├── bootstrap/      Google Wire composition root
├── config/         environment loading
├── database/       postgres, redis, migrate
├── httpserver/     HTTP server, routes, middleware
└── runtime/        App, Runner, graceful shutdown
```

## Rules

- Implements interfaces from `domain/` and `application/contracts/`
- Only `bootstrap/` and `cmd/` wire concrete types together
- Repositories use `postgres.Executor(ctx, db)` — never bypass for transactions
