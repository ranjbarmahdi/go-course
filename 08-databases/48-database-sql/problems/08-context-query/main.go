/*
============================================================
48 — DATABASE/SQL
Problem 8 — Context Query
============================================================

Query a user with a context timeout.

Requirements:

    1. Create function:

           func findUserByEmailContext(
               ctx context.Context,
               db *sql.DB,
               email string,
           ) (id, email string, err error)

    2. Use db.QueryRowContext (NOT QueryRow)

    3. In main():

           ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
           defer cancel()

    4. Query email: mahdi@example.com

    5. Print:

           Context query OK: <email>

    6. Handle sql.ErrNoRows if user does not exist

============================================================
Goal
============================================================

Practice:

    - QueryRowContext
    - context.WithTimeout
    - production-style queries

============================================================
*/

package main

func main() {}
