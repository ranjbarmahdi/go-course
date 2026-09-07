/*
============================================================
50 — REPOSITORY PATTERN
Problem 5 — Map ErrNoRows to Domain Error
============================================================

Create GetUserUseCase that maps repository errors to domain errors.

Requirements:

    1. GetUserUseCase depends on UserRepository

    2. Execute(ctx, email) (User, error)

    3. If repository returns sql.ErrNoRows:

           return User{}, ErrUserNotFound

    4. Other repository errors:

           wrap and return

    5. Test in main:

           existing email → returns ErrUserNotFound? no — returns user
           missing email    → ErrUserNotFound

    6. Use errors.Is for sql.ErrNoRows

============================================================
Goal
============================================================

Practice:

    - error mapping in use case
    - sql.ErrNoRows → domain error

============================================================
Layer: APPLICATION
============================================================
*/

package main

func main() {}
