/*
Problem:

Create a generic function:

func Identity[T any](value T) T

The function should return exactly the value
that it receives.

Requirements:

- Use a type parameter T.
- Accept any type.
- Return the same type T.
- In main(), test it with:
    - int
    - string

Expected output:

100
Mahdi
*/

package main

import "fmt"

func Identity[T any](value T) T {
	return value
}

func main() {
	fmt.Println(Identity(100))
	fmt.Println(Identity("Mahdi"))
}
