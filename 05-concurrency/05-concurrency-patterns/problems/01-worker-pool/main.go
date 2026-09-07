/*
============================================================
Problem 1 — Worker Pool ⭐⭐⭐
============================================================

Build a simple worker pool using goroutines and channels.

Requirements:

1. Create a jobs channel:

    jobs := make(chan int)

2. Create a results channel:

    results := make(chan int)

3. Create 3 worker goroutines.

Each worker must:
- Continuously receive jobs from the jobs channel.
- Process each job by multiplying it by 2.
- Send the processed result to the results channel.
- Stop when the jobs channel is closed.

Example:

    Job:    10
    Result: 20

4. Send 6 jobs from main:

    1, 2, 3, 4, 5, 6

5. Close the jobs channel after all jobs are sent.

6. Use sync.WaitGroup to know when all 3 workers have finished.

7. After all workers finish, close the results channel.

8. Range over the results channel in main and print every result.

Expected results:

    2
    4
    6
    8
    10
    12

IMPORTANT:

The order of results is NOT guaranteed.

For example, this is valid:

    2
    6
    4
    10
    8
    12

because multiple workers are processing jobs concurrently.

Worker pool structure:

    jobs
      │
      ├──> Worker 1 ──┐
      │               │
      ├──> Worker 2 ──┼──> results
      │               │
      └──> Worker 3 ──┘

Use:

- goroutines
- channels
- sync.WaitGroup
- close()
- range over channel

Do NOT use time.Sleep() for synchronization.
*/

package main

import (
	"fmt"
	"sync"
)

func worker(id int, jobs <-chan int, results chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	for job := range jobs {
		fmt.Println("Worker", id, "processing job", job)

		results <- job * 2
	}
}

func main() {
	jobs := make(chan int)
	results := make(chan int)

	var wg sync.WaitGroup

	wg.Add(3)

	for i := 1; i <= 3; i++ {
		go worker(i, jobs, results, &wg)
	}

	go func() {
		for i := 1; i <= 6; i++ {
			jobs <- i
		}

		close(jobs)
	}()

	go func() {
		wg.Wait()
		close(results)
	}()

	for result := range results {
		fmt.Println("Result:", result)
	}
}
