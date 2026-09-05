/*
============================================================
Problem 1 — Basic WaitGroup
============================================================

Create a sync.WaitGroup.

Start TWO goroutines.

Goroutine 1:
    - Print: "Worker 1 started"
    - Print: "Worker 1 finished"

Goroutine 2:
    - Print: "Worker 2 started"
    - Print: "Worker 2 finished"

Use WaitGroup so that main() waits for BOTH goroutines
to finish.

Requirements:
- Import sync.
- Create a sync.WaitGroup.
- Use wg.Add().
- Use wg.Done().
- Use wg.Wait().
- Start exactly TWO goroutines.
- Use defer wg.Done() inside each goroutine.
- Do NOT use time.Sleep().
- Do NOT use channels.
- Do NOT use close().
- Do NOT use select.

Important:

main() must NOT finish until both workers have called
wg.Done().

Mental model:

    wg.Add(2)
        │
        ├──────────> Worker 1
        │              │
        │              └── wg.Done()
        │
        ├──────────> Worker 2
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
Understand that WaitGroup is used to WAIT for goroutines
to finish. It is not used to send data between goroutines.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {

	wg := sync.WaitGroup{}

	wg.Add(2)
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

	wg.Wait()
}
