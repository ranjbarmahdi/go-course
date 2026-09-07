/*
============================================================
46 — MIDDLEWARE
Problem 13 — Request ID Middleware
============================================================

Create:

    GET /users

Create:

    func requestIDMiddleware(next http.Handler) http.Handler

Requirements:

    1. Generate a request ID for every request

       Use a simple counter starting at 1:

           req-1, req-2, req-3, ...

       (package-level variable is fine)

    2. Store the request ID in context using a typed key:

           type contextKey string
           const requestIDKey contextKey = "requestID"

    3. Also set response header:

           X-Request-ID: req-1

    4. usersHandler reads request ID from context and returns:

           Request ID: req-1

    5. Apply requestIDMiddleware globally

    6. Run the server on :8080

Test:

    curl http://localhost:8080/users -i
        → X-Request-ID header + body with same ID

    curl http://localhost:8080/users
        → next ID increments

============================================================
Goal
============================================================

Practice:

    - request tracing
    - response headers in middleware
    - context values

============================================================
*/

package main

func main() {}
