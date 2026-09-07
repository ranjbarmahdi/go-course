/*
============================================================
51 — HANDLING TRANSACTIONS IN GO
Problem 12 — Complete Transaction Flow
============================================================

Build the complete transaction pattern from lecture 51.

Requirements:

    DOMAIN
        User, Profile
        ErrInvalidInput, ErrEmailAlreadyExists

    APPLICATION — PORT
        UserRepository, ProfileRepository, TransactionManager

    APPLICATION — USE CASE
        CreateUserWithProfileUseCase

    INFRASTRUCTURE
        DBTX, getExecutor, txContextKey
        PostgresTransactionManager
        PostgresUserRepository (getExecutor)
        PostgresProfileRepository (getExecutor)

    PRESENTATION
        handler → print HTTP status

    COMPOSITION ROOT
        main() wiring

    Flow:

        1. Create mahdi@example.com + profile "Mahdi" → HTTP 201
        2. Verify users count = 1 and profiles count = 1
        3. Duplicate email → HTTP 409, counts unchanged

    DSN:

        postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

============================================================
Goal
============================================================

Practice:

    - complete context-propagation transaction pattern
    - production-style use case with RunInTx

============================================================
*/

package main

func main() {}
