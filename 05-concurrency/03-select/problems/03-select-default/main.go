/*
============================================================
Problem 3 — default and non-blocking select
============================================================

Create an unbuffered channel:

    ch := make(chan int)

Do NOT send anything to the channel.

Use select with:

    case value := <-ch:
        // receive and print the value

    default:
        // print "No value available"

Requirements:
- Use select.
- Use default.
- Do NOT use goroutines.
- Do NOT use time.Sleep().
- Do NOT use time.After().
- Do NOT use WaitGroup.
- Do NOT use close().

Expected output:

    No value available

Goal:
Understand that adding default makes select NON-BLOCKING.

Without default:

    select {
    case value := <-ch:
        ...
    }

    ↓

    If ch is not ready → BLOCK

With default:

    select {
    case value := <-ch:
        ...
    default:
        ...
    }

    ↓

    If ch is not ready → immediately execute default.

Important:
default does NOT wait for a value to arrive.
It executes immediately if no other case is ready.
*/

package main

import "fmt"

func main() {
	ch := make(chan int)

	select {
	case value := <-ch:
		fmt.Println(value)
	default:
		fmt.Println("No value available")
	}
}
