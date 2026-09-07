/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 2 — TransactionManager Port
============================================================

Define the application port for transactions.

Requirements:

    1. TransactionManager interface:

           RunInTx(ctx context.Context, fn func(ctx context.Context) error) error

    2. Label:

           // LAYER: APPLICATION — PORT

    3. Do NOT implement yet

    4. Add comment explaining:

           use case calls RunInTx
           use case never calls BeginTx or Commit

============================================================
Goal
============================================================

Practice:

    - transaction port interface
    - application layer abstraction

============================================================
*/

package main

func main() {}
