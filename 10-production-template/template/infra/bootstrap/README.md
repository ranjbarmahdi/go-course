# infra/bootstrap

Dependency injection with [Google Wire](https://github.com/google/wire).

## Files

| File | Purpose |
|---|---|
| `wire.go` | `wire.Build(...)` injector (build tag: `wireinject`) |
| `sets.go` | Provider sets (Database, Adapter, Sample, Http, Runtime) |
| `providers.go` | Small provider functions |
| `wire_gen.go` | Generated — do not edit by hand |

## Regenerate

```bash
cd infra/bootstrap
go generate
```

Run this after changing `sets.go`, `providers.go`, or constructor signatures.

## Current graph (simplified)

```
Config → Postgres, Redis
       → SampleRepository, RunInTx, IDGenerator
       → create-sample, get-sample use cases
       → sample HTTP handler
       → router → server → App
```

## Rules

- Add new features as Wire sets (repo + use case + handler)
- Use `wire.Bind` to connect interfaces to implementations
- Unused providers in `wire.Build` cause Wire errors — only include what the graph needs
