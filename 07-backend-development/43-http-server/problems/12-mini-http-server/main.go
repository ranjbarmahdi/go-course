/*
============================================================
43 — HTTP SERVER
Problem 12 — Mini HTTP Server
============================================================

Build a small HTTP server combining everything learned
in Lecture 43.

Create these routes:

	GET  /
	GET  /users
	POST /users
	GET  /users/search
	GET  /headers
	POST /echo
	GET  /context

Requirements:

1. Use an explicit:

	http.NewServeMux()

2. Use an:

	http.Server

3. Implement method checking.

4. Read query parameters.

5. Read request headers.

6. Read request body.

7. Return appropriate status codes.

8. Use http.Error() for errors.

9. Create logging middleware.

10. Pass r.Context() into a service function.

11. Chain the middleware around the router.

Expected architecture:

	Client
	   ↓
	http.Server
	   ↓
	Logging Middleware
	   ↓
	ServeMux
	   ↓
	Route
	   ↓
	Handler
	   ↓
	Service
	   ↓
	Context

Example:

	GET /users?page=2&limit=10

	↓

	Middleware

	↓

	ServeMux

	↓

	usersHandler

	↓

	Query Parameters

	↓

	Response

Do NOT use:

	Gin
	Echo
	Fiber
	Chi

Use only the standard library.

Packages you may need:

	net/http
	context
	fmt
	io
	strconv
	time

============================================================
GOAL
============================================================

This is the final integration problem for:

	43 — HTTP Server

After completing this problem, you should understand:

	Handler
	HandlerFunc
	ServeMux
	http.Server
	Request
	ResponseWriter
	HTTP Methods
	URL
	Query Parameters
	Headers
	Body
	Status Codes
	http.Error
	Context
	Middleware
	Handler chains
*/
package main

import "fmt"

func main() {
	fmt.Println("Not Implemented")
}
