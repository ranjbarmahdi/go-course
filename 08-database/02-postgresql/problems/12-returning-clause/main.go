/*
============================================================
49 — POSTGRESQL
Problem 12 — RETURNING Clause
============================================================

Use PostgreSQL RETURNING to get inserted row data.

Requirements:

    1. Insert user:

           INSERT INTO users (email, password_hash)
           VALUES ($1, $2)
           RETURNING id, email, created_at

    2. Scan returned values into Go variables

    3. Print:

           User created: <id> <email> <created_at>

    4. Use QueryRowContext (not Exec) for RETURNING

    5. Same for product insert with RETURNING id, name

============================================================
Goal
============================================================

Practice:

    - RETURNING clause
    - avoid extra SELECT after INSERT

============================================================
*/

package main

func main() {}
