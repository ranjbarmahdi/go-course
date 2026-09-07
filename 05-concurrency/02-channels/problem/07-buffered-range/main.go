/*
Problem: Buffered Channel + Range

Create a program where a worker sends multiple values through
a buffered channel, closes the channel, and the main goroutine
receives all values using range.

Requirements:

1. Create a buffered int channel with capacity 3.

2. Start a worker goroutine.

3. Inside the worker:

		- Send 10.
		- Send 20.
		- Send 30.
		- Print "Worker: all values sent".
		- Close the channel.

4. In main():

		- Use range to receive all values.
		- Print every received value.
		- After the range loop finishes, print:

			"Main: channel closed"

5. Do NOT use:

		- time.Sleep
		- sync.WaitGroup
		- value, ok := <-ch

Expected output:

	10
	20
	30
	Worker: all values sent
	Main: channel closed

Important:

	- The channel must have capacity 3.
	- The worker must close the channel.
	- range must receive the values.
	- range stops after the channel is closed and all values
	  have been received.
	- Do NOT manually receive individual values.

Goal:

Understand how buffered channels, close(), and range work
together.

Think about this flow:

	Worker
	   │
	   ├── send 10
	   ├── send 20
	   ├── send 30
	   ├── close(ch)
	   │
	   ▼
	Channel
	   │
	   ▼
	Main
	   │
	   └── range over channel
*/

package main

import "fmt"

func main() {
	ch := make(chan int, 3)

	go func() {
		ch <- 10
		ch <- 20
		ch <- 30
		ch <- 40
		close(ch)
	}()

	for v := range ch {
		fmt.Println(v)
	}
}
