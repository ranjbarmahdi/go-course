# cmd/api

Purpose: Application entry point.

Rules:
- Load config, setup slog, handle signals
- Wire dependencies (manual Topic 07, Wire Topic 11)
- Start `runtime.App`

Files:
- `main.go`

Topic 07 adds full wiring: config → db → repo → use case → handler → server
