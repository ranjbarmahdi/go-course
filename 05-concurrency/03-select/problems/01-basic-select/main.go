/*
============================================================
Problem 1 — Basic select
============================================================

Create two channels:

    ch1 := make(chan string)
    ch2 := make(chan string)

Start two goroutines:

    Goroutine 1 → send "Hello from ch1" to ch1
    Goroutine 2 → send "Hello from ch2" to ch2

In main(), use select to receive from ONE of the channels.

Print the received value.

Requirements:
- Use goroutines.
- Use select.
- Use two channels.
- Do NOT use time.Sleep().
- Do NOT use WaitGroup.
- Do NOT use close().
- Do NOT use default.
- Do not worry about which message is printed.
  Either channel may be selected.

Expected output is one of:

    Hello from ch1

or:

    Hello from ch2

Important:
The output is intentionally nondeterministic because both
channel operations may become ready.
*/

package main

import "fmt"

func main() {
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		ch1 <- "Hello from ch1"
	}()

	go func() {
		ch2 <- "Hello from ch2"
	}()

	select {
	case value := <-ch1:
		fmt.Println(value)
	case value := <-ch2:
		fmt.Println(value)
	}
}
