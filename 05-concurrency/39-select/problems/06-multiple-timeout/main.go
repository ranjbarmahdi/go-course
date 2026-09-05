/*
============================================================
Problem 6 — Multiple channels + timeout
============================================================

Create two buffered channels:

    userCh := make(chan string, 1)
    orderCh := make(chan string, 1)

Start two goroutines.

Goroutine 1:
    - Sleep for 1 second.
    - Send "User data" to userCh.

Goroutine 2:
    - Sleep for 3 seconds.
    - Send "Order data" to orderCh.

In main(), use select with THREE cases:

    case user := <-userCh:
        // print the user data

    case order := <-orderCh:
        // print the order data

    case <-time.After(2 * time.Second):
        // print "Request timeout"

Requirements:
- Use two goroutines.
- Use two buffered channels.
- Both channels must have capacity 1.
- Use select.
- Use time.After().
- userCh sends after 1 second.
- orderCh sends after 3 seconds.
- timeout happens after 2 seconds.
- Do NOT use WaitGroup.
- Do NOT use default.
- Do NOT use close().

Expected output:

    User data

Why?

    userCh  ───────────→ 1 second
    timeout ───────────→ 2 seconds
    orderCh ───────────→ 3 seconds

The user data arrives first, so select receives from
userCh and main() finishes.

Important:
select does NOT wait for all channels.

It waits until ONE case is ready.

    userCh  ──┐
              │
    orderCh ──┼──> select ──> FIRST READY CASE
              │
    timeout ──┘
*/

package main

import (
	"fmt"
	"time"
)

func main() {
	userCh := make(chan string, 1)
	orderCh := make(chan string, 1)

	go func() {
		time.Sleep(1 * time.Second)
		userCh <- "User data"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		orderCh <- "Order data"
	}()

	select {
	case user := <-userCh:
		fmt.Println(user)

	case order := <-orderCh:
		fmt.Println(order)

	case <-time.After(2 * time.Second):
		fmt.Println("Timeout")
	}
}
