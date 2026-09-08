# infra/runtime

Purpose: Run long-lived processes (HTTP server, future Kafka consumer).

Rules:
- `App.Run()` starts all Runners with errgroup
- Graceful shutdown on SIGINT/SIGTERM

Files (Topic 08):
- `runtime.go` — Runner and Indicator interfaces
- `app.go` — App runs []Runner

Today: HTTP server is the only Runner.
