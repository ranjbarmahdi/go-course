/*
============================================================
46 — MIDDLEWARE
Problem 1 — Middleware Pattern
============================================================

Create:

    GET /hello

Create a middleware called:

    exampleMiddleware

Signature:

    func exampleMiddleware(next http.Handler) http.Handler

Requirements:

    1. Before calling the next handler, print:

           BEFORE /hello

    2. Call next.ServeHTTP(w, r)

    3. After the handler finishes, print:

           AFTER /hello

    4. Wrap the router with exampleMiddleware

    5. The hello handler must return:

           Hello

    6. Run the server on :8080

Expected console output for GET /hello:

    BEFORE /hello
    AFTER /hello

============================================================
Goal
============================================================

Practice:

    - middleware signature
    - next.ServeHTTP(w, r)
    - before/after handler logic
    - wrapping a router

============================================================
*/

package main

func main() {}
