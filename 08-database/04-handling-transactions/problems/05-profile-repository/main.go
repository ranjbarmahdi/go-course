/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 5 — Profile Repository
============================================================

Create ProfileRepository with getExecutor pattern.

Requirements:

    1. Domain Profile: ID, UserID, Name

    2. ProfileRepository interface:

           Create(ctx, profile Profile) error

    3. PostgresProfileRepository with getExecutor

    4. Table:

           CREATE TABLE IF NOT EXISTS profiles (
               id       TEXT PRIMARY KEY,
               user_id  TEXT NOT NULL UNIQUE REFERENCES users(id),
               name     TEXT NOT NULL
           )

    5. Create must use exec from getExecutor(ctx, r.db)

    6. Test insert in main

============================================================
Goal
============================================================

Practice:

    - second repository for multi-repo transactions

============================================================
*/

package main

func main() {}
