/*
Problem: Basic Goroutine

Create a function:

    func printNumbers()

The function must print the numbers 1 through 5.

In main():

1. Start printNumbers() as a goroutine.
2. Print "Main goroutine" from main().
3. Make sure the program stays alive long enough for printNumbers()
   to finish.

Example output:

    Main goroutine
    1
    2
    3
    4
    5

Important:

The exact order of "Main goroutine" and the numbers is NOT guaranteed.

Requirements:

- You must use the `go` keyword.
- Do not use channels.
- Do not use sync.WaitGroup.
- For this problem, you may use time.Sleep() to keep main alive.
*/

package main

import (
	"fmt"
	"time"
)

func printNumbers() {
	for i := range 5 {
		fmt.Println(i + 1)
	}
}

func main() {

	go printNumbers()

	fmt.Println("Main goroutine")

	time.Sleep(1 * time.Second)
}
