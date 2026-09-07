/*
============================================================
43 — HTTP SERVER
Problem 2 —
HTTP Methods
============================================================

Create:

    /users

The endpoint must support:

    GET
    POST
    PUT
    PATCH
    DELETE

Return:

    GET    → "Get users"
    POST   → "Create user"
    PUT    → "Update user"
    PATCH  → "Patch user"
    DELETE → "Delete user"

For any other method:

    HTTP 405 Method Not Allowed

Use:

    r.Method

and preferably:

    http.MethodGet
    http.MethodPost
    http.MethodPut
    http.MethodPatch
    http.MethodDelete

Practice:

    HTTP methods
    switch
    http.Error()
    status codes
*/

package main

import (
	"net/http"
)

func getUsersHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Write([]byte("Get users"))

	case http.MethodPost:
		w.Write([]byte("Create user"))

	case http.MethodPut:
		w.Write([]byte("Update user"))

	case http.MethodPatch:
		w.Write([]byte("Patch user"))

	case http.MethodDelete:
		w.Write([]byte("Delete user"))

	default:
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
	}
}

// func createUserHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Create user")
// }

// func updateUserHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Update user")
// }

// func patchUserHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Patch user")
// }

// func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Delete user")
// }

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", getUsersHandler)

	server := http.Server{
		Addr:    ":8080",
		Handler: mux,
	}
	server.ListenAndServe()
}
