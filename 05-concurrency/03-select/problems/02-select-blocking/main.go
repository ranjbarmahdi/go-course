/*
============================================================
Problem 2 — select blocks
============================================================

Create one channel:

    ch := make(chan string)

Start a goroutine that:

    1. Waits for 2 seconds using time.Sleep().
    2. Sends "Hello after 2 seconds" to ch.

In main():

    1. Print:
       "Waiting for message..."

    2. Use select to receive from ch.

    3. When the value is received, print it.

Requirements:
- Use a goroutine.
- Use an unbuffered channel.
- Use select.
- You MUST use time.Sleep(2 * time.Second) inside
  the goroutine.
- Do NOT use default.
- Do NOT use WaitGroup.
- Do NOT use close().
- Do NOT use time.After() yet.

Expected behavior:

    Waiting for message...
    [approximately 2 second delay]
    Hello after 2 seconds

Goal:
Understand that select blocks when none of its cases
is currently ready.

Mental model:

    select
       |
       |-- ch ready? --> receive
       |
       |-- not ready --> BLOCK
       |
       |-- eventually ch becomes ready --> receive
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		ch <- "Hello After 2 Seconds"
	}()

	fmt.Println("Waiting for messages...")

	select {
	case value := <-ch:
		fmt.Println(value)
	}

}
