/*
============================================================
Problem 8 — Combined Concurrency Patterns ⭐⭐⭐⭐⭐
============================================================

Build a small backend-style job processing system.

The system should use:

    Worker Pool
        +
    Fan-Out
        +
    Fan-In
        +
    Bounded Concurrency
        +
    Graceful Shutdown

Scenario:

You have 20 image-processing jobs.

Each job has an ID:

    1 through 20

Requirements:

1. Create:

    jobs := make(chan int)
    results := make(chan int)

2. Create 3 workers.

This is the Worker Pool / Fan-Out part.

3. Each worker receives jobs from the same jobs channel.

4. Each worker must process the image job.

For this exercise, simulate processing by calculating:

    result := job * 100

5. Add bounded concurrency.

Only 2 jobs may be actively processed at the same time.

Use a semaphore with capacity 2.

6. Every completed job sends its result to the ONE results
   channel.

This is the Fan-In part.

7. Use WaitGroup to track workers.

8. Close jobs when no more jobs will be submitted.

9. When all workers finish, close results.

10. Main ranges over results and counts how many jobs
    completed.

11. Add a shutdown channel.

If shutdown is triggered, workers must stop accepting
new jobs and exit gracefully.

12. Print:

    Completed jobs: X

The important architecture is:

                 jobs
                   │
          ┌────────┼────────┐
          ↓        ↓        ↓
       Worker 1 Worker 2 Worker 3
          │        │        │
          └────────┼────────┘
                   ↓
                results
                   ↓
                  main

Inside the workers:

             worker
                ↓
          semaphore
          capacity 2
                ↓
           processing
                ↓
             result

This is the final exercise for the lecture.

Focus on understanding how the patterns combine rather
than trying to make the code unnecessarily complicated.
*/

package main

import "fmt"

func main() {
	fmt.Println("Combined Concurrency Pattern")
}
