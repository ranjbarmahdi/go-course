/*
============================================================
46 — MIDDLEWARE
Problem 5 — Recovery Middleware
============================================================

Create:

    GET /panic
    GET /ok

Create:

    func recoveryMiddleware(next http.Handler) http.Handler

Requirements:

    1. recoveryMiddleware must catch panics using defer + recover()

    2. On panic:

           print: PANIC: <error>
           return HTTP 500 Internal Server Error

    3. GET /panic handler must panic with:

           panic("something went wrong")

    4. GET /ok handler returns:

           OK

    5. Apply recoveryMiddleware globally around the router

    6. The server must NOT crash when /panic is called

    7. Run the server on :8080

Test:

    curl http://localhost:8080/panic  → 500, server still running
    curl http://localhost:8080/ok     → 200 OK

============================================================
Goal
============================================================

Practice:

    - panic recovery
    - defer and recover()
    - keeping the server alive

============================================================
*/

package main

func main() {}
