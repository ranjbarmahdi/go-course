/*
============================================================
06 — VALIDATION
Problem 5 — Register Handler with Validation
============================================================

Create:

    POST /register

Request body:

    {
        "email": "me@example.com",
        "password": "secret123"
    }

Requirements:

    1. RegisterRequest with validate tags

    2. Decode JSON → 400 invalid JSON

    3. validate.Struct → 400 with fields map:

           {
               "error": "validation failed",
               "fields": { "Email": "...", "Password": "..." }
           }

    4. Success → 201:

           { "message": "user registered", "email": "..." }

    5. Do NOT store users — validation only

    6. Run on :8080

Test:

    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"bad","password":"123"}'

    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"me@example.com","password":"secret123"}'

============================================================
Goal
============================================================

Practice:

    - validation in HTTP handler
    - 400 Bad Request response

============================================================
*/

package main

func main() {}
