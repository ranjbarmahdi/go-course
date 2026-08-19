package main

import "fmt"

// ============================================================
// 1. Memory Address
// ============================================================
// Every variable is stored somewhere in memory.
//
// `&` returns the memory address of a variable.

func main() {

	x := 10

	fmt.Println(x)
	fmt.Println(&x)

	// ============================================================
	// 2. Pointer Variable
	// ============================================================
	// A pointer stores the memory address of another variable.
	//
	// `*int` means:
	// "pointer to an int"

	x1 := 10

	var p1 *int
	p1 = &x1

	fmt.Println(x1)
	fmt.Println(p1)
	fmt.Println(*p1)

	// `p1`  -> memory address of x1
	// `*p1` -> value stored at that address

	// ============================================================
	// 3. Changing Through a Pointer
	// ============================================================
	// `*p` accesses the value stored at the address.
	//
	// Changing `*p` changes the original variable.

	x2 := 10

	var p *int = &x2

	*p = 50

	fmt.Println(x2)
	fmt.Println(*p)
}
