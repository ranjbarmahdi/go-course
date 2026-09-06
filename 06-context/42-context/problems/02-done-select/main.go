/*
============================================================
Problem 2 — Context Cancellation with Select
============================================================

Create:

	ctx, cancel := context.WithCancel(...)

Create a worker goroutine.

The worker should use select:

	select {
	case <-ctx.Done():
		// stop
	default:
		// work
	}

Requirements:

	1. The worker should run continuously.
	2. Every 300 milliseconds, print:

		"Working..."

	3. The worker must use select to check ctx.Done().
	4. main() should cancel the Context after 2 seconds.
	5. After cancellation, the worker must print:

		"Cancelled:", ctx.Err()

	6. The worker must return after cancellation.

Do NOT use:

	- WaitGroup
	- Global variables

============================================================
Goal
============================================================

Practice:

	- ctx.Done()
	- select
	- context.WithCancel()
	- ctx.Err()
	- Cooperative cancellation

============================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("Implement later")
}
