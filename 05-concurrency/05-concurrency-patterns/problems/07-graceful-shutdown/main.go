/*
============================================================
Problem 7 — Graceful Shutdown ⭐⭐⭐⭐
============================================================

Build a worker that can shut down gracefully.

Create:

    jobs := make(chan int)
    shutdown := make(chan struct{})

Requirements:

1. Start one worker goroutine.

2. The worker should use select:

    select {
    case job := <-jobs:
        // process job

    case <-shutdown:
        // stop worker
    }

3. Send several jobs to the worker.

4. After some jobs have been sent, signal shutdown:

    close(shutdown)

5. The worker must stop accepting new jobs.

6. Before exiting, print:

    Worker shutting down...

7. Main must wait for the worker to finish using
   WaitGroup.

IMPORTANT:

The shutdown signal should NOT be implemented using
time.Sleep() as synchronization.

The purpose is to understand:

    Running worker
          ↓
    shutdown signal
          ↓
    stop accepting work
          ↓
    cleanup
          ↓
    exit
*/

package main

import "fmt"

func main() {
	fmt.Println("graceful shutdown")
}
