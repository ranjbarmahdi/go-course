/*
============================================================
43 — HTTP SERVER
Problem 8 — Custom Handler
============================================================

Create a custom type:

    type HelloHandler struct{}

Make it implement:

    http.Handler

Implement:

    ServeHTTP()

Create:

    GET /hello

Response:

    Hello from custom Handler

Register it using:

    mux.Handle()

Do NOT use:

    mux.HandleFunc()

Goal:

Understand the relationship between:

    Handler
    ServeHTTP
    Handle()
    HandlerFunc
*/

package main

import "fmt"

func main() {
	fmt.Println("Not Implemented")
}
