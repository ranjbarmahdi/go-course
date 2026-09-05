/*
============================================================
Problem 3 — defer wg.Done() with Early Return
============================================================

Create a sync.WaitGroup.

Start TWO goroutines.

Worker 1:
    - Print: "Worker 1 started"
    - Return early from the goroutine.
    - It must NOT print "Worker 1 finished"

Worker 2:
    - Print: "Worker 2 started"
    - Print: "Worker 2 finished"

Use WaitGroup so that main() waits for BOTH goroutines.

Requirements:
- Import sync.
- Create a sync.WaitGroup.
- Use wg.Add(2).
- Use wg.Done().
- Use wg.Wait().
- Start exactly TWO goroutines.
- Use defer wg.Done() inside EACH goroutine.
- Worker 1 must use return.
- Do NOT use time.Sleep().
- Do NOT use channels.
- Do NOT use close().
- Do NOT use select.

Important:

Even though Worker 1 returns early, its deferred
wg.Done() MUST still execute.

Mental model:

    Worker 1
        │
        ├── Print "started"
        │
        ├── return
        │
        └── defer wg.Done()
                  ↓
              counter--

    Worker 2
        │
        ├── Print "started"
        ├── Print "finished"
        └── defer wg.Done()
                  ↓
              counter--

Goal:
Understand why `defer wg.Done()` is especially useful
when a goroutine can return from multiple places.
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
		return
	}()

	go func() {
		defer wg.Done()
		fmt.Println("Worker 2 started")
		fmt.Println("Worker 2 finished")
	}()

	wg.Wait()
}
