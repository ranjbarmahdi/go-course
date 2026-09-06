/*
============================================================
43 — HTTP SERVER
Problem 9 — Middleware
============================================================

Create:

    GET /users

Create a logging middleware.

The middleware must print:

    ----- BEFORE -----
    GET /users

Then call the next handler.

After the handler finishes, print:

    ----- AFTER -----

Use:

    func loggingMiddleware(next http.Handler) http.Handler

The request flow must be:

    Client
       ↓
    Middleware
       ↓
    usersHandler
       ↓
    Response

Register the middleware around the router.

Goal:

Understand:

    next.ServeHTTP(w, r)
*/

package main

import "fmt"

func main() {
	fmt.Println("Not Implemented")
}
