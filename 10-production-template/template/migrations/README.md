# migrations

Versioned SQL schema changes, applied with [goose](https://github.com/pressly/goose).

## Naming

```
<version>_<name>.sql
```

Examples:

- `00001_sample_create.sql`
- `00002_add_index_on_sample_uuid.sql`

Use lowercase `.sql` extension and an underscore after the version number.

## File format

Up and Down in the **same file**:

```sql
-- +goose Up
CREATE TABLE ...;

-- +goose Down
DROP TABLE ...;
```

## Schema

- Operator creates the schema: `CREATE SCHEMA <DB_SCHEMA>;`
- App sets `search_path` via `DATABASE_URL` + `DB_SCHEMA`
- Table names in SQL are unqualified (no schema prefix in migrations)

## Commands

```bash
go run ./cmd/migrate status
go run ./cmd/migrate up
go run ./cmd/migrate down
```

Or set `AUTO_MIGRATE=true` to run `up` on API startup.

## Rules

- Never edit an applied migration — add a new file
- Goose tracks history in `goose_db_version`

## Current migrations

| File | Creates |
|---|---|
| `00001_sample_create.sql` | `sample_table` |
