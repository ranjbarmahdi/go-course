/*
============================================================
46 — MIDDLEWARE
Problem 2 — Logging Middleware
============================================================

Create:

    GET /users

Create:

    func loggingMiddleware(next http.Handler) http.Handler

Requirements:

    1. Before the handler, print:

           → GET /users

       (use r.Method and r.URL.Path — do not hardcode the path)

    2. After the handler, print duration:

           ← GET /users (2ms)

       Format: ← METHOD PATH (duration)

    3. Use time.Now() and time.Since()

    4. The users handler returns:

           Users endpoint

    5. Run the server on :8080

============================================================
Goal
============================================================

Practice:

    - request logging
    - measuring handler duration
    - logging middleware

============================================================
*/

package main

func main() {}
