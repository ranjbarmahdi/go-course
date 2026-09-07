package main

import (
	"fmt"
	"time"
)

// ============================================================
// 37 — GOROUTINES ⭐⭐⭐
// ============================================================
//
// A goroutine is a lightweight unit of execution managed by
// the Go runtime.
//
// Start a function as a goroutine:
//
//     go function()
//
// Important:
//
//     goroutine != OS thread
//
// Goroutines are managed by the Go runtime and scheduled onto
// operating-system threads.
//
// In this lesson:
//
// 1. Concurrency
// 2. What is a goroutine?
// 3. Starting a goroutine
// 4. main() is also a goroutine
// 5. Multiple goroutines
// 6. Why main() can terminate goroutines
// 7. Goroutine lifecycle
// 8. Anonymous goroutines
// 9. Passing arguments
// 10. Concurrency vs parallelism
// 11. Goroutines in backend development
// 12. Important mental model
//
// ============================================================

// ============================================================
// 1. Concurrency ⭐⭐⭐
// ============================================================
//
// Imagine a backend needs to perform three independent tasks:
//
//     Task A → process user
//     Task B → send email
//     Task C → write log
//
// Sequential execution:
//
//     A ──────────>
//                  B ──────────>
//                               C ──────────>
//
// Each task waits for the previous task to finish.
//
// With concurrency, multiple tasks can make progress during
// overlapping periods:
//
//     A ───────────────>
//     B ─────────>
//     C ─────────────>
//
// Important:
//
//     Concurrency != Parallelism
//
// Concurrency:
//     Multiple tasks can make progress during the same period.
//
// Parallelism:
//     Multiple tasks are literally executing at the same time,
//     usually on different CPU cores.
//
// Go is designed to make concurrent programming easy.
//

// ============================================================
// 2. What Is a Goroutine? ⭐⭐⭐
// ============================================================
//
// A goroutine is a lightweight unit of execution managed by
// the Go runtime.
//
// Normal function:
//
//     sayHello()
//
// The current goroutine calls the function and waits for it
// to finish.
//
// Goroutine:
//
//     go sayHello()
//
// A new goroutine is started and the current goroutine
// continues executing.
//
// Think:
//
//     sayHello()
//
//         "Run this and wait."
//
//
//     go sayHello()
//
//         "Run this concurrently and let me continue."
//

func sayHello() {
	fmt.Println("Hello from goroutine")
}

// ============================================================
// 3. Starting a Goroutine
// ============================================================
//
// The `go` keyword is used to start a goroutine.
//
// Example:
//
//     go sayHello()
//
// The current goroutine does NOT wait for sayHello() to finish.
//
// The Go runtime schedules the new goroutine for execution.
//

func firstGoroutine() {
	fmt.Println("Start")

	go sayHello()

	fmt.Println("End")

	// Sleep is used here only so we can observe the
	// goroutine executing.
	time.Sleep(100 * time.Millisecond)
}

// ============================================================
// 4. main() Is Also a Goroutine ⭐⭐⭐
// ============================================================
//
// When a Go program starts, main() executes inside the
// main goroutine.
//
// Conceptually:
//
//                 Go Runtime
//                      │
//              ┌───────┴────────┐
//              │                │
//        main goroutine     new goroutine
//              │                │
//            main()           sayHello()
//
// When we write:
//
//     go sayHello()
//
// another goroutine is created.
//

func mainGoroutineExample() {
	fmt.Println("Main goroutine is running")

	go sayHello()

	time.Sleep(100 * time.Millisecond)
}

// ============================================================
// 5. Multiple Goroutines ⭐⭐⭐
// ============================================================
//
// We can start many goroutines:
//
//     go task("Task A")
//     go task("Task B")
//     go task("Task C")
//
// They execute concurrently.
//
// The order of execution is NOT guaranteed.
//

func task(name string) {
	for i := 1; i <= 3; i++ {
		fmt.Println(name, i)
		time.Sleep(100 * time.Millisecond)
	}
}

func multipleGoroutines() {
	go task("Task A")
	go task("Task B")
	go task("Task C")

	// Only used here to keep main alive long enough
	// to observe the goroutines.
	time.Sleep(1 * time.Second)
}

/*
Possible output:

    Task A 1
    Task B 1
    Task C 1
    Task B 2
    Task A 2
    Task C 2
    Task A 3
    Task C 3
    Task B 3

Another execution may produce a completely different order.

Never rely on a specific execution order between goroutines.
*/

// ============================================================
// 6. Why main() Can Terminate Goroutines ⭐⭐⭐
// ============================================================
//
// This is one of the most important concepts.
//
// Consider:
//
//     func main() {
//         go sayHello()
//         fmt.Println("Main finished")
//     }
//
// The execution can be:
//
//     main goroutine
//          │
//          ├── start sayHello goroutine
//          │
//          ├── print "Main finished"
//          │
//          └── main() returns
//                   │
//                   ▼
//             PROGRAM EXITS
//
// If the program exits before the goroutine executes,
// the goroutine will not finish.
//
// The Go runtime does NOT automatically wait for all
// goroutines when main() returns.
//
// Therefore:
//
//     main() returns
//          ↓
//     program terminates
//
// This is why we temporarily use time.Sleep() in these
// examples.
//
// Later we will learn proper synchronization using:
//
//     sync.WaitGroup
//     channels
//     mutexes
//

func mainCanFinishFirst() {
	go func() {
		fmt.Println("Background task started")

		time.Sleep(1 * time.Second)

		fmt.Println("Background task finished")
	}()

	fmt.Println("Main finished")

	// Without this Sleep, the program may terminate
	// before the goroutine finishes.
	time.Sleep(2 * time.Second)
}

