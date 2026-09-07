/*
============================================================
Problem 1 — Manual Cancellation with context.WithCancel
============================================================

Create a cancellable Context using:

	context.WithCancel()

Requirements:

	1. Create a Context using context.Background().
	2. Start a goroutine called worker.
	3. The worker should repeatedly print:

		"Working..."

	4. The worker must listen to:

		ctx.Done()

	5. When the Context is cancelled, the worker must print:

		"Worker stopped"

	   and return.

	6. From main(), wait for a short period and then
	   call cancel().

	7. main() must wait long enough to clearly show
	   that the worker stopped because of cancellation.

Important:

	Do NOT use a WaitGroup.

	The worker must stop because of Context cancellation.

Expected behavior:

	Working...
	Working...
	Working...
	Worker stopped

============================================================
Goal
============================================================

Practice:

	- context.WithCancel()
	- ctx.Done()
	- Goroutines
	- Cooperative cancellation

============================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("Implement later")
}
