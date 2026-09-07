/*
============================================================
46 — MIDDLEWARE
Problem 10 — CORS Middleware
============================================================

Create:

    GET /users

Create:

    func corsMiddleware(next http.Handler) http.Handler

Requirements:

    1. Set these response headers on every request:

           Access-Control-Allow-Origin: *
           Access-Control-Allow-Methods: GET, POST, PUT, DELETE, OPTIONS
           Access-Control-Allow-Headers: Content-Type, Authorization

    2. If the request method is OPTIONS:

           return 204 No Content
           DO NOT call the next handler

    3. For other methods, continue to the next handler

    4. GET /users returns:

           Users endpoint

    5. Apply corsMiddleware globally

    6. Run the server on :8080

Test:

    curl -X OPTIONS http://localhost:8080/users -i
        → 204 with CORS headers

    curl http://localhost:8080/users -i
        → 200 with CORS headers

============================================================
Goal
============================================================

Practice:

    - CORS headers
    - OPTIONS preflight handling
    - global CORS middleware

============================================================
*/

package main

func main() {}
