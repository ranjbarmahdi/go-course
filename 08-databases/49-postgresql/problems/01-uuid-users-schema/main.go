/*
============================================================
49 — POSTGRESQL
Problem 1 — UUID Users Schema
============================================================

Create a production-style users table with UUID primary key.

Requirements:

    1. Connect to PostgreSQL (pgx driver)

    2. DSN:

           postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

    3. Enable extension:

           CREATE EXTENSION IF NOT EXISTS "pgcrypto";

    4. Create table:

           CREATE TABLE IF NOT EXISTS users (
               id            UUID PRIMARY KEY DEFAULT gen_random_uuid(),
               email         VARCHAR(255) NOT NULL UNIQUE,
               password_hash TEXT NOT NULL,
               created_at    TIMESTAMPTZ NOT NULL DEFAULT NOW(),
               updated_at    TIMESTAMPTZ NOT NULL DEFAULT NOW()
           )

    5. Print:

           Users table created

    6. Use IF NOT EXISTS (safe to run twice)

============================================================
Goal
============================================================

Practice:

    - UUID primary keys
    - PostgreSQL data types
    - schema design

============================================================
*/

package main

func main() {}
