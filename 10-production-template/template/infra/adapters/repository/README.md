# infra/adapters/repository

PostgreSQL implementations of domain repository interfaces.

## Current

- `sample-repository.go` — `CreateSample`, `FindByID` for `domain/sample`

## Rules

- Only place for SQL queries
- Always use `postgres.Executor(ctx, r.db)` — supports ambient transactions
- Reconstruct entities with domain constructors (`NewSample`, not raw struct literals)
- Map `sql.ErrNoRows` to domain sentinel errors

## Example pattern

```go
func NewSampleRepository(db *sql.DB) *SampleRepository { ... }

var _ domainsample.SampleRepository = (*SampleRepository)(nil)

func (r *SampleRepository) FindByID(ctx context.Context, id valueobjects.SampleID) (*domainsample.Sample, error) {
    exec := postgres.Executor(ctx, r.db)
    // SELECT ... Scan ... NewSample(...)
}
```

When adding a feature, create `<feature>-repository.go` and bind it in Wire.
