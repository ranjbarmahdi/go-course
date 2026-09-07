/*
============================================================
47 — AUTHENTICATION
Problem 12 — Protected vs Public Routes
============================================================

Create:

    POST /register     (public)
    POST /login        (public)
    POST /refresh      (public)
    GET  /health       (public)
    GET  /profile      (protected)
    GET  /settings     (protected)

Requirements:

    1. /health returns:

           {"status": "ok"}

       No authentication required.

    2. /profile and /settings require valid access token

    3. /profile returns:

           {"user_id": "...", "email": "..."}

    4. /settings returns:

           {"message": "Settings for <email>"}

    5. Protected routes without token → 401

    6. Use jwtAuthMiddleware on protected routes only

    7. Run server on :8080

Test:

    curl http://localhost:8080/health
        → 200 {"status":"ok"}

    curl http://localhost:8080/profile
        → 401

    curl http://localhost:8080/profile \
      -H "Authorization: Bearer <access_token>"
        → 200

============================================================
Goal
============================================================

Practice:

    - public vs protected routes
    - applying auth middleware selectively

============================================================
*/

package main

func main() {}
