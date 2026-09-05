/*
============================================================
Problem 5 — Timeout with select
============================================================

Create a buffered channel:

    ch := make(chan string, 1)

Start a goroutine that:

    1. Sleeps for 2 seconds.
    2. Sends "Data received" to ch.

In main():

    1. Use select with TWO cases:

       case value := <-ch:
           // print the received value

       case <-time.After(1 * time.Second):
           // print "Timeout"

Requirements:
- Use a goroutine.
- Use a buffered channel with capacity 1.
- Use select.
- Use time.After().
- The worker must sleep for 2 seconds.
- The timeout must be 1 second.
- Do NOT use WaitGroup.
- Do NOT use close().
- Do NOT use default.

Expected output:

    Timeout

Why?

    Worker:
        sleep 2 seconds
              ↓
        send "Data received"

    select:
        ch --------------------> 2 seconds
        time.After(1 second) --> 1 second

    The timeout happens FIRST.

Important:
The buffered channel is intentional.

After the timeout happens, main() finishes.
The worker may later wake up and send its value.

Because the channel has capacity 1, that send can complete
without getting stuck waiting for a receiver.
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string, 1)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "Data Received"
	}()

	select {
	case v := <-ch:
		fmt.Println(v)
	case <-time.After(time.Second):
		fmt.Println("Timeout")
	}
}
