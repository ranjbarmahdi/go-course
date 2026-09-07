/*
Problem:

Create a generic function:

func PrintPair[T any, U any](a T, b U)

The function should print both values.

Requirements:

- Use two type parameters.
- T can be any type.
- U can be any type.
- Print a.
- Print b.

In main(), call:

PrintPair(27, "Mahdi")

Expected output:

27
Mahdi

Then call:

PrintPair("Age", 27)

Expected output:

Age
27
*/

package main

import "fmt"

func PrintPair[T any, U any](a T, b U) {
	fmt.Println(a)
	fmt.Println(b)
}

func main() {
	PrintPair(27, "Mahdi")
	PrintPair("Age", 27)
}
