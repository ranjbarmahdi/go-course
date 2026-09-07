/*
============================================================
47 — AUTHENTICATION
Problem 4 — Generate JWT
============================================================

Create a TokenService with:

    func NewTokenService(secret string) *TokenService
    func (s *TokenService) GenerateAccessToken(userID, email string) (string, error)

Requirements:

    1. Use github.com/golang-jwt/jwt/v5

    2. Signing method: HS256

    3. Claims must include:

           user_id  → string
           email    → string
           type     → "access"
           exp      → 15 minutes from now
           iat      → issued at

    4. Use jwt.RegisteredClaims for exp and iat

    5. In main(), generate a token for:

           user_id: "user-1"
           email:   "mahdi@example.com"

    6. Print the token string

    7. Token must have 3 dot-separated parts (header.payload.signature)

Allowed packages:

    fmt
    time
    github.com/golang-jwt/jwt/v5

============================================================
Goal
============================================================

Practice:

    - JWT structure
    - custom claims
    - token signing
    - expiration

============================================================
*/

package main

func main() {}
