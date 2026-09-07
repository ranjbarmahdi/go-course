/*
============================================================
Problem 3 — Operation Timeout
============================================================

Create a function:

	func doWork(ctx context.Context) error

The function simulates an operation that takes
5 seconds.

However, main() should create a Context with
a timeout of 2 seconds.

Requirements:

	1. Create:

		ctx, cancel := context.WithTimeout(
			context.Background(),
			2*time.Second,
		)

	2. Use defer cancel().
	3. Start doWork(ctx).
	4. doWork must listen for Context cancellation.
	5. If the work finishes first, return nil.
	6. If the Context expires first, return:

		ctx.Err()

Expected result:

	context deadline exceeded

The important part:

	The 5-second operation must NOT continue
	after the 2-second timeout.

============================================================
Goal
============================================================

Practice:

	- context.WithTimeout()
	- ctx.Done()
	- ctx.Err()
	- Timeout handling

============================================================
*/

package main

import "fmt"

func main() {
	fmt.Println("Implement later")
}
