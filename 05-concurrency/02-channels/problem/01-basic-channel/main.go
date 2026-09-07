/*
Problem: Basic Channel Communication

Create a worker function that communicates with the main goroutine
using an unbuffered channel.

Requirements:

1. Create a function:

		func worker(ch chan int)

2. Inside worker():

		- Send 10 into the channel.
		- Send 20 into the channel.

3. In main():

		- Create an unbuffered int channel.
		- Start worker() as a goroutine.
		- Receive both values from the channel.
		- Print both values.

Expected output:

	10
	20

Restrictions:

	- Do NOT use time.Sleep.
	- Do NOT use sync.WaitGroup.
	- Do NOT use close().
	- Do NOT use a buffered channel.
	- Only use goroutines and an unbuffered channel.

Goal:

Understand basic communication between a goroutine and the
main goroutine using an unbuffered channel.
*/

package main

import "fmt"

func worker(ch chan int) {
	ch <- 10
	ch <- 20
}

func main() {
	ch := make(chan int)

	go worker(ch)

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
