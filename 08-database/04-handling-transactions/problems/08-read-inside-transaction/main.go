/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 8 — Read Inside Transaction
============================================================

Read and write in the same transaction.

Requirements:

    1. Inside RunInTx:

           user, err := userRepo.FindByEmail(ctx, email)
           if not found → create user
           profileRepo.Create(ctx, profile)

    2. FindByEmail must use getExecutor (same tx)

    3. Create user only if email does not exist

    4. If email exists → return ErrEmailAlreadyExists

    5. Test:

           first run  → creates user + profile
           second run → ErrEmailAlreadyExists, no partial data

============================================================
Goal
============================================================

Practice:

    - SELECT + INSERT in same transaction
    - getExecutor for reads and writes

============================================================
*/

package main

func main() {}
