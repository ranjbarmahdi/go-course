/*
============================================================
50 — REPOSITORY PATTERN
Problem 8 — HTTP Error Mapping
============================================================

Map domain errors to HTTP status codes in a handler.

Requirements:

    1. Create registerHTTPHandler(uc *RegisterUserUseCase)

    2. Handler simulates POST /register (hardcoded input is OK)

    3. Map errors:

           ErrInvalidInput       → 400
           ErrEmailAlreadyExists → 409
           other errors          → 500
           success               → 201

    4. Print status code as text:

           HTTP 201: user registered
           HTTP 409: email already exists

    5. Handler must NOT contain SQL

    6. Wire with real PostgresUserRepository in main

============================================================
Goal
============================================================

Practice:

    - presentation layer
    - error mapping at HTTP boundary

============================================================
Layer: PRESENTATION
============================================================
*/

package main

func main() {}
