/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 1 — DBTX and getExecutor
============================================================

Create DBTX interface and getExecutor helper.

Requirements:

    1. DBTX interface with:

           ExecContext
           QueryRowContext
           QueryContext

    2. txContextKey struct{} for context key

    3. getExecutor(ctx, db *sql.DB) DBTX

           if ctx has *sql.Tx → return tx
           else               → return db

    4. In main(), demonstrate:

           without tx in ctx → getExecutor returns db
           (manual test with BeginTx + WithValue optional)

    5. Label code:

           // LAYER: INFRASTRUCTURE

============================================================
Goal
============================================================

Practice:

    - DBTX abstraction
    - reading tx from context

============================================================
*/

package main

func main() {}
