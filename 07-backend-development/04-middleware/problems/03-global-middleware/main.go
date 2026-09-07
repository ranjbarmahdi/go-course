/*
============================================================
46 — MIDDLEWARE
Problem 3 — Global Middleware
============================================================

Create two routes:

    GET /users
    GET /products

Create:

    func loggingMiddleware(next http.Handler) http.Handler

Requirements:

    1. Apply loggingMiddleware GLOBALLY to the router

    2. Both routes must pass through loggingMiddleware

    3. Logging must print for every request:

           → METHOD PATH

    4. Handlers:

           GET /users     → Users endpoint
           GET /products  → Products endpoint

    5. Use http.NewServeMux()

    6. Run the server on :8080

Expected:

    GET /users     → logs → GET /users
    GET /products  → logs → GET /products

============================================================
Goal
============================================================

Practice:

    - global middleware
    - wrapping ServeMux
    - shared logic for all routes

============================================================
*/

package main

func main() {}
