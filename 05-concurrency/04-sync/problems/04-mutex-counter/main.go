/*
============================================================
Problem 4 — Mutex and Shared Counter
============================================================

Create a shared integer:

    counter := 0

Start FIVE goroutines.

Each goroutine must:
    - Lock a sync.Mutex
    - Increment counter by 1
    - Unlock the mutex

Use a sync.WaitGroup so main() waits for all five
goroutines to finish.

After wg.Wait(), print:

    Counter: 5

Requirements:
- Import sync.
- Create a sync.WaitGroup.
- Create a sync.Mutex.
- Use wg.Add(5).
- Start exactly FIVE goroutines.
- Use defer wg.Done() inside each goroutine.
- Use mu.Lock() before modifying counter.
- Use mu.Unlock() after modifying counter.
- Do NOT use channels.
- Do NOT use time.Sleep().
- Do NOT use select.
- Do NOT use atomic operations.

Important:

The counter is shared data.

Without a mutex, multiple goroutines could access and
modify counter concurrently.

We want the increment operation to be protected:

    mu.Lock()
    counter++
    mu.Unlock()

Mental model:

    counter = 0

    Worker 1 ──┐
    Worker 2 ──┤
    Worker 3 ──┤──> Mutex ──> counter++
    Worker 4 ──┤
    Worker 5 ──┘

Only ONE goroutine can hold the mutex at a time.

Goal:
Understand that:

    WaitGroup = "Wait for goroutines"

while:

    Mutex = "Protect shared data"
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

	wg.Add(5)

	for i := range 5 {
		go func() {
			defer wg.Done()
			mu.Lock()
			fmt.Println("Locked goroutine:", i)
			counter++
			mu.Unlock()
			fmt.Println("UnLocked goroutine:", i)

		}()
	}

	wg.Wait()
	fmt.Println("Counter:", counter)
}
