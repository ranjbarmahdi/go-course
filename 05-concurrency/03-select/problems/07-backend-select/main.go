/*
============================================================
Problem 7 — Backend-style select
============================================================

Imagine a backend endpoint that needs to fetch data from
two independent sources:

    User Service
    Order Service

Create two buffered channels:

    userCh := make(chan string, 1)
    orderCh := make(chan string, 1)

Start two goroutines.

User goroutine:
    - Sleep for 1 second.
    - Send "User: Mehdi" to userCh.

Order goroutine:
    - Sleep for 3 seconds.
    - Send "Orders: 5" to orderCh.

In main(), use select with THREE cases:

    case user := <-userCh:
        // print the user

    case orders := <-orderCh:
        // print the orders

    case <-time.After(2 * time.Second):
        // print "Request timeout"

Requirements:
- Use two buffered channels.
- Capacity must be 1.
- Use two goroutines.
- User takes 1 second.
- Orders take 3 seconds.
- Timeout is 2 seconds.
- Use select.
- Use time.After().
- Do NOT use WaitGroup.
- Do NOT use default.
- Do NOT use close().
- Do NOT use time.Sleep() in main().
- time.Sleep() should only be used inside the goroutines.

Expected output:

    User: Mehdi

Why?

    User Service
        │
        │ 1 second
        ↓
      userCh ──────────┐
                       │
    Timeout ───────────┤ 2 seconds
                       ↓
                     select
                       ↑
    Order Service ─────┤
        │              │
        │ 3 seconds     │
        ↓              │
      orderCh ─────────┘

The user result arrives first, so select receives it
and main() finishes.

Important backend concept:

select allows a server to wait for multiple asynchronous
operations and also define a maximum amount of time it is
willing to wait.

This pattern is commonly useful for:

    - database queries
    - external APIs
    - microservice calls
    - cache operations
    - request timeouts
    - cancellation

Do not worry about implementing a real HTTP server yet.
Focus only on understanding select.
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
		time.Sleep(time.Second)
		userCh <- "User: Mahdi"
	}()

	go func() {
		time.Sleep(time.Second * 3)
		orderCh <- "Orders: 5"
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
