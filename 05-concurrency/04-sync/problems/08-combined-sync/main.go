/*
============================================================
Problem 8 — Combined sync Exercise
============================================================

Create:

    counter := 0

Create:
    - sync.WaitGroup
    - sync.Mutex
    - sync.Once

Start FIVE goroutines.

Each goroutine must:

    1. Run an initialization function using once.Do().
    2. Lock the mutex.
    3. Increment counter by 1.
    4. Unlock the mutex.
    5. Call wg.Done() using defer.

The initialization function must print:

    "Initializing application"

The initialization message must appear EXACTLY ONCE.

After all five goroutines finish, print:

    Counter: 5

Requirements:
- Import sync.
- Create a sync.WaitGroup.
- Create a sync.Mutex.
- Create a sync.Once.
- Use wg.Add(5).
- Start exactly FIVE goroutines.
- Use defer wg.Done() inside every goroutine.
- Use once.Do().
- Use mu.Lock().
- Increment counter by 1.
- Use mu.Unlock().
- Use wg.Wait().
- Do NOT use channels.
- Do NOT use time.Sleep().
- Do NOT use select.
- Do NOT use atomic operations.
- Do NOT use a boolean to implement Once.

Expected output contains:

    Initializing application

exactly once, and ends with:

    Counter: 5

Important:

This exercise demonstrates that the three synchronization
primitives have DIFFERENT responsibilities.

sync.Once:

    "Run initialization exactly once."

sync.Mutex:

    "Protect shared counter."

sync.WaitGroup:

    "Wait until all workers finish."

Mental model:

             FIVE GOROUTINES
                    │
          ┌─────────┼─────────┐
          ↓         ↓         ↓
       once.Do()  once.Do()  ...
          │
          └──────> initialization
                    │
                ONLY ONCE
                    │
                    ↓
               Mutex Lock
                    │
                 counter++
                    │
               Mutex Unlock
                    │
                    ↓
                wg.Done()
                    │
                    ↓
                 wg.Wait()
                    │
                    ↓
              Counter: 5

Goal:

Understand that synchronization primitives can be
combined, but each one solves a different problem:

    Once    → initialization
    Mutex   → shared state
    WaitGroup → completion
*/

package main

import (
	"fmt"
	"sync"
)

func initialization() {
	fmt.Println("initialization application")
}
func main() {
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}
	once := sync.Once{}

	counter := 0

	wg.Add(5)
	for range 5 {
		go func() {
			defer wg.Done()
			once.Do(initialization)
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}

	wg.Wait()

	fmt.Println("Counter:", counter)
}
