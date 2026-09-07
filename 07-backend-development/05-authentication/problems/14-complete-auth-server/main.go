/*
============================================================
47 — AUTHENTICATION
Problem 14 — Complete Auth Server
============================================================

Create a complete authentication API.

Routes:

    POST /register     (public)
    POST /login        (public)
    POST /refresh      (public)
    GET  /profile      (protected)

============================================================
REQUIREMENTS
============================================================

    1. Register:

           accept email + password
           hash password with bcrypt
           validate password (min 8 chars)
           return 201 on success
           return 409 if email exists

    2. Login:

           verify email + password
           return access_token + refresh_token + token_type
           return 401 for invalid credentials

    3. Refresh:

           accept refresh_token in JSON body
           validate type == "refresh"
           return new access_token

    4. Profile (protected):

           require Bearer access token
           return user_id and email from JWT context

    5. Use in-memory user store

    6. Use TokenService for all JWT logic

    7. Secret: "super-secret-key"

    8. Access token:  15 minutes
       Refresh token: 7 days

    9. All JSON responses must set Content-Type: application/json

   10. Run server on :8080

   11. Allowed packages:

           net/http
           encoding/json
           golang.org/x/crypto/bcrypt
           github.com/golang-jwt/jwt/v5

       Do NOT use a database.

============================================================
TEST
============================================================

    # 1. Register
    curl -X POST http://localhost:8080/register \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"secret123"}'

    # 2. Login
    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"secret123"}'

    # 3. Profile
    curl http://localhost:8080/profile \
      -H "Authorization: Bearer <access_token>"

    # 4. Refresh
    curl -X POST http://localhost:8080/refresh \
      -H "Content-Type: application/json" \
      -d '{"refresh_token":"<refresh_token>"}'

    # 5. Profile without token → 401
    curl http://localhost:8080/profile

============================================================
Goal
============================================================

Practice:

    - complete authentication flow
    - register, login, refresh, protected routes
    - bcrypt + JWT together
    - production-style auth API

============================================================
*/

package main

func main() {}
