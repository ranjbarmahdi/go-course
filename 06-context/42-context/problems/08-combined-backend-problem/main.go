/*
============================================================
Problem 8 — Complete Backend Context Problem
============================================================

Build a small backend-style program that combines:

	- HTTP
	- Context
	- Timeout
	- Cancellation
	- Goroutines
	- Channels
	- Select
	- Request-scoped values

Architecture:

	HTTP Request
	     |
	     v
	Handler
	     |
	     v
	Service
	     |
	     +--------> Worker 1
	     |
	     +--------> Worker 2
	     |
	     +--------> Worker 3

Requirements:

------------------------------------------------------------
1. HTTP Handler
------------------------------------------------------------

Create an HTTP endpoint:

	GET /process

The handler should:

	- Get r.Context()
	- Create a 3-second timeout Context
	- Create a request ID
	- Store the request ID in the Context
	- Call service(ctx)

------------------------------------------------------------
2. Service
------------------------------------------------------------

Create:

	func service(ctx context.Context) error

The service should:

	- Start 3 worker goroutines.
	- Give each worker some work.
	- Workers must receive the Context.
	- Workers must stop if the Context is cancelled.

------------------------------------------------------------
3. Workers
------------------------------------------------------------

Each worker should simulate work that takes
different amounts of time.

For example:

	Worker 1 → 1 second
	Worker 2 → 2 seconds
	Worker 3 → 5 seconds

Workers must use select:

	select {
	case <-time.After(duration):
		// work completed

	case <-ctx.Done():
		// cancelled
	}

------------------------------------------------------------
4. Result Channel
------------------------------------------------------------

Workers should send their results to a channel.

Example:

	"Worker 1 completed"
	"Worker 2 completed"

The service should collect the results.

------------------------------------------------------------
5. Timeout
------------------------------------------------------------

The request timeout is:

	3 seconds

Therefore:

	Worker 1 → completes
	Worker 2 → completes
	Worker 3 → gets cancelled

The service must handle this correctly.

------------------------------------------------------------
6. Request ID
------------------------------------------------------------

Store:

	"request-123"

inside the Context.

Workers should retrieve the request ID
from the Context and include it in their output.

Example:

	"request-123 - Worker 1 completed"

------------------------------------------------------------
7. Graceful Cancellation
------------------------------------------------------------

When the Context is cancelled:

	- Workers must stop.
	- No worker should continue unnecessary work.
	- The service should return an error when appropriate.
	- No goroutine should remain blocked forever.

------------------------------------------------------------
Expected Architecture
------------------------------------------------------------

	Client
	   |
	   v
	HTTP Handler
	   |
	   | Context
	   v
	Service
	   |
	   +------------------+
	   |        |         |
	   v        v         v
	Worker1  Worker2   Worker3
	   |        |         |
	   +--------+---------+
	            |
	            v
	       Result Channel

Context:

	HTTP Request
	      |
	      v
	   Timeout
	      |
	      v
	    Workers

------------------------------------------------------------
Important
------------------------------------------------------------

Do not use:

	- Global variables
	- time.Sleep() as synchronization
	- WaitGroup as the main communication mechanism

Use:

	- Context
	- Channels
	- Select
	- Goroutines

============================================================
Goal
============================================================

This problem combines almost everything
from the Context lesson.

You should understand:

	- WithCancel
	- WithTimeout
	- Done
	- Err
	- WithValue
	- Value
	- Parent → child Context
	- HTTP request Context
	- Goroutines
	- Channels
	- Select
	- Cooperative cancellation
	- Backend architecture

============================================================
End of Context Problems
============================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("Implement later")
}
