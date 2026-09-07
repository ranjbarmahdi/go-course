/*
============================================================
Problem 6 — RWMutex: Multiple Readers + One Writer
============================================================

Create a shared integer:

    data := 100

Create:
    - sync.RWMutex
    - sync.WaitGroup

Start THREE reader goroutines.

Each reader must:
    - Use RLock()
    - Read and print data
    - Use RUnlock()

Start ONE writer goroutine.

The writer must:
    - Use Lock()
    - Change data to 200
    - Use Unlock()

Use WaitGroup so main() waits for ALL FOUR goroutines.

Requirements:
- Import sync.
- Create a sync.RWMutex.
- Create a sync.WaitGroup.
- Use wg.Add(4).
- Start exactly THREE readers.
- Start exactly ONE writer.
- Use defer wg.Done() inside every goroutine.
- Readers must use RLock() / RUnlock().
- Writer must use Lock() / Unlock().
- Use wg.Wait().
- Do NOT use channels.
- Do NOT use time.Sleep().
- Do NOT use select.
- Do NOT use sync.Mutex.

Important:

RWMutex allows:

    Multiple readers at the same time.

But:

    Only one writer at a time.

And while a writer holds the lock:

    Readers must wait.

Reader:

    mu.RLock()
    fmt.Println(data)
    mu.RUnlock()

Writer:

    mu.Lock()
    data = 200
    mu.Unlock()

Mental model:

                 RWMutex
                    │
          ┌─────────┴─────────┐
          ↓                   ↓
      Readers              Writer
     RLock()                Lock()
          │                   │
   multiple allowed       exclusive
          │                   │
      RUnlock()            Unlock()

Goal:
Understand the difference between:

    sync.Mutex
        ↓
    One goroutine at a time.

and:

    sync.RWMutex
        ↓
    Multiple readers OR one writer.
*/

package main

import (
	"fmt"
	"sync"
)

func main() {
	mu := sync.RWMutex{}
	wg := sync.WaitGroup{}

	data := 100

	wg.Add(4)
	for i := range 3 {
		go func() {
			defer wg.Done()

			mu.RLock()
			fmt.Println("Reader", i, "Locked")

			fmt.Println("Data:", data)

			fmt.Println("Reader", i, "UnLocked")
			mu.RUnlock()
		}()
	}

	go func() {
		defer wg.Done()

		mu.Lock()
		fmt.Println("Writer Locked")
		data = 200
		fmt.Println("Writer UnLocked")
		mu.Unlock()
	}()

	wg.Wait()
}
