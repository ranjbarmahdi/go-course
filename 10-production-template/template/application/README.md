# application

Use cases, application errors, and shared ports.

## Contains

```
application/
├── contracts/          IDGenerator, RunInTx
├── errors/             apperrors Kind + KindOf
├── shared/             pagination (Page, PageResult)
├── utils/error.go      TranslateDomainError
└── usecase/
    └── sample/         reference vertical slice
        ├── response.go, mapper.go
        ├── create-sample/
        └── get-sample/
```

## Rules

- Import `domain/` only — not infra, not `net/http`
- Business rules live in use cases, not in HTTP validator tags
- Use cases return `application/errors` — handlers map them with `WriteAppError`
- Call `TranslateDomainError` when domain errors surface from entities or repos

## Adding a use case

1. Create `application/usecase/<feature>/<action>/`
2. Define `UseCase` interface + `Implementation` + `New(...)` constructor
3. Input struct with **no** `json` or `validate` tags
4. Register in `infra/bootstrap/sets.go` with `wire.Bind`

See `create-sample/` and `get-sample/` as examples.
