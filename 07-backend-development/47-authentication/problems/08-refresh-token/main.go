/*
============================================================
47 — AUTHENTICATION
Problem 8 — Refresh Token
============================================================

Extend TokenService with:

    func (s *TokenService) GenerateRefreshToken(userID, email string) (string, error)

Requirements:

    1. Refresh token is also a JWT (HS256)

    2. Claims must include:

           user_id
           email
           type  → "refresh"
           exp   → 7 days from now

    3. Update login to return both tokens:

           {
               "access_token":  "<jwt>",
               "refresh_token": "<jwt>",
               "token_type":    "Bearer"
           }

    4. Access token expires in 15 minutes

    5. Refresh token expires in 7 days

    6. Run server on :8080

Test:

    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"secret123"}'

    Response must contain both access_token and refresh_token.

============================================================
Goal
============================================================

Practice:

    - access vs refresh token
    - different expiration times
    - token type claim

============================================================
*/

package main

func main() {}
