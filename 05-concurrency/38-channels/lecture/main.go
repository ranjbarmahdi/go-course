package main

import "fmt"

// ============================================================
// Helper functions
// ============================================================

// Send-only channel.
// This function can only send values into the channel.
func send(ch chan<- int) {
	ch <- 10
}

// Receive-only channel.
// This function can only receive values from the channel.
func receive(ch <-chan int) {
	value := <-ch
	fmt.Println(value)
}

func main() {

	// ============================================================
	// 1. Why do we need channels?
	// ============================================================

	fmt.Println("1. Why do we need channels?")

	/*
		Goroutine A
			│
			│ send
			▼
		┌───────────┐
		│  Channel  │
		└───────────┘
			│
			│ receive
			▼
		Goroutine B

		A channel allows goroutines to communicate with each other.
	*/

	// Channels are one of the main ways goroutines communicate
	// and synchronize their work.

	// ============================================================
	// 2. What is a channel?
	// ============================================================

	fmt.Println("\n2. What is a channel?")

	// A channel is a Go value that allows goroutines
	// to send and receive values.

	ch := make(chan int)

	fmt.Println(ch)

	// ============================================================
	// 3. Sending a value
	// ============================================================

	fmt.Println("\n3. Sending a value")

	/*
		To send a value:

			ch <- value
	*/

	ch3 := make(chan int)

	go func() {
		ch3 <- 10
	}()

	// The receiver waits until the goroutine sends 10.
	fmt.Println(<-ch3)

	// ============================================================
	// 4. Receiving multiple values
	// ============================================================

	fmt.Println("\n4. Receiving multiple values")

	ch4 := make(chan int)

	go func() {
		ch4 <- 10
		ch4 <- 20
	}()

	// Receive the values in the same order they were sent.
	fmt.Println(<-ch4)
	fmt.Println(<-ch4)

	// ============================================================
	// 5. Channels connect goroutines ⭐⭐⭐
	// ============================================================

	fmt.Println("\n5. Channels connect goroutines")

	ch5 := make(chan int)

	go func() {
		ch5 <- 100
	}()

	value := <-ch5

	fmt.Println("Received:", value)

	/*
		The important idea:

			Goroutine
			    │
			    │ send
			    ▼
			Channel
			    │
			    │ receive
			    ▼
			Another goroutine
	*/

	// ============================================================
	// 6. Channels can synchronize goroutines ⭐⭐⭐
	// ============================================================

	fmt.Println("\n6. Channels can synchronize goroutines")

	ch6 := make(chan string)

	go func() {
		fmt.Println("Worker started")

		// Sending blocks until the receiver is ready.
		ch6 <- "Worker finished"
	}()

	// Receiving blocks until the worker sends a value.
	message := <-ch6

	fmt.Println(message)

	/*
		The channel is doing two things:

		1. Communication
		   Worker sends "Worker finished".

		2. Synchronization
		   main waits until the worker sends.
	*/

	// ============================================================
	// 7. Blocking behavior ⭐⭐⭐
	// ============================================================

	fmt.Println("\n7. Blocking behavior")

	/*
		With an unbuffered channel:

		Sender                         Receiver

		ch <- 10  ──────────────────►  <-ch
		    │                           │
		    │                           │
		    └── waits until receiver ───┘

		A send blocks until another goroutine receives.

		A receive blocks until another goroutine sends.
	*/

	// ============================================================
	// 8. Unbuffered channel
	// ============================================================

	fmt.Println("\n8. Unbuffered channel")

	/*
		make(chan int)

		creates an unbuffered channel.

		Capacity = 0

		There is no storage space inside the channel.
	*/

	ch8 := make(chan int)

	go func() {
		fmt.Println("Sending 10...")
		ch8 <- 10
		fmt.Println("Send completed")
	}()

	fmt.Println("Waiting for value...")

	value8 := <-ch8

	fmt.Println("Received:", value8)

	/*
		The send and receive synchronize directly.

		Sender
		   │
		   │ 10
		   ▼
		┌───────────┐
		│  Channel  │
		│ capacity 0│
		└───────────┘
		   │
		   ▼
		Receiver
	*/

	// ============================================================
	// 9. Buffered channels ⭐⭐⭐
	// ============================================================

	fmt.Println("\n9. Buffered channels")

	/*
		make(chan int, 3)

		creates a buffered channel with capacity 3.
	*/

	ch9 := make(chan int, 3)

	// These three sends do not block because
	// the channel has three available slots.
	ch9 <- 1
	ch9 <- 2
	ch9 <- 3

	fmt.Println("Channel length:", len(ch9))
	fmt.Println("Channel capacity:", cap(ch9))

	/*
		Buffer:

			┌──────────────────────┐
			│ 1 │ 2 │ 3 │
			└──────────────────────┘
			  capacity = 3

		The next send would block because the buffer is full.
	*/

	fmt.Println("Received:", <-ch9)

	// Now there is one free slot.
	ch9 <- 4

	fmt.Println("Received:", <-ch9)
	fmt.Println("Received:", <-ch9)
	fmt.Println("Received:", <-ch9)

	// ============================================================
	// 10. Buffered channel with a goroutine
	// ============================================================

	fmt.Println("\n10. Buffered channel with a goroutine")

	ch10 := make(chan int, 3)

	go func() {
		ch10 <- 10
		ch10 <- 20
		ch10 <- 30

		fmt.Println("Worker finished sending")
	}()

	// The worker can send all three values immediately
	// because the buffer has capacity 3.

	fmt.Println(<-ch10)
	fmt.Println(<-ch10)
	fmt.Println(<-ch10)

	// ============================================================
	// 11. Closing a channel
	// ============================================================

	fmt.Println("\n11. Closing a channel")

	/*
		close(ch)

		means:

			"No more values will be sent."

		Closing a channel does NOT delete its existing values.
	*/

	ch11 := make(chan int, 3)

	ch11 <- 10
	ch11 <- 20
	ch11 <- 30

	// The sender closes the channel after finishing sending.
	close(ch11)

	// Already-buffered values can still be received.
	fmt.Println(<-ch11)
	fmt.Println(<-ch11)
	fmt.Println(<-ch11)

	// ============================================================
	// 12. Checking whether a channel is closed
	// ============================================================

	fmt.Println("\n12. Checking whether a channel is closed")

	ch12 := make(chan int, 2)

	// The sender sends values.
	ch12 <- 10
	ch12 <- 20

	// The sender closes the channel because it is finished.
	close(ch12)

	value12, ok := <-ch12

	fmt.Println("Value:", value12)
	fmt.Println("Open:", ok)

	value12, ok = <-ch12

	fmt.Println("Value:", value12)
	fmt.Println("Open:", ok)

	/*
		Receive syntax:

			value, ok := <-ch

		ok == true

			A value was successfully received.

		ok == false

			The channel is closed and there are no
			more values to receive.
	*/

	value12, ok = <-ch12

	fmt.Println("Value:", value12)
	fmt.Println("Open:", ok)

	// ============================================================
	// 13. Range over a channel ⭐⭐⭐
	// ============================================================

	fmt.Println("\n13. Range over a channel")

	/*
		range can continuously receive values:

			for value := range ch {
				fmt.Println(value)
			}

		The loop stops when the channel is closed
		and all remaining values have been received.
	*/

	ch13 := make(chan int)

	go func() {
		ch13 <- 10
		ch13 <- 20
		ch13 <- 30

		// Sender tells the receiver:
		// "I am finished sending."
		close(ch13)
	}()

	for value := range ch13 {
		fmt.Println(value)
	}

	/*
		Sender                  Receiver

		  │                        │
		  │──── 10 ───────────────►│
		  │──── 20 ───────────────►│
		  │──── 30 ───────────────►│
		  │                        │
		  │──── close(ch) ────────►│
		  │                        │
		  │                   range stops
	*/

	// ============================================================
	// 14. Directional channels
	// ============================================================

	fmt.Println("\n14. Directional channels")

	/*
		Go allows us to restrict how a function can use a channel.

		Send-only:

			chan<- int

		Receive-only:

			<-chan int
	*/

	ch14 := make(chan int)

	go send(ch14)

	receive(ch14)

	/*
		send():

			func send(ch chan<- int)

			This function can ONLY send.

		receive():

			func receive(ch <-chan int)

			This function can ONLY receive.

		This makes APIs safer and easier to understand.
	*/

	// ============================================================
	// 15. Who should close a channel? ⭐⭐⭐
	// ============================================================

	fmt.Println("\n15. Who should close a channel?")

	/*
		General rule:

		The sender should close the channel.

		Why?

		The sender knows when there will be no more values.

		Example:

			go func() {
				ch <- 1
				ch <- 2
				ch <- 3

				close(ch)
			}()

			for value := range ch {
				fmt.Println(value)
			}

		The receiver normally does NOT close the channel.
	*/

	ch15 := make(chan int)

	go func() {
		for i := 1; i <= 3; i++ {
			ch15 <- i
		}

		close(ch15)
	}()

	for value := range ch15 {
		fmt.Println(value)
	}

	// ============================================================
	// 16. Important difference: buffered vs unbuffered
	// ============================================================

	fmt.Println("\n16. Buffered vs unbuffered")

	/*
		UNBUFFERED

			ch := make(chan int)

			Capacity = 0

			Send waits for a receiver.
			Receive waits for a sender.


		BUFFERED

			ch := make(chan int, 3)

			Capacity = 3

			Up to 3 values can wait inside the channel.

			Send blocks only when the buffer is full.
			Receive blocks only when the buffer is empty.
	*/

	// ============================================================
	// 17. Channel mental model ⭐⭐⭐
	// ============================================================

	fmt.Println("\n17. Channel mental model")

	/*
		Think of a channel as a communication pipe.

		Goroutine A
		     │
		     │ ch <- value
		     ▼
		┌───────────────┐
		│    CHANNEL    │
		└───────────────┘
		     │
		     │ value := <-ch
		     ▼
		Goroutine B


		Main operations:

			Send:

				ch <- value

			Receive:

				value := <-ch

			Close:

				close(ch)

			Check closed:

				value, ok := <-ch

			Range:

				for value := range ch
	*/

	// ============================================================
	// Important mental model
	// ============================================================

	fmt.Println("\nImportant mental model")

	/*
		1. Channels allow goroutines to communicate.

		2. Unbuffered channels provide direct synchronization.

		3. Buffered channels provide temporary storage.

		4. Send:

				ch <- value

		5. Receive:

				value := <-ch

		6. Close:

				close(ch)

		7. A receive from a closed channel can still get
		   remaining buffered values.

		8. After all values are consumed, receiving from
		   a closed channel gives the zero value and ok=false.

		9. range over a channel stops when the channel is closed
		   and all remaining values have been received.

		10. Normally, the sender closes the channel.

		11. A channel does NOT automatically make every operation
		    non-blocking.

		12. Blocking is an important part of channel synchronization.

		13. If a goroutine sends but nobody receives,
		    the sender can block forever.

		14. If a goroutine receives but nobody sends,
		    the receiver can block forever.

		15. Incorrect channel coordination can cause:

				fatal error: all goroutines are asleep - deadlock!
	*/

	fmt.Println("\n38 — Channels finished")
}
