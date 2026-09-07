/*
Problem: Buffered Channel

Create a program that demonstrates how a buffered channel
can store values without an immediate receiver.

Requirements:

1. Create a buffered int channel with capacity 3.

2. Send these values into the channel:

		10
		20
		30

3. Before receiving anything, print:

		- len(ch)
		- cap(ch)

4. Receive and print all three values.

Expected output:

	Length: 3
	Capacity: 3
	10
	20
	30

Important:

	- Do NOT use time.Sleep.
	- Do NOT use sync.WaitGroup.
	- Do NOT use close().
	- Do NOT use goroutines for sending.
	- Use a buffered channel.

Goal:

Understand that a buffered channel can temporarily store
values and that sending does not immediately require a receiver
as long as there is free space in the buffer.
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	fmt.Println("Before Send")
	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	fmt.Println("\nAfter Send 10")
	ch <- 10
	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	fmt.Println("\nAfter Send 20")
	ch <- 20
	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	fmt.Println("\nAfter Send 30")
	ch <- 30
	fmt.Println(len(ch))
	fmt.Println(cap(ch))

	fmt.Println("\nReceiving")

	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
