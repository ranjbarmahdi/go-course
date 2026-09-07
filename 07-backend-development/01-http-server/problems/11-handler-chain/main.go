/*
============================================================
43 — HTTP SERVER
Problem 11 — Middleware Chain
============================================================

Create two routes:

    GET /users
    GET /products

Create three middleware:

    loggingMiddleware
    authMiddleware
    recoveryMiddleware

============================================================
GLOBAL MIDDLEWARE
============================================================


The following middleware must apply to ALL routes:

    loggingMiddleware
    recoveryMiddleware

So:

    Client
       ↓
    Logging
       ↓
    Recovery
       ↓
    Route Handler

Both /users and /products must pass through them.

Logging should print:

    Request: GET /users
or:

    Request: GET /products

============================================================
ROUTE-SPECIFIC MIDDLEWARE
============================================================

The /users route must additionally use:

    authMiddleware

The /products route must NOT use authMiddleware.

So /users flow:

    Client
       ↓
    loggingMiddleware      GLOBAL
       ↓
    recoveryMiddleware     GLOBAL
       ↓
    authMiddleware         ROUTE-SPECIFIC
       ↓
    usersHandler

And /products flow:

    Client
       ↓
    loggingMiddleware      GLOBAL
       ↓
    recoveryMiddleware     GLOBAL
       ↓
    productsHandler

============================================================
AUTHENTICATION
============================================================

authMiddleware should check:

    Authorization

If Authorization is missing:

    HTTP 401 Unauthorized

and:

    DO NOT call usersHandler

If Authorization exists:

    continue to usersHandler

============================================================
HANDLERS
============================================================

/users should return:

    Users endpoint

/products should return:

    Products endpoint

============================================================
REQUIREMENTS
============================================================

1. Create an explicit:

       http.NewServeMux()

2. Create:

       loggingMiddleware

3. Create:

       recoveryMiddleware

4. Create:

       authMiddleware

5. Apply loggingMiddleware globally.

6. Apply recoveryMiddleware globally.

7. Apply authMiddleware ONLY to /users.

8. Do NOT apply authMiddleware to /products.

9. Use middleware composition.

10. Understand the difference between:

       Global middleware
       Route-specific middleware

============================================================
EXPECTED STRUCTURE
============================================================

Global:

    recoveryMiddleware(
        loggingMiddleware(
            mux,
        ),
    )

Route-specific:

    authMiddleware(
        usersHandler,
    )

Conceptually:

    Server
      │
      ▼
    Global Middleware
      │
      ├────────────── /users
      │                 │
      │                 ▼
      │             Auth Middleware
      │                 │
      │                 ▼
      │             usersHandler
      │
      └────────────── /products
                        │
                        ▼
                   productsHandler

Practice:

    middleware composition
    http.Handler
    http.HandlerFunc
    ServeMux
    global middleware
    route-specific middleware
*/

package main

import (
	"fmt"
	"net/http"
	"time"
)

func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		t := time.Now()
		fmt.Println("===============================")
		fmt.Println("Request:", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)

		fmt.Println("Exit From Logger")
		fmt.Printf("========== <<Takes %f>> ===========\n", time.Since(t).Seconds())
	})
}

func recoveryMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Enter To Recovery")

		next.ServeHTTP(w, r)

		fmt.Println("Exit From Recovery")
	})
}

func authMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Enter To Auth")

		auth := r.Header.Get("Authorization")

		fmt.Println("Checking auth:", auth)

		if auth == "" {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)

		fmt.Println("Exit From Auth")
	})
}

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	time.Sleep(time.Second)
	fmt.Fprint(w, "Get Users")
}

func getProductsHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Get Products")
}

func main() {
	mux := http.NewServeMux()

	userHandler := authMiddleware(
		http.HandlerFunc(getUsersHandler),
	)

	mux.Handle("GET /users", userHandler)
	mux.HandleFunc("GET /products", getProductsHandler)

	handler := loggingMiddleware(
		recoveryMiddleware(mux),
	)

	server := http.Server{
		Addr:    ":8080",
		Handler: handler,
	}

	server.ListenAndServe()
}
