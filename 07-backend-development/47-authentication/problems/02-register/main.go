/*
============================================================
47 — AUTHENTICATION
Problem 2 — Register Endpoint
============================================================

Create:

    POST /register

Request body (JSON):

    {
        "email": "mahdi@example.com",
        "password": "secret123"
    }

Requirements:

    1. Create RegisterRequest struct with JSON tags

    2. Decode body with json.NewDecoder

    3. Validate:

           email    → required
           password → required, minimum 8 characters

    4. Hash password with bcrypt before storing

    5. Store users in memory (map or slice)

    6. Return 409 Conflict if email already exists

    7. Return 201 Created on success:

           {
               "message": "user registered",
               "email": "mahdi@example.com"
           }

    8. Return 400 Bad Request for invalid JSON or validation errors

    9. Set Content-Type: application/json

   10. Run server on :8080

    Do NOT use a database.

Test:

    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"secret123"}'

============================================================
Goal
============================================================

Practice:

    - register flow
    - password hashing
    - JSON request DTO
    - validation
    - in-memory user store

============================================================
*/

package main

func main() {}
