/*
============================================================
46 — MIDDLEWARE
Problem 14 — Complete Middleware Server
============================================================

Create a complete API with middleware.

Routes:

    GET /users
    GET /products
    GET /profile
    GET /panic

Middleware:

    corsMiddleware       (global)
    loggingMiddleware    (global)
    recoveryMiddleware   (global)
    authMiddleware       (route-specific on /users and /profile)

============================================================
REQUIREMENTS
============================================================

    1. Global chain order:

           cors → logging → recovery → mux

    2. loggingMiddleware logs method, path, and duration

    3. recoveryMiddleware catches panics → 500

    4. authMiddleware checks Authorization header → 401 if missing

    5. Protected routes: /users, /profile
       Public routes:   /products, /panic

    6. Handlers:

           GET /users     → Users endpoint
           GET /products  → Products endpoint
           GET /profile   → Profile endpoint
           GET /panic     → panic("test panic")

    7. corsMiddleware handles OPTIONS with 204

    8. Use http.Server on :8080

    9. Do NOT use third-party packages

============================================================
TEST
============================================================

    curl http://localhost:8080/products
        → 200, CORS headers, logged

    curl http://localhost:8080/users
        → 401

    curl -H "Authorization: token" http://localhost:8080/users
        → 200

    curl http://localhost:8080/panic
        → 500, server still running

============================================================
Goal
============================================================

Practice:

    - complete middleware architecture
    - global + route-specific middleware
    - logging, recovery, auth, CORS together
    - production-style HTTP server setup

============================================================
*/

package main

func main() {}
