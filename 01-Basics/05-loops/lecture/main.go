package main

import "fmt"

func main() {

	// ============================================================
	// 1. Classic for Loop
	// ============================================================

	for i := 0; i <= 5; i++ {
		fmt.Println(i)
	}

	// ============================================================
	// 2. While-style Loop
	// ============================================================
	// Go does not have a separate `while` keyword.
	// A `for` loop can be used as a while loop.

	i := 0

	for i <= 10 {
		fmt.Println(i)
		i++
	}

	// ============================================================
	// 3. Infinite Loop
	// ============================================================
	// `for {}` creates an infinite loop.
	// Use `break` when you want to stop it.

	j := 0

	for {
		fmt.Println(j)
		j++

		if j == 5 {
			break
		}
	}

	// ============================================================
	// 4. continue
	// ============================================================
	// `continue` skips the current iteration
	// and moves to the next iteration.

	for i := 0; i <= 5; i++ {

		if i == 3 {
			continue
		}

		fmt.Println(i)
	}

	// ============================================================
	// 5. break
	// ============================================================
	// `break` immediately stops the loop.

	for i := 0; i <= 10; i++ {

		if i == 5 {
			break
		}

		fmt.Println(i)
	}

	// ============================================================
	// 6. Nested Loops
	// ============================================================
	// A loop can be placed inside another loop.

	for row := 1; row <= 3; row++ {

		for column := 1; column <= 3; column++ {
			fmt.Println("row:", row, "column:", column)
		}
	}
}
