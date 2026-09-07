/*
============================================================
48 — DATABASE/SQL
Problem 13 — Delete User
============================================================

Delete a user by id.

Requirements:

    1. Ensure users table exists

    2. Insert test user if not exists:

           id:    user-delete
           email: delete@example.com

    3. Create function:

           func deleteUserByID(db *sql.DB, id string) error

    4. Query:

           DELETE FROM users WHERE id = $1

    5. Use db.Exec

    6. After delete, FindByEmail must return sql.ErrNoRows

    7. Print:

           User deleted: user-delete

============================================================
Goal
============================================================

Practice:

    - DELETE with Exec
    - confirming deletion with SELECT

============================================================
*/

package main

func main() {}
