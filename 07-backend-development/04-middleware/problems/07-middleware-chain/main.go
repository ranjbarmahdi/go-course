/*
============================================================
46 — MIDDLEWARE
Problem 7 — Middleware Chain
============================================================

Create:

    GET /users

Create two middleware:

    loggingMiddleware
    recoveryMiddleware

Requirements:

    1. Apply BOTH globally around the router

    2. Chain order (request direction):

           logging → recovery → mux → handler

    3. loggingMiddleware prints:

           → GET /users

       before the handler

    4. recoveryMiddleware catches panics

    5. usersHandler returns:

           Users endpoint

    6. Use nested wrapping:

           loggingMiddleware(
               recoveryMiddleware(mux),
           )

    7. Run the server on :8080

============================================================
Goal
============================================================

Practice:

    - middleware chaining
    - nesting wrappers
    - execution order

============================================================
*/

package main

func main() {}
