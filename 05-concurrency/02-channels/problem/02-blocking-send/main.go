/*
Problem: Blocking Send

Create a program that demonstrates that sending to an
unbuffered channel blocks until another goroutine receives
the value.

Requirements:

1. Create an unbuffered int channel.

2. Start a goroutine that:

		- Prints "Worker: before send"
		- Sends 100 into the channel.
		- Prints "Worker: after send"

3. In main():

		- Print "Main: before receive"
		- Receive the value from the channel.
		- Print the received value.
		- Print "Main: after receive"

Expected output should demonstrate that:

	"Worker: after send"

cannot happen until the receiver receives the value.

Important:

	Do NOT use time.Sleep.
	Do NOT use WaitGroup.
	Do NOT use a buffered channel.
	Do NOT use close().

Goal:

Understand exactly where an unbuffered channel send
blocks and how the sender and receiver synchronize.
*/

package main

import "fmt"

func worker(ch chan int) {
	fmt.Println("Worker: before send")
	ch <- 100
	fmt.Println("Worker: after send")
}

func main() {
	ch := make(chan int)

	go worker(ch)

	fmt.Println("Main: before receive")

	value := <-ch

	fmt.Println("Received:", value)
	fmt.Println("Main: after receive")
}
