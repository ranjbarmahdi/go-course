/*
============================================================
46 — MIDDLEWARE
Problem 11 — Chain Helper
============================================================

Create:

    GET /users
    GET /products

Create a chain helper:

    func chain(
        middlewares ...func(http.Handler) http.Handler,
    ) func(http.Handler) http.Handler

Requirements:

    1. chain must apply middleware in the correct order:

           cors → logging → recovery → mux

    2. Create these middleware:

           corsMiddleware
           loggingMiddleware
           recoveryMiddleware

    3. loggingMiddleware prints:

           → METHOD PATH

    4. corsMiddleware sets:

           Access-Control-Allow-Origin: *

    5. Handlers:

           GET /users     → Users endpoint
           GET /products  → Products endpoint

    6. Use chain like this:

           handler := chain(
               corsMiddleware,
               loggingMiddleware,
               recoveryMiddleware,
           )(mux)

    7. Run the server on :8080

============================================================
Goal
============================================================

Practice:

    - reusable chain helper
    - applying multiple middleware cleanly
    - middleware order

============================================================
*/

package main

func main() {}
