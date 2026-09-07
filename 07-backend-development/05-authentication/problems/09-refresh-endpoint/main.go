/*
============================================================
47 — AUTHENTICATION
Problem 9 — Refresh Endpoint
============================================================

Create:

    POST /login
    POST /refresh

Refresh request body (JSON):

    {
        "refresh_token": "<refresh_token>"
    }

Requirements:

    1. Login returns access_token and refresh_token

    2. POST /refresh accepts refresh_token in JSON body

    3. Validate refresh token:

           must be valid JWT
           type must be "refresh"
           must not be expired

    4. On success return 200 OK:

           {
               "access_token": "<new_access_token>",
               "token_type":   "Bearer"
           }

    5. Return 401 if refresh token is invalid or expired

    6. Return 400 if refresh_token field is missing

    7. Do NOT accept access_token on /refresh (type must be "refresh")

    8. Run server on :8080

Test:

    # login
    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"secret123"}'

    # refresh
    curl -X POST http://localhost:8080/refresh \
      -H "Content-Type: application/json" \
      -d '{"refresh_token":"<refresh_token>"}'

============================================================
Goal
============================================================

Practice:

    - refresh token flow
    - issuing new access tokens
    - token type validation

============================================================
*/

package main

func main() {}
