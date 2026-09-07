/*
============================================================
06 — VALIDATION
Problem 11 — Validation vs Business Rules
============================================================

Create POST /register with BOTH validation and business logic.

Requirements:

    1. validator tags for email format and password min=8

    2. In-memory map of registered emails

    3. If email already exists → 409 Conflict:

           { "error": "email already exists" }

    4. If validation fails → 400 Bad Request

    5. If JSON invalid → 400 invalid JSON

    6. Success → 201 Created

    7. Order must be:

           decode → validate → business rule → success

Test:

    # first register — 201
    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"me@example.com","password":"secret123"}'

    # duplicate — 409
    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"me@example.com","password":"secret123"}'

    # bad email — 400 (not 409)
    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"bad","password":"secret123"}'

============================================================
Goal
============================================================

Practice:

    - validation (400) vs business rules (409)
    - correct handler order

============================================================
*/

package main

func main() {}
