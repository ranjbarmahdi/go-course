/*
============================================================
47 — AUTHENTICATION
Problem 13 — TokenService
============================================================

Create a complete TokenService without HTTP handlers.

Type:

    type TokenService struct { ... }

Functions:

    NewTokenService(secret string) *TokenService
    GenerateAccessToken(userID, email string) (string, error)
    GenerateRefreshToken(userID, email string) (string, error)
    ParseToken(tokenString string) (*TokenClaims, error)

Requirements:

    1. Use github.com/golang-jwt/jwt/v5

    2. TokenClaims struct:

           UserID string
           Email  string
           Type   string   ("access" or "refresh")
           jwt.RegisteredClaims

    3. Access token TTL:  15 minutes
    4. Refresh token TTL: 7 days
    5. Signing method:    HS256

    6. ParseToken must reject wrong signing method

    7. In main(), demonstrate:

           generate access token  → parse → print UserID, Email, Type
           generate refresh token → parse → print UserID, Email, Type
           parse with wrong secret → error

Allowed packages:

    fmt
    time
    errors
    github.com/golang-jwt/jwt/v5

============================================================
Goal
============================================================

Practice:

    - isolating token logic
    - reusable TokenService
    - access and refresh token generation

============================================================
*/

package main

func main() {}
