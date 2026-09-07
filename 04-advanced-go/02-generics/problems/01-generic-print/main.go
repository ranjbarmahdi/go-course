/*
Problem:

Create a generic function:

func Print[T any](value T)

The function should print the provided value.

Requirements:

- Use a type parameter T.
- T must accept any type.
- Use fmt.Println().
- In main(), call Print() with:
    - an int
    - a string
    - a bool
    - a float64

Expected output:

100
Mahdi
true
3.14
*/

package main

import "fmt"

func Print[T any](value T) {
	fmt.Println(value)
}

func main() {
	Print(100)
	Print("Mahdi")
	Print(true)
	Print(3.14)
}
