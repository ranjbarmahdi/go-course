/*
============================================================
49 — POSTGRESQL
Problem 2 — Products With CHECK Constraint
============================================================

Create a products table with a price validation constraint.

Requirements:

    1. Connect to PostgreSQL (same DSN)

    2. Create table:

           CREATE TABLE IF NOT EXISTS products (
               id          UUID PRIMARY KEY DEFAULT gen_random_uuid(),
               name        VARCHAR(255) NOT NULL,
               price       NUMERIC(10, 2) NOT NULL CHECK (price > 0),
               created_at  TIMESTAMPTZ NOT NULL DEFAULT NOW()
           )

    3. Insert valid product:

           name:  Go Book
           price: 29.99

    4. Print:

           Product inserted: Go Book

    5. Attempt insert with price 0 or negative must fail

       (test in psql or print note — do not crash program)

============================================================
Goal
============================================================

Practice:

    - NUMERIC for money
    - CHECK constraints
    - never use FLOAT for money

============================================================
*/

package main

func main() {}
