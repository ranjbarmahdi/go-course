/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 3 — PostgresTransactionManager
============================================================

Implement TransactionManager.

Requirements:

    1. PostgresTransactionManager struct { db *sql.DB }

    2. NewPostgresTransactionManager(db)

    3. RunInTx must:

           BeginTx
           ctx = context.WithValue(ctx, txKey, tx)
           call fn(txCtx)
           on error → Rollback
           on success → Commit

    4. defer tx.Rollback() after BeginTx

    5. In main(), test RunInTx:

           success fn → prints "committed"
           failing fn → prints "rolled back"

    6. Use getExecutor from Problem 1 pattern

============================================================
Goal
============================================================

Practice:

    - RunInTx implementation
    - BeginTx / Commit / Rollback hidden in infrastructure

============================================================
Layer: INFRASTRUCTURE
============================================================
*/

package main

func main() {}
