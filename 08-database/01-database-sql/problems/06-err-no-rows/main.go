/*
============================================================
48 — DATABASE/SQL
Problem 6 — Handle sql.ErrNoRows
============================================================

Handle "user not found" separately from database errors.

Requirements:

    1. Create function:

           func findUserByEmail(db *sql.DB, email string) (string, string, error)

       Returns: id, email, error

    2. Query user by email (same as Problem 5)

    3. In main(), test two cases:

       Case A — existing email:
           mahdi@example.com
           → print: Found: <id> <email>

       Case B — missing email:
           unknown@example.com
           → print: User not found
           → do NOT log.Fatal

    4. Use:

           if err == sql.ErrNoRows { ... }

    5. Real database errors must still use log.Fatal

============================================================
Goal
============================================================

Practice:

    - sql.ErrNoRows
    - not found vs database error

============================================================
*/

package main

func main() {}
