/*
============================================================
48 — DATABASE/SQL
Problem 4 — Insert User
============================================================

Insert a user into the users table.

Requirements:

    1. Ensure users table exists (CREATE TABLE IF NOT EXISTS)

    2. Insert with db.Exec:

           INSERT INTO users (id, email, password_hash)
           VALUES ($1, $2, $3)

    3. Insert this user:

           id:            user-1
           email:         mahdi@example.com
           password_hash: $2a$10$hashedpasswordhere

    4. Print:

           User inserted: mahdi@example.com

    5. Use PostgreSQL placeholders ($1, $2, $3)

    6. Handle errors with log.Fatal

DSN:

    postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

Note:

    If you run twice, you may get a duplicate email error.
    That is expected — topic 49 covers handling this properly.

============================================================
Goal
============================================================

Practice:

    - INSERT with Exec
    - parameterized queries

============================================================
*/

package main

func main() {}
