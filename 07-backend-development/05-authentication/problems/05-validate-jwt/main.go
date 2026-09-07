/*
============================================================
47 — AUTHENTICATION
Problem 5 — Validate JWT
============================================================

Extend TokenService with:

    func (s *TokenService) ParseToken(tokenString string) (*TokenClaims, error)

Requirements:

    1. Parse token with jwt.ParseWithClaims

    2. Verify signing method is HS256

    3. Verify signature with the secret key

    4. Return *TokenClaims on success

    5. Return error if:

           token is invalid
           token is expired
           signature is wrong

    6. In main(), demonstrate:

           generate valid token   → parse succeeds, print user_id and email
           parse tampered token   → parse fails
           parse expired token    → parse fails

    7. Use a custom TokenClaims struct embedding jwt.RegisteredClaims

Allowed packages:

    fmt
    time
    errors
    github.com/golang-jwt/jwt/v5

============================================================
Goal
============================================================

Practice:

    - JWT validation
    - jwt.ParseWithClaims
    - handling expired tokens

============================================================
*/

package main

func main() {}
