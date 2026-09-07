/*
============================================================
48 — DATABASE/SQL
Problem 1 — Connect and Ping
============================================================

Connect to PostgreSQL and verify the connection.

Requirements:

    1. Use database/sql

    2. Use pgx driver:

           _ "github.com/jackc/pgx/v5/stdlib"

    3. DSN:

           postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

    4. Open with sql.Open("pgx", dsn)

    5. Call db.Ping()

    6. Print:

           Connected to PostgreSQL

    7. defer db.Close()

    8. Exit with log.Fatal on error

Prerequisite:

    Docker PostgreSQL must be running on port 5432.

============================================================
Goal
============================================================

Practice:

    - sql.Open
    - db.Ping
    - pgx driver setup

============================================================
*/

package main

func main() {}
