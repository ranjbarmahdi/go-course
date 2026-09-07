/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 7 — Transaction Rollback
============================================================

Prove rollback works when second insert fails.

Requirements:

    1. Use RunInTx with userRepo + profileRepo

    2. Inside RunInTx:

           userRepo.Create(ctx, user)     → succeeds
           profileRepo.Create(ctx, profile with DUPLICATE user_id or invalid FK)

    3. Transaction must rollback

    4. After RunInTx returns error:

           user must NOT exist in database

    5. Print:

           Transaction rolled back — user not saved

    6. Use profiles.user_id UNIQUE REFERENCES users(id)

============================================================
Goal
============================================================

Practice:

    - atomic rollback
    - why RunInTx matters

============================================================
*/

package main

func main() {}
