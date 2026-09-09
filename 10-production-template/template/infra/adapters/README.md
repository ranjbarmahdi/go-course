# infra/adapters

Implementations of domain and application interfaces.

## Contains

```
adapters/
├── id-generator.go         UUIDGenerator (google/uuid)
└── repository/
    └── sample-repository.go
```

## Rules

- SQL and external SDK calls live here only
- Map DB rows ↔ domain entities
- Translate driver errors at the boundary (`sql.ErrNoRows` → domain `ErrNotFound`)
- Use `postgres.Executor(ctx, db)` in every query

## Adding an adapter

1. Implement the domain repository interface in `repository/`
2. Add compile-time check: `var _ domain.XRepository = (*XRepository)(nil)`
3. Register in `infra/bootstrap/sets.go` with `wire.Bind`

Future: Kafka producer, email sender, etc. as separate subfolders implementing `application/contracts` ports.
