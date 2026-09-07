/*
============================================================
48 — DATABASE/SQL
Problem 5 — Select One User
============================================================

Find a user by email using QueryRow.

Requirements:

    1. Assume users table exists with at least one row:

           email: mahdi@example.com

       (run Problem 4 first if needed)

    2. Query:

           SELECT id, email, password_hash
           FROM users
           WHERE email = $1

    3. Use db.QueryRow(...).Scan(...)

    4. Print:

           Found: <id> <email>

    5. Create function:

           func findUserByEmail(db *sql.DB, email string) (id, email, passwordHash string, err error)

============================================================
Goal
============================================================

Practice:

    - QueryRow
    - Scan
    - SELECT with WHERE

============================================================
*/

package main

func main() {}
