/*
============================================================
48 — DATABASE/SQL
Problem 3 — Create Table
============================================================

Connect to PostgreSQL and create a users table.

Requirements:

    1. Connect and ping (same DSN as Problem 1)

    2. Run CREATE TABLE with db.Exec:

           CREATE TABLE IF NOT EXISTS users (
               id            TEXT PRIMARY KEY,
               email         TEXT NOT NULL UNIQUE,
               password_hash TEXT NOT NULL,
               created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
           )

    3. Print:

           Table users created

    4. Running the program twice must NOT fail (use IF NOT EXISTS)

    5. defer db.Close()

============================================================
Goal
============================================================

Practice:

    - db.Exec for DDL
    - CREATE TABLE
    - PostgreSQL column types

============================================================
*/

package main

func main() {}
