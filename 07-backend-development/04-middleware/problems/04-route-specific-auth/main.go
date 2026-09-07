/*
============================================================
46 — MIDDLEWARE
Problem 4 — Route-Specific Auth
============================================================

Create two routes:

    GET /users
    GET /products

Create:

    func authMiddleware(next http.Handler) http.Handler

Requirements:

    1. Apply authMiddleware ONLY to GET /users

    2. GET /products must NOT use authMiddleware

    3. authMiddleware checks the Authorization header

    4. If Authorization is missing:

           HTTP 401 Unauthorized
           DO NOT call the users handler

    5. If Authorization exists (any non-empty value):

           continue to users handler

    6. Handlers:

           GET /users     → Users endpoint
           GET /products  → Products endpoint

    7. Run the server on :8080

Test:

    curl http://localhost:8080/users              → 401
    curl http://localhost:8080/products           → 200
    curl -H "Authorization: token" http://localhost:8080/users → 200

============================================================
Goal
============================================================

Practice:

    - route-specific middleware
    - wrapping individual handlers
    - early return on auth failure

============================================================
*/

package main

func main() {}
