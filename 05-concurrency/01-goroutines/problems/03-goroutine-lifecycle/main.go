/*
Problem: Goroutine Lifecycle

Create a function:

    func worker()

Inside worker():

1. Print:
       "Worker started"

2. Wait for 500 milliseconds.

3. Print:
       "Worker finished"

In main():

1. Start worker() as a goroutine.

2. Print:
       "Main finished"

3. Make sure the worker has enough time to finish.

Expected output should contain:

    Worker started
    Main finished
    Worker finished

However, remember:

The exact order of "Worker started" and "Main finished"
is NOT guaranteed.

Requirements:

- Use a goroutine.
- Use time.Sleep().
- Do not use channels.
- Do not use sync.WaitGroup.
- Understand why main() must stay alive long enough for
  the worker goroutine to finish.

Bonus question:

What happens if you remove the final time.Sleep() from main()?
Explain why.
*/

package main

import (
	"fmt"
	"time"
)

func worker() {
	fmt.Println("Worker started")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Worker finished")
}

func main() {
	go worker()

	fmt.Println("Main goroutine")

	time.Sleep(2 * time.Second)
}
