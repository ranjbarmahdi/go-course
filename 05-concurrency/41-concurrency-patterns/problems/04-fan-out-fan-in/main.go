/*
============================================================
Problem 4 — Fan-Out + Fan-In ⭐⭐⭐
============================================================

Combine Fan-Out and Fan-In.

Build this system:

                 ┌──> Worker 1 ──┐
    jobs ────────┼──> Worker 2 ──┼──> results
                 └──> Worker 3 ──┘

Requirements:

1. Create:

    jobs := make(chan int)
    results := make(chan int)

2. Start 3 workers.

3. Every worker receives jobs from the SAME jobs channel.

4. Each worker:
- Receives a job.
- Squares the job.
- Sends the result to results.

Example:

    5 → 25
    7 → 49

5. Send 10 jobs:

    1 through 10

6. Close jobs after sending.

7. Use WaitGroup to detect when all workers finish.

8. Close results after all workers finish.

9. Main ranges over results and prints them.

Expected values:

    1, 4, 9, 16, 25,
    36, 49, 64, 81, 100

Order is NOT guaranteed.

This problem combines:

    Fan-Out
        +
    Fan-In

*/

package main

import "fmt"

func main() {
	fmt.Println("fan in out")
}
