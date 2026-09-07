/*
============================================================
50 — REPOSITORY PATTERN
Problem 3 — Postgres User Repository
============================================================

Implement UserRepository with PostgreSQL.

Requirements:

    1. PostgresUserRepository struct with *sql.DB

    2. NewPostgresUserRepository(db *sql.DB)

    3. Implement Create and FindByEmail

    4. Ensure users table exists:

           id TEXT PRIMARY KEY
           email TEXT UNIQUE
           password_hash TEXT
           created_at TIMESTAMPTZ DEFAULT NOW()

    5. DSN:

           postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

    6. In main():

           insert user mahdi@example.com
           find by email and print id + email

    7. SQL only inside repository — not in main

============================================================
Goal
============================================================

Practice:

    - infrastructure layer
    - implementing port interface

============================================================
Layer: INFRASTRUCTURE
============================================================
*/

package main

func main() {}
