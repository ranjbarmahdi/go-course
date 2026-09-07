/*
Problem: Concurrent Backend Tasks

Imagine a backend request that needs to perform three independent tasks:

    1. Fetch user data
    2. Fetch orders
    3. Fetch notifications

Create these three functions:

    func fetchUser()
    func fetchOrders()
    func fetchNotifications()

Each function must:

1. Print "<task> started"
2. Sleep for 500 milliseconds.
3. Print "<task> finished"

For example:

    User started
    User finished

In main():

1. Start all three functions as goroutines.
2. Print:
       "Request processing..."

3. Keep the program alive long enough for all three
   goroutines to finish.

Important:

If the functions were executed normally:

    fetchUser()
    fetchOrders()
    fetchNotifications()

the total execution time would be approximately:

    500ms + 500ms + 500ms = 1500ms

When executed as goroutines, they can run concurrently,
so the total time should be closer to:

    ~500ms

Requirements:

- Use goroutines.
- Use time.Sleep().
- Do not use channels.
- Do not use sync.WaitGroup.
- Do not assume a specific output order.

Bonus:

Measure the total execution time using:

    time.Now()
    time.Since()

and compare the concurrent version with the sequential version.
*/

package main

import (
	"fmt"
	"time"
)

func fetchUser() {
	fmt.Println("User Started")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("User finished")
}

func fetchOrders() {
	fmt.Println("Order Started")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Order finished")
}

func fetchNotifications() {
	fmt.Println("Notification Started")
	time.Sleep(500 * time.Millisecond)
	fmt.Println("Notification finished")
}

func main() {
	now := time.Now()

	go fetchUser()
	go fetchOrders()
	go fetchNotifications()

	fmt.Println("Main goroutine")

	time.Sleep(time.Second)

	diff := time.Since(now)
	fmt.Printf("Take %d Milliseconds", diff.Milliseconds())
}
