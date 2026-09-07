/*
Problem: Buffered Channel Blocking

Create a program that demonstrates that a send to a buffered
channel blocks when the buffer is full.

Requirements:

1. Create a buffered int channel with capacity 2.

2. Start a goroutine that:

  - Prints "Worker: before sends"
  - Sends 10 into the channel.
  - Sends 20 into the channel.
  - Prints "Worker: buffer is full"
  - Sends 30 into the channel.
  - Prints "Worker: after third send"

3. In main():

  - Wait for the first two values by receiving them
    from the channel.
  - Print both received values.
  - Then receive the third value.
  - Print it.

Important:

  - Do NOT use time.Sleep.
  - Do NOT use sync.WaitGroup.
  - Do NOT use close().
  - Use a buffered channel with capacity 2.

Goal:

Understand that:

  - A buffered channel can accept values until its buffer is full.
  - Once the buffer is full, the next send blocks.
  - Receiving a value creates space in the buffer.
  - The blocked sender can then continue.

Expected behavior:

	Worker sends 10 and 20 successfully.

	Worker blocks when trying to send 30.

	Main receives values from the channel.

	The receive creates space, allowing the worker to send 30.
*/
package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 2)

	go func() {
		fmt.Println("Worker: before sends")
		ch <- 10
		ch <- 20
		fmt.Println("Worker: buffer is full")

		ch <- 30
		fmt.Println("Worker: after third send")

		ch <- 40
		fmt.Println("Worker: after fourth send")
	}()

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)

	time.Sleep(time.Second)
	fmt.Println(<-ch)
}
