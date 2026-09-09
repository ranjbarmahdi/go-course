# infra/runtime

Run long-lived processes and coordinate graceful shutdown.

## Files

| File | Purpose |
|---|---|
| `runner.go` | `Runner` and `Indicator` interfaces |
| `app.go` | `App` — starts all runners with errgroup |

## How it works

1. Wire builds `[]Runner` (today: HTTP server only)
2. `App.Run(ctx)` starts each runner in errgroup
3. Context cancellation triggers graceful shutdown
4. HTTP server handles its own `Shutdown` on `ctx.Done()`

## Indicator

Database packages implement `Indicator` for `/readyz`:

```go
type Indicator interface {
    Name() string
    Ready(ctx context.Context) bool
}
```

## Adding a runner

Implement `Runner` (e.g. Kafka consumer), add to `provideRunners` in bootstrap.
