/*
============================================================
43 — HTTP SERVER
Problem 10 — Request Context
============================================================

Create:

    GET /context

Get the request context:

    ctx := r.Context()

Create a service function:

    func service(ctx context.Context) error

The service should simulate work for:

    500 milliseconds

Use:

    select
    time.After()
    ctx.Done()

The handler should:

    r.Context()
        ↓
    service(ctx)

If the context is cancelled:

    return an appropriate error

Goal:

Connect HTTP requests with the Context lecture.

Flow:

    HTTP Request
         ↓
    Handler
         ↓
    Service
         ↓
    Context
*/

package main

import "fmt"

func main() {
	fmt.Println("Not Implemented")
}
