/*
============================================================
47 — AUTHENTICATION
Problem 7 — Profile With Context
============================================================

Create:

    POST /login
    GET  /profile

Requirements:

    1. Login returns access_token (JWT with user_id and email)

    2. jwtAuthMiddleware must attach to context:

           user_id
           email

       Use typed context keys:

           type contextKey string
           const userIDKey contextKey = "userID"
           const userEmailKey contextKey = "userEmail"

    3. GET /profile reads from context and returns JSON:

           {
               "user_id": "user-1",
               "email": "mahdi@example.com"
           }

    4. Without token → 401

    5. With valid token → 200 and correct user data from token

    6. Run server on :8080

Test:

    curl -X POST http://localhost:8080/login \
      -H "Content-Type: application/json" \
      -d '{"email":"mahdi@example.com","password":"secret123"}'

    curl http://localhost:8080/profile \
      -H "Authorization: Bearer <access_token>"

============================================================
Goal
============================================================

Practice:

    - context.WithValue in auth middleware
    - reading user from context in handler
    - connecting login → protected route

============================================================
*/

package main

func main() {}
