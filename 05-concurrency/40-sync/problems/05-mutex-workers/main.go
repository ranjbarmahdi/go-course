/*
============================================================
Problem 5 — Mutex with Multiple Workers
============================================================

Create a shared integer:

    counter := 0

Start TEN goroutines.

Each goroutine must increment the shared counter
TEN times.

So the final result must be:

    Counter: 100

Use:
    - sync.WaitGroup
    - sync.Mutex

Requirements:
- Import sync.
- Create a sync.WaitGroup.
- Create a sync.Mutex.
- Use wg.Add(10).
- Start exactly TEN goroutines.
- Each goroutine must increment counter TEN times.
- Protect every counter increment with the mutex.
- Use defer wg.Done() inside each goroutine.
- Use wg.Wait().
- Do NOT use channels.
- Do NOT use time.Sleep().
- Do NOT use select.
- Do NOT use atomic operations.

The critical section should be:

    mu.Lock()
    counter++
    mu.Unlock()

Important:

The mutex must protect EACH access that modifies the
shared counter.

Mental model:

    counter = 0

    Worker 1 ──┐
    Worker 2 ──┤
    Worker 3 ──┤
       ...     ├──> Mutex ──> counter++
       ...     │
    Worker 10 ─┘

Each worker:

    counter++
    counter++
    counter++
       ...
    counter++    // 10 times

Total:

    10 workers × 10 increments = 100

Goal:
Understand that Mutex protects shared state while
WaitGroup waits for all workers to finish.

Important distinction:

    WaitGroup → "Are all workers finished?"

    Mutex → "Who can modify the shared data right now?"
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	counter := 0

	wg.Add(10)

	for i := range 10 {
		go func() {
			defer wg.Done()

			mu.Lock()
			fmt.Println("Goroutine", i, "Locked")

			for range 10 {
				counter++
			}

			mu.Unlock()
			fmt.Println("Goroutine", i, "UnLocked")

		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
