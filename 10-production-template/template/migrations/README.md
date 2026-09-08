# migrations

Purpose: Versioned SQL schema changes.

Rules:
- One file per migration: `001_create_users.sql`, `002_create_products.sql`
- Never edit applied migrations — add new files
- Run with migrate tool or Makefile (Topic 13)

Files (Topic 06+):
- `001_create_users.sql`
