/*
============================================================
43 — HTTP SERVER
Problem 4 — Query Parameters
============================================================

Create:

    GET /users

The endpoint should accept:

    page
    limit
    search

Example:

    /users?page=2&limit=10&search=mahdi

Read the values using:

    r.URL.Query()

Then return:

    Page: 2
    Limit: 10
    Search: mahdi

Important:

Query parameter values are strings.

Bonus:

Convert page and limit to integers using strconv.Atoi().

Practice:

    r.URL.Query()
    query.Get()
    strconv.Atoi()
*/

package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func queryParamHandler(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	page, _ := strconv.Atoi(params.Get("page"))
	limit, _ := strconv.Atoi(params.Get("limit"))
	search := params.Get("search")

	fmt.Fprintf(w, "Page: %d\n", page)
	fmt.Fprintf(w, "Limit: %d\n", limit)
	fmt.Fprintf(w, "Search: %s", search)
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/users", queryParamHandler)

	http.ListenAndServe(":8080", mux)
}
