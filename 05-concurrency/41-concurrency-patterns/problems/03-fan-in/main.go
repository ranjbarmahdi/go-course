/*
============================================================
Problem 3 — Fan-In ⭐⭐⭐
============================================================

Build a simple Fan-In pattern.

The idea:

    Worker 1 ──┐
    Worker 2 ──┼──> ONE results channel
    Worker 3 ──┘

Multiple goroutines produce results into ONE channel.

Requirements:

1. Create 3 separate input channels:

    ch1 := make(chan int)
    ch2 := make(chan int)
    ch3 := make(chan int)

2. Create ONE output channel:

    results := make(chan int)

3. Create 3 producer goroutines.

Producer 1 must send:

    1, 2, 3

Producer 2 must send:

    4, 5, 6

Producer 3 must send:

    7, 8, 9

4. Each producer must close its own channel when
   it has finished sending.

5. Create a fan-in function:

    func fanIn(
        ch1 <-chan int,
        ch2 <-chan int,
        ch3 <-chan int,
        results chan<- int,
    )

The fan-in function must receive values from all 3 input
channels and send them into the ONE results channel.

6. After all three input channels are completely consumed,
   close the results channel.

7. In main, range over results and print every value.

Expected values:

    1 2 3 4 5 6 7 8 9

IMPORTANT:

The exact order is NOT guaranteed.

For example:

    1
    4
    7
    2
    5
    8
    3
    6
    9

is also valid.

The important part is that ALL 9 values arrive through
the SAME results channel.

Fan-In means:

    MANY sources
         ↓
      ONE channel
         ↓
       main

Use:

- goroutines
- channels
- select
- sync.WaitGroup
- close()
- range over channel

Do NOT use time.Sleep() for synchronization.

Think about who is responsible for closing `results`.
The input producers must NOT close the shared `results`
channel.
*/

package main

import "fmt"

func main() {
	fmt.Println("Fain in")
}
