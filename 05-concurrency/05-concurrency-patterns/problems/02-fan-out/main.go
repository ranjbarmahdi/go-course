/*
============================================================
Problem 2 — Fan-Out ⭐⭐⭐
============================================================

Build a simple Fan-Out pattern.

The idea:

    One source of jobs
             │
             ├──> Worker 1
             │
             ├──> Worker 2
             │
             └──> Worker 3

Multiple workers should consume jobs from the same channel.

Requirements:

1. Create a jobs channel:

    jobs := make(chan int)

2. Create 3 worker goroutines.

Each worker must:
- Receive jobs from the jobs channel.
- Print which worker received the job.
- Process the job by multiplying it by 10.
- Print the processed result.

3. Send these 9 jobs:

    1, 2, 3, 4, 5, 6, 7, 8, 9

4. Close the jobs channel after sending all jobs.

5. Use sync.WaitGroup so main waits until all workers finish.

6. Do NOT create a separate channel for each worker.

All workers must receive from the SAME jobs channel.

Example output:

    Worker 1 received job 1
    Worker 1 result: 10

    Worker 2 received job 2
    Worker 2 result: 20

    Worker 3 received job 3
    Worker 3 result: 30

The exact order and distribution are NOT guaranteed.

For example, Worker 1 might process jobs:

    1, 4, 7

while Worker 2 processes:

    2, 5, 8

and Worker 3 processes:

    3, 6, 9

Or the distribution could be completely different.

IMPORTANT:

Fan-Out means:

    ONE source
        ↓
    MANY workers

This problem is about distributing work among multiple
goroutines.

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

func worker(jobs <-chan int, wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for v := range jobs {
		fmt.Println("Worker", i, "received job", v)
		fmt.Println("Worker", i, "receive:", v*10)
	}
}

func main() {
	wg := sync.WaitGroup{}
	jobs := make(chan int)

	wg.Add(3)
	for i := range 3 {
		go worker(jobs, &wg, i+1)
	}

	go func() {
		for i := 1; i <= 9; i++ {
			jobs <- i
		}

		close(jobs)
	}()

	wg.Wait()

}
