/*
============================================================
50 — REPOSITORY PATTERN
Problem 7 — Composition Root Wiring
============================================================

Wire all layers in main (composition root).

Requirements:

    1. openDB + ensure users table

    2. Create:

           userRepo := NewPostgresUserRepository(db)
           registerUC := NewRegisterUserUseCase(userRepo)
           getUserUC    := NewGetUserUseCase(userRepo)

    3. Register user: wire@example.com / secret123

    4. Get user by email and print result

    5. main() is the ONLY place that knows PostgresUserRepository

    6. Use cases receive UserRepository interface

    7. DSN:

           postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

============================================================
Goal
============================================================

Practice:

    - composition root
    - dependency injection

============================================================
Layer: COMPOSITION ROOT
============================================================
*/

package main

func main() {}