// ============================================================
// 7. Goroutine Lifecycle ⭐⭐⭐
// ============================================================
//
// A goroutine can be thought of as moving through several
// states:
//
//     Created
//        │
//        ▼
//     Runnable
//        │
//        ▼
//      Running
//        │
//     ┌──┴────┐
//     │       │
//     ▼       ▼
//  Waiting  Running
//     │       │
//     └──┬────┘
//        │
//        ▼
//     Finished
//
// A simplified lifecycle:
//
//     Create goroutine
//          ↓
//       Runnable
//          ↓
//     Scheduler runs it
//          ↓
//       Function executes
//          ↓
//     Function returns
//          ↓
//     Goroutine finishes
//
// A goroutine can temporarily wait because it is blocked on
// things such as:
//
//     I/O
//     channels
//     mutexes
//     timers
//
// We will understand these blocking situations better when
// we learn channels and synchronization.
//

// ============================================================
// 8. Anonymous Goroutines
// ============================================================
//
// We don't always need to create a separate function.
//
// We can create an anonymous function and immediately run it
// as a goroutine:
//
//     go func() {
//         // work
//     }()
//
// This is very common in Go.
//

func anonymousGoroutine() {
	go func() {
		fmt.Println("Hello from anonymous goroutine")
	}()

	time.Sleep(100 * time.Millisecond)
}

// ============================================================
// 9. Passing Arguments to Goroutines
// ============================================================
//
// We can pass arguments to a goroutine just like a normal
// function call.
//

func greet(name string) {
	fmt.Println("Hello", name)
}

func goroutineArguments() {
	go greet("Mahdi")

	time.Sleep(100 * time.Millisecond)
}

/*
We can also use an anonymous function:

    name := "Mahdi"

    go func(value string) {
        fmt.Println("Hello", value)
    }(name)

The value is passed into the function when the goroutine
is started.
*/

// ============================================================
// 10. Concurrency vs Parallelism ⭐⭐⭐
// ============================================================
//
// Concurrency:
//
//     Multiple tasks can make progress during the same period.
//
// Example:
//
//     Task A: ----work----wait----work----
//     Task B: --work----wait----work------
//
//
//
// Parallelism:
//
//     Multiple tasks are executing literally at the same time.
//
//     CPU Core 1 → Task A
//     CPU Core 2 → Task B
//
//
//
// Go supports concurrency.
//
// Whether goroutines execute in parallel depends on the
// Go runtime scheduler, available CPU resources, and runtime
// configuration.
//
// The important idea:
//
//     Goroutine = concurrency
//
//     Parallel execution = possible, but not guaranteed
//     simply because you created a goroutine.
//

// ============================================================
// 11. Goroutines in Backend Development ⭐⭐⭐
// ============================================================
//
// Goroutines are extremely useful in backend applications.
//
// Examples:
//
//     HTTP request handling
//     Database operations
//     External API calls
//     Background jobs
//     Message processing
//     File processing
//     Periodic tasks
//
// Imagine an HTTP request:
//
//                 HTTP Request
//                      │
//              ┌───────┼────────┐
//              │       │        │
//              ▼       ▼        ▼
//            User    Orders   Notifications
//
// These operations may be independent and can potentially
// be handled concurrently.
//
// But goroutines alone don't solve communication or
// synchronization.
//
// Once multiple goroutines need to:
//
//     communicate
//     wait for each other
//     share data
//     coordinate work
//
// we need additional concurrency tools.
//
// We will learn:
//
//     Channels
//     Select
//     Mutex
//     WaitGroup
//     Once
//
// in the next topics.
//

// ============================================================
// 12. Important Mental Model ⭐⭐⭐
// ============================================================
//
// Remember these:
//
// 1. A goroutine is a lightweight unit of concurrent execution.
//
// 2. Start a goroutine with:
//
//        go function()
//
// 3. main() runs inside the main goroutine.
//
// 4. main() returning terminates the entire program.
//
// 5. Other goroutines are NOT automatically waited for.
//
// 6. The execution order of goroutines is not guaranteed.
//
// 7. time.Sleep() is useful for demonstrations,
//    but it is NOT proper synchronization.
//
// 8. Proper synchronization will come later with:
//
//        WaitGroup
//        Channels
//        Mutex
//
//
//
// The most important difference:
//
//     function()
//
//         Current goroutine
//              │
//              ▼
//           function
//              │
//              ▼
//            wait
//              │
//              ▼
//           continue
//
//
//     go function()
//
//         Current goroutine
//              │
//              ├──────────────► continue
//              │
//              ▼
//         New goroutine
//              │
//              ▼
//           function
//
//
// Mental model:
//
//     go function()
//
//     =
//
//     "Go runtime, start this function concurrently.
//      The current goroutine can continue."
//
// ============================================================

func main() {
	fmt.Println("\n1. Concurrency")

	fmt.Println("\n2. What is a goroutine?")
	go sayHello()
	time.Sleep(100 * time.Millisecond)

	fmt.Println("\n3. Starting a goroutine")
	firstGoroutine()

	fmt.Println("\n4. main() is also a goroutine")
	mainGoroutineExample()

	fmt.Println("\n5. Multiple goroutines")
	multipleGoroutines()

	fmt.Println("\n6. Why main() can terminate goroutines")
	mainCanFinishFirst()

	fmt.Println("\n8. Anonymous goroutines")
	anonymousGoroutine()

	fmt.Println("\n9. Passing arguments to goroutines")
	goroutineArguments()

	fmt.Println("\n37 — Goroutines finished")
}
