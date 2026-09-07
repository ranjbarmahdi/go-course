/*
============================================================
47 — AUTHENTICATION
Problem 3 — Login Endpoint
============================================================

Create:

    POST /register
    POST /login

Login request body (JSON):

    {
        "email": "mahdi@example.com",
        "password": "secret123"
    }

Requirements:

    1. Register must work first (reuse logic from Problem 2)

    2. Login must:

           find user by email
           verify password with bcrypt
           return tokens on success

    3. Return 401 Unauthorized for:

           unknown email
           wrong password

    4. On success return 200 OK:

           {
               "access_token": "<jwt>",
               "token_type": "Bearer"
           }

    5. access_token must be a valid JWT signed with HS256

    6. JWT payload must include:

           user_id
           email
           exp

    7. Use github.com/golang-jwt/jwt/v5

    8. Secret key: "my-secret-key"

    9. Token expires in 15 minutes

   10. Run server on :8080

Test:

    # register first
    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"secret123"}'

    # then login
    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"secret123"}'

============================================================
Goal
============================================================

Practice:

    - login flow
    - bcrypt password verification
    - JWT generation
    - auth response DTO

============================================================
*/

package main

func main() {}
