package main

import (
	"fmt"
	"sync"
)

// ============================================================
// 40 — sync
// ============================================================
//
// The sync package provides synchronization primitives for
// coordinating concurrent goroutines.
//
// Main tools we will learn:
//
//   1. sync.WaitGroup
//   2. sync.Mutex
//   3. sync.RWMutex
//   4. sync.Once
//
// Important mental model:
//
//   WaitGroup → wait for goroutines to finish
//   Mutex     → protect shared data from concurrent access
//   RWMutex   → multiple readers, exclusive writer
//   Once      → execute something exactly once
//
// Channels are different:
//
//   Channel → goroutines communicate through values
//
// ============================================================

func main() {

	// ============================================================
	// 1. What is sync?
	// ============================================================

	fmt.Println("1. What is sync?")

	/*
		The sync package provides synchronization primitives
		for coordinating goroutines.

		Concurrency creates situations where multiple goroutines
		may need to:

		- wait for each other
		- access shared data safely
		- perform initialization only once

		The sync package helps solve these problems.
	*/

	// ============================================================
	// 2. sync.WaitGroup
	// ============================================================

	fmt.Println("\n2. sync.WaitGroup")

	/*
		A WaitGroup lets one goroutine wait for a collection
		of goroutines to finish.

		Think of WaitGroup as a counter.

		    counter = 0

		wg.Add(1)

		    counter = 1

		wg.Done()

		    counter = 0

		wg.Wait()

		    Wait until counter == 0
	*/

	// ============================================================
	// 3. Basic WaitGroup
	// ============================================================

	fmt.Println("\n3. Basic WaitGroup")

	var wg sync.WaitGroup

	// We are waiting for one goroutine.
	wg.Add(1)

	go func() {
		fmt.Println("Worker finished")

		// Tell WaitGroup that this goroutine is finished.
		wg.Done()
	}()

	// Block until the counter becomes zero.
	wg.Wait()

	fmt.Println("Main finished")

	/*
		Execution:

		    wg.Add(1)
		        │
		        ↓
		    counter = 1
		        │
		        ├──────────────> Worker
		        │                  │
		        │                  ├── work
		        │                  │
		        │                  └── wg.Done()
		        │                        │
		        │                        ↓
		        │                   counter = 0
		        │
		    wg.Wait()
		        │
		        ↓
		    Main continues
	*/

	// ============================================================
	// 4. Why defer wg.Done()?
	// ============================================================

	fmt.Println("\n4. Why defer wg.Done()?")

	/*
		We usually write:

		    go func() {
		        defer wg.Done()

		        // work
		    }()

		Instead of:

		    go func() {
		        // work

		        wg.Done()
		    }()

		defer guarantees that Done() is called when the
		goroutine function returns.

		Mental model:

		    goroutine starts
		          │
		          ↓
		    defer wg.Done()
		    registered
		          │
		          ↓
		         work
		          │
		          ↓
		    function returns
		          │
		          ↓
		      wg.Done()
	*/

	var wg2 sync.WaitGroup

	wg2.Add(1)

	go func() {
		defer wg2.Done()

		fmt.Println("Doing work")
	}()

	wg2.Wait()

	// ============================================================
	// 5. Multiple Goroutines
	// ============================================================

	fmt.Println("\n5. Multiple Goroutines")

	/*
			             Add(3)
			               │
			               ↓
			          counter = 3
			          /     |     \
			         /      |      \
			    Worker1  Worker2  Worker3
			       │        │        │
			    Done()   Done()   Done()
			       │        │        │
			       ↓        ↓        ↓
			       2        1        0
			                         │
			                         ↓
			                    Wait returns
			                         │
			                         ↓
			                All workers finished

		The workers may finish in ANY order.

		WaitGroup only guarantees that Wait() returns after
		all workers have called Done().
	*/

	var wg3 sync.WaitGroup

	wg3.Add(3)

	go func() {
		defer wg3.Done()

		fmt.Println("Worker 1")
	}()

	go func() {
		defer wg3.Done()

		fmt.Println("Worker 2")
	}()

	go func() {
		defer wg3.Done()

		fmt.Println("Worker 3")
	}()

	wg3.Wait()

	fmt.Println("All workers finished")

	// ============================================================
	// 6. Add() must match Done()
	// ============================================================

	fmt.Println("\n6. Add() must match Done()")

	/*
		If we do:

		    wg.Add(3)

		we need three corresponding Done() calls.

		    Add(3)
		      │
		      ↓
		    3
		      │
		    Done()
		      ↓
		    2
		      │
		    Done()
		      ↓
		    1
		      │
		    Done()
		      ↓
		    0
		      │
		      ↓
		    Wait() returns

		If one Done() is missing:

		    Add(3)

		    Done() → 2
		    Done() → 1

		    Wait()
		      ↓
		    BLOCK FOREVER

		The counter never reaches zero.
	*/

	// ============================================================
	// 7. WaitGroup does NOT communicate data
	// ============================================================

	fmt.Println("\n7. WaitGroup does NOT communicate data")

	/*
		WaitGroup answers:

		    "Are all these goroutines finished?"

		It does NOT transfer values between goroutines.

		For communication, use channels:

		    ch := make(chan string)

		    go func() {
		        ch <- "result"
		    }()

		    result := <-ch

		Here the channel is responsible for transferring data.
	*/

	// ============================================================
	// 8. WaitGroup vs Channels
	// ============================================================

	fmt.Println("\n8. WaitGroup vs Channels")

	/*
			                 Concurrency
			                     │
			          ┌──────────┴──────────┐
			          │                     │
			       Channels              sync
			          │                     │
			          │              ┌──────┼──────┐
			          │              │      │      │
			          │         WaitGroup Mutex  Once
			          │                     │
			          │                  RWMutex
			          │
			    Communication
			    + synchronization

		Channel:

		    "Send/receive data between goroutines."

		WaitGroup:

		    "Wait until goroutines finish."

		Mutex:

		    "Protect shared data."

		RWMutex:

		    "Protect shared data with multiple readers."

		Once:

		    "Execute initialization exactly once."
	*/

	// ============================================================
	// 9. sync.Mutex
	// ============================================================

	fmt.Println("\n9. sync.Mutex")

	/*
		A Mutex provides mutual exclusion.

		It allows only one goroutine at a time to enter
		a protected critical section.

		Example:

		    mu.Lock()

		    sharedData++

		    mu.Unlock()

		Mental model:

		    Goroutine 1
		        │
		        ↓
		       Lock
		        │
		        ↓
		   modify shared data
		        │
		        ↓
		      Unlock
		        │
		        ↓
		    Goroutine 2 can enter

		This is useful when multiple goroutines access
		the same shared variable.
	*/

	stock := 12

	var mu sync.Mutex

	mu.Lock()
	stock--
	mu.Unlock()

	fmt.Println("Stock:", stock)

	// ============================================================
	// 10. sync.RWMutex
	// ============================================================

	fmt.Println("\n10. sync.RWMutex")

	/*
		RWMutex stands for Read/Write Mutex.

		It is useful when we have:

		    Many readers
		    Few writers

		For example:

		        shared configuration
		                 │
		           ┌─────┼─────┐
		           ↓     ↓     ↓
		        Reader Reader Reader

		A normal Mutex allows only one goroutine to hold
		the lock at a time.

		RWMutex allows multiple readers to hold a read lock
		at the same time.

		Reader:

		    mu.RLock()
		    defer mu.RUnlock()

		    // read data

		Writer:

		    mu.Lock()
		    defer mu.Unlock()

		    // modify data


		Mental model:

		    RLock   RLock   RLock
		      │       │       │
		      └───────┼───────┘
		           READ TOGETHER


		    BUT


		    Lock
		      │
		      ↓
		    WRITER
		      │
		      ↓
		    Exclusive access
	*/

	var mu2 sync.RWMutex
	var wg4 sync.WaitGroup

	data := 100

	wg4.Add(3)

	// Reader 1
	go func() {
		defer wg4.Done()

		mu2.RLock()
		defer mu2.RUnlock()

		fmt.Println("Reader 1:", data)
	}()

	// Reader 2
	go func() {
		defer wg4.Done()

		mu2.RLock()
		defer mu2.RUnlock()

		fmt.Println("Reader 2:", data)
	}()

	// Writer
	go func() {
		defer wg4.Done()

		mu2.Lock()
		defer mu2.Unlock()

		data = 200

		fmt.Println("Writer:", data)
	}()

	wg4.Wait()

	// ============================================================
	// 11. sync.Once
	// ============================================================

	fmt.Println("\n11. sync.Once")

	/*
		sync.Once guarantees that a function is executed
		exactly once.

		    var once sync.Once

		    once.Do(func() {
		        // initialization
		    })

		Even if multiple goroutines call once.Do(),
		the function executes only once.

		Useful for:

		    - application initialization
		    - configuration loading
		    - lazy initialization
		    - singleton-style initialization
	*/

	var once sync.Once

	once.Do(func() {
		fmt.Println("Initialize")
	})

	once.Do(func() {
		fmt.Println("Initialize")
	})

	// Only the first Do() executes the function.

	// ============================================================
	// Important Mental Model
	// ============================================================

	/*
		WaitGroup
		    ↓
		"Wait until goroutines finish."

		Mutex
		    ↓
		"Only one goroutine accesses this critical section."

		RWMutex
		    ↓
		"Multiple readers, but writers need exclusive access."

		Once
		    ↓
		"Execute this initialization exactly once."

		Channel
		    ↓
		"Goroutines communicate through values."
	*/

	fmt.Println("\n40. sync lecture finished")
}
