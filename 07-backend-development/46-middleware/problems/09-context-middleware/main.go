/*
============================================================
46 — MIDDLEWARE
Problem 9 — Context Middleware
============================================================

Create:

    GET /profile

Create:

    func authContextMiddleware(next http.Handler) http.Handler

Requirements:

    1. Check Authorization header

    2. If missing → 401 Unauthorized

    3. If present:

           attach userID "user-42" to request context
           using context.WithValue

    4. Use a typed context key (NOT a plain string):

           type contextKey string
           const userIDKey contextKey = "userID"

    5. profileHandler reads userID from context and returns:

           Profile for user: user-42

    6. Run the server on :8080

Test:

    curl -H "Authorization: Bearer token" http://localhost:8080/profile
        → Profile for user: user-42

============================================================
Goal
============================================================

Practice:

    - context.WithValue
    - r.WithContext
    - passing data through middleware
    - typed context keys

============================================================
*/

package main

func main() {}
