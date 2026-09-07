/*
============================================================
46 — MIDDLEWARE
Problem 6 — Auth Middleware
============================================================

Create:

    GET /profile

Create:

    func authMiddleware(next http.Handler) http.Handler

Requirements:

    1. Protect GET /profile with authMiddleware

    2. Read the Authorization header

    3. If Authorization is empty:

           HTTP 401 Unauthorized
           body: Unauthorized
           DO NOT call profileHandler

    4. If Authorization is present:

           call profileHandler

    5. profileHandler returns:

           Profile endpoint

    6. Run the server on :8080

Test:

    curl http://localhost:8080/profile
        → 401 Unauthorized

    curl -H "Authorization: Bearer mytoken" http://localhost:8080/profile
        → 200 Profile endpoint

============================================================
Goal
============================================================

Practice:

    - authentication middleware
    - http.Error for 401
    - blocking unauthorized requests

============================================================
*/

package main

func main() {}
