/*
============================================================
43 — HTTP SERVER
Problem 5 — Request Headers
============================================================

Create:

    GET /headers

Read these request headers:

    Authorization
    Content-Type
    User-Agent

Print them to the terminal.

Return a response showing their values.

Example request:

    Authorization: Bearer abc123
    Content-Type: application/json

Practice:

    r.Header
    r.Header.Get()
*/

package main

import (
	"fmt"
	"net/http"
)

func headersHandler(w http.ResponseWriter, r *http.Request) {
	headers := r.Header
	fmt.Fprintf(w, "Authorization: %s\n", headers.Get("Authorization"))
	fmt.Fprintf(w, "Content-Type: %s\n", headers.Get("Content-Type"))
	fmt.Fprintf(w, "User-Agent: %s", headers.Get("User-Agent"))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/headers", headersHandler)

	http.ListenAndServe(":8080", mux)
}
