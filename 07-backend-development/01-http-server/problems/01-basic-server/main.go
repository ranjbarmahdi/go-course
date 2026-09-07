/*
============================================================
43 — HTTP SERVER
Problem 1 — Basic Server + Routes
============================================================

Create an HTTP server using net/http.

Requirements:

1. Create an explicit ServeMux.

2. Create these routes:

       GET /
       GET /users
       GET /products

3. Responses:

       /
       → "Welcome to API"

       /users
       → "Users endpoint"

       /products
       → "Products endpoint"

4. Create an http.Server:

       Addr: ":8080"
       Handler: mux

5. Start the server.

Do NOT use the default global router.

Practice:

    http.NewServeMux()
    mux.HandleFunc()
    http.Server
    http.ResponseWriter
    http.Request
*/

package main

import (
	"fmt"
	"net/http"
)

func welcomeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome to API")
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Users endpoint")
}

func productHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Products endpoint")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", welcomeHandler)
	mux.HandleFunc("/users", usersHandler)
	mux.HandleFunc("/products", productHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	server.ListenAndServe()
}
