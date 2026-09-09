# infra/httpserver/user

**Placeholder** for a future user/auth feature.

## Reference implementation

Use **`routes/sample/`** as the template for new HTTP features:

```
routes/sample/
├── routes.go       mount + Register calls
├── handler.go      decode, validate, use case, write
├── requests.go     HTTP DTOs (json + validate tags)
└── responses.go    HTTP DTOs + ToApiResponse mappers
```

## Planned user routes (not implemented)

- `POST /api/v2/register`
- `POST /api/v2/login`

When implementing, follow the same pattern as sample and register in `router.go` + Wire.
