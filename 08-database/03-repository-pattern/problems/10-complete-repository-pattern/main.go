/*
============================================================
50 — REPOSITORY PATTERN
Problem 10 — Complete Repository Pattern
============================================================

Build a complete layered user registration flow.

Layers required:

    DOMAIN
        User struct
        ErrInvalidInput, ErrEmailAlreadyExists, ErrUserNotFound

    APPLICATION — PORT
        UserRepository interface

    APPLICATION — USE CASE
        RegisterUserUseCase
        GetUserUseCase

    INFRASTRUCTURE
        PostgresUserRepository

    PRESENTATION
        registerHandler (simulated HTTP, print status)

    COMPOSITION ROOT
        main() wires everything

Requirements:

    1. Register mahdi@example.com / secret123 → HTTP 201

    2. Register same email again           → HTTP 409

    3. Get user by email → print id + email

    4. Label each section with layer comment

    5. DSN:

           postgres://postgres:secret@localhost:5432/postgres?sslmode=disable

    6. No SQL outside repository

============================================================
Goal
============================================================

Practice:

    - complete repository pattern
    - all layers together

============================================================
*/

package main

func main() {}
