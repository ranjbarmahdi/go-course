# application/errors

Application-level errors with a `Kind` for HTTP mapping.

## File

- `error.go` — `Kind`, `Error`, `New`, `KindOf`

## Kinds

| Kind | Typical HTTP status |
|---|---|
| `InvalidInput` | 400 |
| `Unauthorized` | 401 |
| `Domain` | 400 |
| `NotFound` | 404 |
| `Conflict` | 409 |
| `Internal` | 500 |

## Usage

Use cases return these errors. HTTP handlers call `utils.WriteAppError(w, err)`.

Domain errors are translated first via `application/utils.TranslateDomainError`.
