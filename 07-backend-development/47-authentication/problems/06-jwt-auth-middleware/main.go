/*
============================================================
47 — AUTHENTICATION
Problem 6 — JWT Auth Middleware
============================================================

Create:

    GET /profile

Create:

    func jwtAuthMiddleware(tokenService *TokenService) func(http.Handler) http.Handler

Requirements:

    1. Read Authorization header

    2. Expected format:

           Bearer <token>

    3. If header is missing → 401 Unauthorized

    4. If format is wrong → 401 Unauthorized

    5. Parse and validate JWT with TokenService

    6. If token is invalid or expired → 401 Unauthorized

    7. If token type is not "access" → 401 Unauthorized

    8. On success, call next handler

    9. profileHandler returns:

           Profile endpoint

   10. Without token → 401
       With valid access token → 200

   11. Run server on :8080

Test:

    curl http://localhost:8080/profile
        → 401

    # login first to get token, then:
    curl http://localhost:8080/profile \
      -H "Authorization: Bearer <access_token>"
        → 200 Profile endpoint

============================================================
Goal
============================================================

Practice:

    - Bearer token extraction
    - JWT auth middleware
    - protecting routes

============================================================
*/

package main

func main() {}
