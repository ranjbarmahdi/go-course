/*
============================================================
Problem 6 — Request-Scoped Value
============================================================

Create a custom Context key:

	type contextKey string

Create:

	const requestIDKey contextKey = "requestID"

Requirements:

	1. Create a Context using context.Background().
	2. Store this request ID:

		"request-123"

	   using context.WithValue().

	3. Create a function:

		func service(ctx context.Context)

	4. Inside service(), retrieve the request ID.
	5. Use a type assertion to safely retrieve it.
	6. Print:

		Request ID: request-123

Then create another function:

		func repository(ctx context.Context)

The flow should be:

	main
	  |
	  v
	service
	  |
	  v
	repository

Both service and repository should be able
to access the request ID from the Context.

Do NOT pass requestID as a normal function parameter.

============================================================
Goal
============================================================

Practice:

	- context.WithValue()
	- ctx.Value()
	- Custom Context keys
	- Type assertions
	- Request-scoped values
	- Passing Context through layers

============================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("Implement later")
}
