/*
============================================================
50 — REPOSITORY PATTERN
Problem 4 — Register User Use Case
============================================================

Create RegisterUserUseCase (single repository, no transaction).

Requirements:

    1. RegisterUserUseCase depends on UserRepository interface

    2. Constructor:

           NewRegisterUserUseCase(users UserRepository)

    3. Execute(ctx, email, password) error

    4. Business rules:

           email and password required
           password min 8 characters → ErrInvalidInput
           email already exists      → ErrEmailAlreadyExists

    5. Hash password (simple prefix "hashed:" is OK for practice)

    6. Call users.Create(ctx, user)

    7. NO SQL in use case

    8. Wire with PostgresUserRepository in main and test

============================================================
Goal
============================================================

Practice:

    - use case layer
    - business logic without SQL

============================================================
Layer: APPLICATION
============================================================
*/

package main

func main() {}
