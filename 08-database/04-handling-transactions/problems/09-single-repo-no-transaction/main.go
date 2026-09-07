/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 9 — Single Repo Without RunInTx
============================================================

Demonstrate when NOT to use RunInTx.

Requirements:

    1. GetProfileUseCase with only ProfileRepository (no TransactionManager)

    2. ProfileRepository.FindByUserID(ctx, userID) (Profile, error)

    3. Single read — no RunInTx

    4. Repo uses getExecutor — works with normal ctx (no tx in context)

    5. In main():

           seed user + profile (from previous problems or inline)
           get profile by user_id
           print profile name

    6. Add comment:

           // Single-table / single-repo → no RunInTx needed

============================================================
Goal
============================================================

Practice:

    - when to skip transactions

============================================================
*/

package main

func main() {}
