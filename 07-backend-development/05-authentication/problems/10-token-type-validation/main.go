/*
============================================================
47 — AUTHENTICATION
Problem 10 — Reject Access Token on Refresh
============================================================

Create:

    POST /login
    POST /refresh

Requirements:

    1. Login returns access_token (type: "access") and refresh_token (type: "refresh")

    2. POST /refresh must REJECT access_token:

           if token type is "access" → 401
           body: {"error": "refresh token required"}

    3. POST /refresh must ACCEPT refresh_token:

           return new access_token

    4. jwtAuthMiddleware on GET /profile must REJECT refresh_token:

           if token type is "refresh" → 401
           body: {"error": "access token required"}

    5. GET /profile accepts only access_token

    6. Run server on :8080

Test:

    # use access_token on /refresh → 401
    curl -X POST http://localhost:8080/refresh \
      -H "Content-Type: application/json" \
      -d '{"refresh_token":"<access_token>"}'

    # use refresh_token on /profile → 401
    curl http://localhost:8080/profile \
      -H "Authorization: Bearer <refresh_token>"

============================================================
Goal
============================================================

Practice:

    - access vs refresh token separation
    - type claim validation
    - correct token for each endpoint

============================================================
*/

package main

func main() {}
