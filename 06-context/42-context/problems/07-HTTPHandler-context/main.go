/*
============================================================
Problem 7 — HTTP Handler → Service → Repository
============================================================

Create a small HTTP server with this architecture:

	HTTP Handler
	     |
	     v
	Service
	     |
	     v
	Repository

The HTTP request Context must flow through
all three layers.

Requirements:

	1. Create an HTTP handler:

		func handler(
			w http.ResponseWriter,
			r *http.Request,
		)

	2. Get the request Context using:

		ctx := r.Context()

	3. Pass ctx to service().
	4. Pass ctx from service() to repository().
	5. repository() should simulate a 5-second operation.
	6. The handler should create a timeout Context
	   of 2 seconds from r.Context():

		ctx, cancel := context.WithTimeout(
			r.Context(),
			2*time.Second,
		)

	7. repository() must respect ctx.Done().
	8. If the timeout happens, return ctx.Err().
	9. The handler should return an appropriate HTTP error.

Architecture:

	Request
	   |
	   v
	Handler
	   |
	   | Context
	   v
	Service
	   |
	   | Context
	   v
	Repository
	   |
	   v
	5-second work

But:

	Timeout = 2 seconds

Therefore the repository should stop
before the 5-second operation finishes.

============================================================
Goal
============================================================

Practice real backend usage of Context:

	- r.Context()
	- WithTimeout()
	- Handler
	- Service
	- Repository
	- Context propagation
	- Cancellation

============================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("Implement later")
}
