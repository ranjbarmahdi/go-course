/*
Problem:

Create a generic function:

func Contains[T comparable](values []T, target T) bool

The function should check whether target exists
inside the slice.

Requirements:

- Use the comparable constraint.
- Loop through the slice.
- Compare values using ==.
- Return true if found.
- Return false otherwise.

Test:

Contains([]int{1, 2, 3, 4}, 3)

Contains([]int{1, 2, 3, 4}, 10)

Contains([]string{"Ali", "Mahdi"}, "Mahdi")

Expected output:

true
false
true
*/

package main

import "fmt"

func Contains[T comparable](values []T, target T) bool {
	for _, v := range values {
		if v == target {
			return true
		}
	}

	return false
}

func main() {
	fmt.Println(Contains([]int{1, 2, 3, 4}, 3))

	fmt.Println(Contains([]int{1, 2, 3, 4}, 10))

	fmt.Println(Contains([]string{"Ali", "Mahdi"}, "Mahdi"))
}
