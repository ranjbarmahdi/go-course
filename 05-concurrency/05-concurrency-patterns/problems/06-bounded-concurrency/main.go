/*
============================================================
Problem 6 — Bounded Concurrency ⭐⭐⭐
============================================================

Simulate processing database requests with a maximum
of 3 concurrent operations.

Requirements:

1. Create a list of 10 jobs:

    1 through 10

2. Create a semaphore:

    sem := make(chan struct{}, 3)

The capacity of 3 means:

    maximum 3 operations at once

3. Start one goroutine per job.

4. Before processing a job, acquire the semaphore:

    sem <- struct{}{}

5. Print:

    Starting job X

6. Process the job.

For demonstration, you may use:

    time.Sleep(500 * time.Millisecond)

7. Print:

    Finished job X

8. Release the semaphore:

    <-sem

9. Use WaitGroup so main waits for all jobs.

IMPORTANT:

There may be 10 goroutines,
but only 3 jobs may be inside the processing section
at the same time.

Concept:

    10 goroutines
         ↓
    ┌─────────────┐
    │  semaphore  │
    │ capacity 3  │
    └─────────────┘
         ↓
    max 3 active

*/

package main

import "fmt"

func main() {
	fmt.Println("bounded concurrency")
}
