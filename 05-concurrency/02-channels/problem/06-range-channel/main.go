/*
Problem: Range Over a Channel

Create a program where a worker sends multiple values through
a channel and closes the channel when it is finished.

Requirements:

1. Create an unbuffered int channel.

2. Start a worker goroutine.

3. Inside the worker:

		- Send 10.
		- Send 20.
		- Send 30.
		- Close the channel.

4. In main():

		- Use range to receive values from the channel.
		- Print every received value.

5. Do NOT use:

		- time.Sleep
		- sync.WaitGroup
		- value, ok := <-ch
		- buffered channels

Expected output:

	10
	20
	30

Important:

	- range receives values from the channel automatically.
	- The range loop stops when the channel is closed.
	- The worker must close the channel after sending all values.
	- range does NOT close the channel automatically.

Goal:

Understand how range simplifies receiving multiple values
from a channel until the sender is finished.
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

	for v := range ch {
		fmt.Println(v)
	}
}
