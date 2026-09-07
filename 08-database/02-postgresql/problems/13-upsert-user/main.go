/*
============================================================
49 — POSTGRESQL
Problem 13 — Upsert User
============================================================

Insert user or get existing on duplicate email.

Requirements:

    1. Insert user:

           email: shop@example.com
           password_hash: $2a$10$demo

    2. Use ON CONFLICT:

           INSERT INTO users (email, password_hash)
           VALUES ($1, $2)
           ON CONFLICT (email) DO UPDATE SET email = EXCLUDED.email
           RETURNING id, email

    3. Run twice — must return same user (no duplicate error)

    4. Print both times:

           User ID: <id>  Email: <email>

============================================================
Goal
============================================================

Practice:

    - ON CONFLICT (upsert)
    - idempotent seed data

============================================================
*/

package main

func main() {}
