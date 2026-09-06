/*
============================================================
43 — HTTP SERVER
Problem 6 — Request Body
============================================================

Create:

	POST /echo

The client sends any body.

Example:

	Hello Go

Read the request body using:

	io.ReadAll(r.Body)

Then return:

	Received: Hello Go

If reading the body fails:

	HTTP 400 Bad Request

Practice:

	r.Body
	io.ReadAll()
	http.Error()
*/

package main

import (
	"fmt"
	"io"
	"net/http"
)

func bodyHandler(w http.ResponseWriter, r *http.Request) {
	res, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, "Reading Failed", http.StatusBadRequest)
		return
	}

	fmt.Println(string(res))
	fmt.Fprintln(w, "Received:", string(res))
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/echo", bodyHandler)

	http.ListenAndServe(":8080", mux)
}
