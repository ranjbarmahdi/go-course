/*
Problem: Closing a Channel

Create a program where a worker sends multiple values and then
closes the channel to tell the receiver that no more values will
be sent.

Requirements:

1. Create an unbuffered int channel.

2. Start a worker goroutine.

3. Inside the worker:

		- Send 10.
		- Send 20.
		- Send 30.
		- Close the channel.

4. In main():

		- Receive values from the channel using:

			value, ok := <-ch

		- Print both value and ok for every receive.
		- Continue receiving until the channel is closed.

5. Do NOT use:

		- time.Sleep
		- sync.WaitGroup
		- range
		- buffered channels

Expected output:

	10 true
	20 true
	30 true
	0 false

Goal:

Understand:

	- Why a sender closes a channel.
	- What close(ch) means.
	- What the second value (ok) means when receiving.
	- What happens when receiving from a closed channel.
*/

package main

import "fmt"

func main() {
	ch := make(chan int)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30

		close(ch)
	}()

	v, ok := <-ch
	fmt.Println(v, ok)

	v, ok = <-ch
	fmt.Println(v, ok)

	v, ok = <-ch
	fmt.Println(v, ok)

	v, ok = <-ch
	fmt.Println(v, ok)
}
