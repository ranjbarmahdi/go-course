/*
============================================================
49 — POSTGRESQL
Problem 5 — Run Migrations
============================================================

Apply schema migrations from embedded SQL strings.

Requirements:

    1. Create function:

           func runMigrations(db *sql.DB) error

    2. Apply in order:

           001 — users (UUID schema)
           002 — products
           003 — orders
           004 — order_items

    3. Print after each:

           Migration applied: 001_create_users
           Migration applied: 002_create_products
           ...

    4. Use CREATE TABLE IF NOT EXISTS

    5. Running twice must NOT fail

    6. Use same DSN as lecture

Reference SQL is in:

    08-databases/49-postgresql/migrations/

============================================================
Goal
============================================================

Practice:

    - migration pattern
    - ordered schema setup
    - idempotent migrations (IF NOT EXISTS)

============================================================
*/

package main

func main() {}
