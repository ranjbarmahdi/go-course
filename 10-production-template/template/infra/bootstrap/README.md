# infra/bootstrap

Purpose: Dependency injection with Google Wire (Topic 11).

Rules:
- `wire_gen.go` is generated — do not edit by hand
- Run `make wire` or `go generate` after changing sets

Files (Topic 11):
- `wire.go`, `sets.go`, `providers.go`, `wire_gen.go`

Before Wire (Topic 07): wire manually in `cmd/api/main.go`.
