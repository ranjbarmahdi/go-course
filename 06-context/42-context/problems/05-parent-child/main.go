/*
============================================================
Problem 5 — Parent → Child Cancellation
============================================================

Create a Context hierarchy:

	Background
	    |
	    v
	Parent Context
	    |
	    v
	Child Context
	    |
	    v
	Worker

Use:

	context.WithCancel()

Requirements:

	1. Create a parent Context.
	2. Create a child Context from the parent.
	3. Start a worker using the child Context.
	4. The worker should repeatedly print:

		"Worker working..."

	5. After 2 seconds, cancel the PARENT Context.
	6. The worker must detect cancellation through
	   the CHILD Context.
	7. Print:

		"Worker stopped"
		"Error:", childCtx.Err()

Important:

	Do NOT cancel the child directly.

	Cancel only the parent.

Expected behavior:

	cancel(parentCtx)
	    |
	    v
	childCtx cancelled
	    |
	    v
	worker stops

============================================================
Goal
============================================================

Understand:

	- Context trees
	- Parent → child propagation
	- WithCancel()
	- ctx.Done()
	- ctx.Err()

============================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("Implement later")
}
