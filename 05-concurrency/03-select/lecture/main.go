package main

import (
	"fmt"
	"time"
)

// ============================================================
// Helper functions
// ============================================================

func fetchUser(userCh chan<- string) {
	userCh <- "user"
}

func fetchOrders(orderCh chan<- string) {
	orderCh <- "orders"
}

func main() {

	// ============================================================
	// 1. What is select?
	// ============================================================

	fmt.Println("1. What is select?")

	/*
		select is essentially the channel equivalent of switch.

		switch chooses between conditions:

			switch value {
			case 1:
				...
			case 2:
				...
			}

		select chooses between channel operations:

			select {
			case value := <-ch1:
				fmt.Println(value)

			case value := <-ch2:
				fmt.Println(value)
			}
		}

		Think:

			switch → Which condition matches?

			select → Which channel operation is ready?
	*/

	// ============================================================
	// 2. Why do we need select?
	// ============================================================

	fmt.Println("\n2. Why do we need select?")

	/*
		Imagine we have two goroutines:

		Goroutine A ──► ch1
		                 │
		                 ▼
		               main

		Goroutine B ──► ch2
		                 │
		                 ▼
		               main


		We don't know which goroutine will send first.

		Without select:

			value := <-ch1

		main is now waiting specifically for ch1.

		Even if ch2 is ready, main cannot receive from ch2
		until it finishes waiting for ch1.

		select allows us to wait for both.
	*/

	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	go func() {
		ch1 <- "from channel 1"
	}()

	go func() {
		ch2 <- "from channel 2"
	}()

	select {
	case value := <-ch1:
		fmt.Println("ch1:", value)

	case value := <-ch2:
		fmt.Println("ch2:", value)
	}

	/*
		Whichever channel is ready first can be selected.

		Important:

		We cannot rely on ch1 always winning
		or ch2 always winning.
	*/

	// ============================================================
	// 3. select blocks ⭐⭐⭐
	// ============================================================

	fmt.Println("\n3. select blocks")

	/*
		If there is no default case and no channel operation
		is ready, select blocks.

			select {
			case value := <-ch1:
				...

			case value := <-ch2:
				...
			}

		If neither channel is ready:

			ch1 ❌
			ch2 ❌

			  ↓

			select BLOCKS

		As soon as one operation becomes ready,
		select can continue.
	*/

	// ============================================================
	// 4. Multiple ready channels ⭐⭐⭐
	// ============================================================

	fmt.Println("\n4. Multiple ready channels")

	ch4a := make(chan int, 1)
	ch4b := make(chan int, 1)

	// Both channels are already ready.
	ch4a <- 10
	ch4b <- 20

	select {
	case value := <-ch4a:
		fmt.Println("Received from ch4a:", value)

	case value := <-ch4b:
		fmt.Println("Received from ch4b:", value)
	}

	/*
		When multiple cases are ready at the same time,
		select chooses one of them.

		Do NOT depend on a specific case always being selected.

		For example, don't assume:

			ch4a always wins

		or:

			ch4b always wins
	*/

	// Receive the remaining value so the example is complete.
	select {
	case value := <-ch4a:
		fmt.Println("Remaining:", value)

	case value := <-ch4b:
		fmt.Println("Remaining:", value)
	}

	// ============================================================
	// 5. default — non-blocking select ⭐⭐⭐
	// ============================================================

	fmt.Println("\n5. default")

	ch5 := make(chan int)

	go func() {
		ch5 <- 10
	}()

	select {
	case value := <-ch5:
		fmt.Println("Received:", value)

	default:
		fmt.Println("No value available")
	}

	/*
		default changes the behavior.

		Without default:

			select waits.

		With default:

			select does not wait.

		If no channel operation is ready,
		default executes immediately.

		Important:

		The result of this example is timing-dependent.

		The goroutine may or may not have reached
		the send before select runs.

		Therefore both of these are possible:

			Received: 10

		or:

			No value available
	*/

	// ============================================================
	// 6. Timeout with select ⭐⭐⭐
	// ============================================================

	fmt.Println("\n6. Timeout with select")

	/*
		One of the most important uses of select
		is implementing timeouts.

		time.After(duration)

		returns a channel that becomes ready after
		the specified duration.
	*/

	ch6 := make(chan string, 1)

	go func() {
		// Simulate slow work.
		time.Sleep(2 * time.Second)

		ch6 <- "result"
	}()

	select {
	case result := <-ch6:
		fmt.Println("Result:", result)

	case <-time.After(1 * time.Second):
		fmt.Println("Timeout")
	}

	/*
		The select waits for whichever happens first:

			┌── result arrives
			│
		select
			│
			└── 1 second timeout

		Here the work takes 2 seconds,
		but the timeout is 1 second.

		Therefore:

			Timeout

		The channel is buffered, so the worker can still
		finish and send the result even after the timeout.
	*/

	// ============================================================
	// 7. Multiple channels + timeout ⭐⭐⭐
	// ============================================================

	fmt.Println("\n7. Multiple channels + timeout")

	ch7 := make(chan string, 1)
	ch8 := make(chan string, 1)

	go func() {
		time.Sleep(2 * time.Second)
		ch7 <- "result from ch7"
	}()

	go func() {
		time.Sleep(3 * time.Second)
		ch8 <- "result from ch8"
	}()

	select {
	case value := <-ch7:
		fmt.Println("ch7:", value)

	case value := <-ch8:
		fmt.Println("ch8:", value)

	case <-time.After(5 * time.Second):
		fmt.Println("Timeout")
	}

	/*
		There are now three possible events:

			ch7 responds first
				↓
			ch7 case executes

			ch8 responds first
				↓
			ch8 case executes

			5 seconds pass first
				↓
			Timeout case executes
	*/

	// ============================================================
	// 8. Backend example ⭐⭐⭐
	// ============================================================

	fmt.Println("\n8. Backend example")

	/*
		Imagine an HTTP request needs data from two services:

		                 ┌── User Service
		Request ─────────┤
		                 └── Order Service


		We can start both operations concurrently:

			go fetchUser(userCh)
			go fetchOrders(orderCh)

		Then select whichever response becomes available first.
	*/

	userCh := make(chan string, 1)
	orderCh := make(chan string, 1)

	go fetchUser(userCh)
	go fetchOrders(orderCh)

	select {
	case user := <-userCh:
		fmt.Println("User received:", user)

	case orders := <-orderCh:
		fmt.Println("Orders received:", orders)

	case <-time.After(2 * time.Second):
		fmt.Println("Request timeout")
	}

	/*
		This pattern is useful for:

		- HTTP requests
		- database operations
		- external API calls
		- background workers
		- timeouts
		- cancellation
		- concurrent services
	*/

	// ============================================================
	// 9. select vs switch
	// ============================================================

	fmt.Println("\n9. select vs switch")

	/*
		switch:

			switch status {
			case 200:
				...
			case 404:
				...
			}

		Chooses based on a value or condition.


		select:

			select {
			case value := <-ch1:
				...

			case value := <-ch2:
				...
			}

		Chooses based on which channel operation
		is ready.
	*/

	// ============================================================
	// 10. The three important forms of select ⭐⭐⭐
	// ============================================================

	fmt.Println("\n10. The three important forms of select")

	/*
		1. Wait for one of several channels:

			select {
			case value := <-ch1:
				...

			case value := <-ch2:
				...
			}


		2. Non-blocking channel operation:

			select {
			case value := <-ch:
				...

			default:
				...
			}


		3. Channel operation with timeout:

			select {
			case value := <-ch:
				...

			case <-time.After(time.Second):
				...
			}
	*/

	// ============================================================
	// 11. Important mental model ⭐⭐⭐
	// ============================================================

	fmt.Println("\n11. Important mental model")

	/*
		select waits for channel operations.

			select {
			case value := <-ch1:
				// ch1 is ready

			case value := <-ch2:
				// ch2 is ready
			}


		No channel ready:

			ch1 ❌
			ch2 ❌

			  ↓

			BLOCK


		With default:

			ch1 ❌
			ch2 ❌

			  ↓

			default executes


		With timeout:

			ch1 ❌
			ch2 ❌

			  ↓
			wait
			  ↓
			timeout
			  ↓
			timeout case executes


		If multiple channels are ready:

			ch1 ✅
			ch2 ✅

			  ↓

			select chooses one ready case.


		The core idea:

			select = wait for one of multiple
			         channel operations.
	*/

	// ============================================================
	// Final summary
	// ============================================================

	fmt.Println("\nFinal summary")

	/*
		1. select works with channel operations.

		2. select can wait for multiple channels.

		3. Without default, select blocks when no case is ready.

		4. If multiple cases are ready, select chooses one.

		5. default makes select non-blocking.

		6. time.After can be used to implement timeouts.

		7. select is heavily used in concurrent backend systems.

		8. select does NOT receive from every ready channel.
		   It executes ONE selected case.

		9. A select case can be:

			Receive:

				case value := <-ch:

			Send:

				case ch <- value:

			Timeout:

				case <-time.After(time.Second):

			Default:

				default:

		10. Mental model:

				┌──────────────┐
				│   select     │
				├──────────────┤
				│ ch1 ready?   │
				│ ch2 ready?   │
				│ timeout?     │
				│ default?     │
				└──────────────┘
				       │
				       ▼
				  ONE case runs
	*/

	fmt.Println("\n39 — Select finished")
}
