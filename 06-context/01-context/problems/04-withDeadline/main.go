/*
============================================================
Problem 4 — Deadline-Based Cancellation
============================================================

Create a Context using:

	context.WithDeadline()

The deadline should be approximately 2 seconds
from the current time.

Create a worker that performs work continuously.

Requirements:

	1. Create a deadline:

		deadline := time.Now().Add(2 * time.Second)

	2. Create a Context using WithDeadline().
	3. Start a worker goroutine.
	4. The worker should print:

		"Working..."

	   every 500 milliseconds.

	5. The worker must stop when:

		<-ctx.Done()

	6. When stopped, print:

		"Stopped:", ctx.Err()

Expected result:

	Working...
	Working...
	Working...
	Working...
	Stopped: context deadline exceeded

Do not manually call cancel() to stop the worker.

The deadline itself should cause cancellation.

============================================================
Goal
============================================================

Practice:

	- context.WithDeadline()
	- time.Time
	- ctx.Done()
	- ctx.Err()

============================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("Implement later")
}
