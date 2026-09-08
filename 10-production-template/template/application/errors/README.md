# application/errors

Purpose: Application-level errors with Kind for HTTP mapping.

Rules:
- Use cases return these errors (not raw HTTP status codes)
- Handlers call `WriteAppError(err)` to map Kind → 400/401/404/500

Files (Topic 10):
- `error.go` — Kind enum, Error struct, KindOf()

Kinds: InvalidInput, Unauthorized, NotFound, Internal
