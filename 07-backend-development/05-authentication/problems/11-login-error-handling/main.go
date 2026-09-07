/*
============================================================
47 — AUTHENTICATION
Problem 11 — Login Error Handling
============================================================

Create:

    POST /register
    POST /login

Requirements:

    1. Register works as in Problem 2

    2. Login returns JSON errors with Content-Type: application/json

    3. Unknown email:

           HTTP 401
           {"error": "invalid email or password"}

    4. Wrong password:

           HTTP 401
           {"error": "invalid email or password"}

       (same message for both — do not reveal which failed)

    5. Invalid JSON body:

           HTTP 400
           {"error": "invalid JSON"}

    6. Missing email or password:

           HTTP 400
           {"error": "email and password are required"}

    7. Successful login:

           HTTP 200
           {"access_token": "...", "refresh_token": "...", "token_type": "Bearer"}

    8. Run server on :8080

Test:

    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"email":"unknown@example.com","password":"secret123"}'

    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"wrongpassword"}'

============================================================
Goal
============================================================

Practice:

    - secure login error messages
    - consistent JSON error responses
    - validation

============================================================
*/

package main

func main() {}
