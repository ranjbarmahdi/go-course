/*
============================================================
Problem 4 — Multiple ready channels
============================================================

Create two buffered channels:

    ch1 := make(chan int, 1)
    ch2 := make(chan int, 1)

Put a value into BOTH channels:

    ch1 <- 10
    ch2 <- 20

Now use select:

    case value := <-ch1:
        // print the value

    case value := <-ch2:
        // print the value

Requirements:
- Use two channels.
- Both channels must be buffered.
- Capacity must be 1.
- Both channels must contain a value BEFORE select.
- Use select.
- Do NOT use goroutines.
- Do NOT use default.
- Do NOT use time.Sleep().
- Do NOT use time.After().
- Do NOT use WaitGroup.
- Do NOT use close().

Expected output is either:

    Received: 10

or:

    Received: 20

Important:
Both cases are ready at the same time.

Go does NOT guarantee that ch1 will be selected first.

The purpose of this exercise is to understand:

    ch1 ready ──┐
                ├──> select ──> ONE case
    ch2 ready ──┘

If both cases are ready, select chooses one.
*/

package main

import (
	"fmt"
)

func main() {
	ch1 := make(chan int, 1)
	ch2 := make(chan int, 1)

	ch1 <- 10
	ch2 <- 20

	select {
	case v := <-ch1:
		fmt.Println(v)
	case v := <-ch2:
		fmt.Println(v)
	}

}
