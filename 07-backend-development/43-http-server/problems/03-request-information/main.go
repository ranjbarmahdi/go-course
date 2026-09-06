/*
============================================================
43 — HTTP SERVER
Problem 3 — Request Information
============================================================

Create:

    GET /request

Print the following information to the terminal:

    Method
    URL
    Path
    Host

Also return:

    "Request information received"

Use:

    r.Method
    r.URL
    r.URL.Path
    r.Host

Practice understanding:

    *http.Request
*/

package main

import (
	"fmt"
	"net/http"
)

func requestInformation(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Method:", r.Method)
	fmt.Println("Url:", r.URL)
	fmt.Println("Path:", r.URL.Path)
	fmt.Println("Host:", r.Host)
	fmt.Fprint(w, "Request information received")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/request", requestInformation)

	http.ListenAndServe(":8080", mux)
}
