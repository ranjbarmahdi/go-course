/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 4 — Repository With getExecutor
============================================================

Create PostgresUserRepository that uses getExecutor.

Requirements:

    1. Create(ctx, user) uses getExecutor(ctx, r.db)

    2. FindByEmail(ctx, email) uses getExecutor(ctx, r.db)

    3. Works both:

           outside RunInTx → uses *sql.DB
           inside RunInTx  → uses *sql.Tx from context

    4. Ensure users table exists

    5. Test in main():

           a) Create user WITHOUT RunInTx
           b) Create user INSIDE RunInTx

    6. Both must persist to database

============================================================
Goal
============================================================

Practice:

    - context-aware repository
    - same repo inside/outside transaction

============================================================
Layer: INFRASTRUCTURE
============================================================
*/

package main

func main() {}
