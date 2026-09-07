/*
============================================================
Problem 5 — Pipeline ⭐⭐⭐
============================================================

Build a 3-stage pipeline:

    Generate → Square → Print

Requirements:

Stage 1 — Generate

Create:

    func generate() <-chan int

Generate the numbers:

    1 through 10

Send them to a channel.

Close the channel when finished.

Stage 2 — Square

Create:

    func square(in <-chan int) <-chan int

For every received value:

    result := value * value

Send the result to another channel.

Close the output channel when the input channel is
finished.

Stage 3 — Main

Range over the output of square() and print every value.

Expected values:

    1
    4
    9
    16
    25
    36
    49
    64
    81
    100

The important structure is:

    generate()
        ↓
      channel
        ↓
     square()
        ↓
      channel
        ↓
      main

Each stage should communicate through channels.

Do NOT use shared global variables.
*/

package main

import "fmt"

func main() {
	fmt.Println("pipeline")
}
