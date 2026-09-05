/*
============================================================
Problem 2 — WaitGroup with Multiple Workers
============================================================

Create a sync.WaitGroup.

Start THREE goroutines.

Goroutine 1:
    - Print: "Worker 1 started"
    - Print: "Worker 1 finished"

Goroutine 2:
    - Print: "Worker 2 started"
    - Print: "Worker 2 finished"

Goroutine 3:
    - Print: "Worker 3 started"
    - Print: "Worker 3 finished"

Use WaitGroup so that main() waits for ALL THREE
goroutines to finish.

Requirements:
- Import sync.
- Create a sync.WaitGroup.
- Use wg.Add(3).
- Use wg.Done().
- Use wg.Wait().
- Start exactly THREE goroutines.
- Use defer wg.Done() inside each goroutine.
- Do NOT use time.Sleep().
- Do NOT use channels.
- Do NOT use close().
- Do NOT use select.

Important:

main() must NOT finish until all three workers have
called wg.Done().

Mental model:

    wg.Add(3)
        │
        ├──────────> Worker 1
        │              │
        │              └── wg.Done()
        │
        ├──────────> Worker 2
        │              │
        │              └── wg.Done()
        │
        ├──────────> Worker 3
        │              │
        │              └── wg.Done()
        │
        ↓
     wg.Wait()
        │
        │ BLOCK
        │
        ↓
    counter = 0
        │
        ↓
    main continues

Goal:
Understand how WaitGroup scales from 2 goroutines
to multiple goroutines.

The number passed to wg.Add() must match the number
of times wg.Done() will be called.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := sync.WaitGroup{}

	wg.Add(3)
	go func() {
		defer wg.Done()
		fmt.Println("Worker 1 started")
		fmt.Println("Worker 1 finished")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Worker 2 started")
		fmt.Println("Worker 2 finished")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Worker 3 started")
		fmt.Println("Worker 3 finished")
	}()

	wg.Wait()
}
