/*
============================================================
46 — MIDDLEWARE
Problem 8 — Global + Route-Specific Middleware
============================================================

Create two routes:

    GET /users
    GET /products

Create three middleware:

    loggingMiddleware
    recoveryMiddleware
    authMiddleware

============================================================
GLOBAL (all routes)
============================================================

    loggingMiddleware
    recoveryMiddleware

============================================================
ROUTE-SPECIFIC (/users only)
============================================================

    authMiddleware

============================================================
REQUEST FLOWS
============================================================

GET /users:

    logging → recovery → auth → usersHandler

GET /products:

    logging → recovery → productsHandler

============================================================
REQUIREMENTS
============================================================

    1. loggingMiddleware prints:

           Request: GET /users
       or:
           Request: GET /products

    2. authMiddleware checks Authorization header

    3. Missing Authorization on /users → 401

    4. Handlers:

           GET /users     → Users endpoint
           GET /products  → Products endpoint

    5. Run the server on :8080

============================================================
Goal
============================================================

Practice:

    - global middleware
    - route-specific middleware
    - combining both levels

============================================================
*/

package main

func main() {}
