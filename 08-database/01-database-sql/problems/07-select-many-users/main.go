/*
============================================================
48 — DATABASE/SQL
Problem 7 — Select Many Users
============================================================

List all users ordered by created_at.

Requirements:

    1. Assume users table has at least one row

    2. Query:

           SELECT id, email, created_at
           FROM users
           ORDER BY created_at

    3. Use db.Query (not QueryRow)

    4. Loop with rows.Next() and Scan

    5. defer rows.Close()

    6. Check rows.Err() after the loop

    7. Print each user:

           <id>  <email>  <created_at>

    8. Create function:

           func listUsers(db *sql.DB) error

============================================================
Goal
============================================================

Practice:

    - Query
    - rows.Next loop
    - rows.Err

============================================================
*/

package main

func main() {}
