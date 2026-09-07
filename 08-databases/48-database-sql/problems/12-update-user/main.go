/*
============================================================
48 — DATABASE/SQL
Problem 12 — Update User Email
============================================================

Update a user's email by id.

Requirements:

    1. Ensure users table exists with user:

           id:    user-1
           email: mahdi@example.com

       (insert manually or via Problem 4 if missing)

    2. Create function:

           func updateUserEmail(db *sql.DB, id, newEmail string) error

    3. Query:

           UPDATE users
           SET email = $1
           WHERE id = $2

    4. Use db.Exec

    5. Update user-1 email to:

           mahdi.updated@example.com

    6. Select user-1 and print new email to confirm

    7. Print:

           Email updated: mahdi.updated@example.com

============================================================
Goal
============================================================

Practice:

    - UPDATE with Exec
    - verifying updates with SELECT

============================================================
*/

package main

func main() {}
