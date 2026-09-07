/*
============================================================
46 — MIDDLEWARE
Problem 12 — Block Request Middleware
============================================================

Create:

    GET /admin

Create:

    func apiKeyMiddleware(next http.Handler) http.Handler

Requirements:

    1. Read the X-API-Key header

    2. If X-API-Key is missing or not equal to "secret-key":

           HTTP 403 Forbidden
           body: Forbidden
           DO NOT call adminHandler

    3. If X-API-Key is "secret-key":

           continue to adminHandler

    4. adminHandler returns:

           Admin panel

    5. Run the server on :8080

Test:

    curl http://localhost:8080/admin
        → 403 Forbidden

    curl -H "X-API-Key: wrong" http://localhost:8080/admin
        → 403 Forbidden

    curl -H "X-API-Key: secret-key" http://localhost:8080/admin
        → 200 Admin panel

============================================================
Goal
============================================================

Practice:

    - custom header validation
    - blocking requests in middleware
    - 403 Forbidden

============================================================
*/

package main

func main() {}
