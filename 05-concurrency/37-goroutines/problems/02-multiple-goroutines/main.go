/*
Problem: Multiple Goroutines

Create a function:

    func printTask(name string)

The function must print the task name 3 times.

For example:

    printTask("Task A")

should print:

    Task A
    Task A
    Task A

In main():

1. Start three goroutines:
       - printTask("Task A")
       - printTask("Task B")
       - printTask("Task C")

2. Print:
       "Main goroutine"

3. Make sure all goroutines have enough time to finish.

Requirements:

- You must use goroutines.
- You must start all three tasks concurrently.
- Do not use channels.
- Do not use sync.WaitGroup.
- You may use time.Sleep().
- Do not assume any specific output order.

Possible output:

    Main goroutine
    Task A
    Task B
    Task C
    Task A
    Task C
    Task B
    Task A
    Task B
    Task C

The exact order may be different.
*/

package main

import (
	"fmt"
	"time"
)

func printTask(name string) {
	for range 3 {
		fmt.Println(name)
	}
}

func main() {
	go printTask("Task A")
	go printTask("Task B")
	go printTask("Task C")

	fmt.Println("Main goroutine")

	time.Sleep(5 * time.Second)
}
