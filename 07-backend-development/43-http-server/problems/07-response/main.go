/*
============================================================
43 — HTTP SERVER
Problem 7 — Response
============================================================

Create:

    POST /users

Return:

    Status: 201 Created

    Body:
    User created

Set:

    Content-Type: text/plain

Requirements:

1. Set the response header.

2. Set HTTP status to:

       http.StatusCreated

3. Write the response body.

Correct order:

    Header
      ↓
    Status
      ↓
    Body

Practice:

    w.Header()
    w.WriteHeader()
    fmt.Fprintln()
*/

package main

import (
	"fmt"
	"net/http"
)

func responseHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusCreated)

	fmt.Fprintln(w, "User created")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /users", responseHandler)

	http.ListenAndServe(":8080", mux)
}
