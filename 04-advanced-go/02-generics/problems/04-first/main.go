/*
Problem:

Create a generic function:

func First[T any](values []T) T

The function should return the first element
of the slice.

Requirements:

- Accept []T.
- Return T.
- Do not use any type assertion.
- Test with []int.
- Test with []string.

Example:

Input:

[]int{10, 20, 30}

Output:

10

Input:

[]string{"Ali", "Mahdi", "Sara"}

Output:

Ali
*/

package main

import "fmt"

func First[T any](values []T) T {
	return values[0]
}

func main() {
	names := []string{"Ali", "Mahdi", "Sara"}
	fmt.Println(First(names))

	numbers := []int{10, 20, 30}
	fmt.Println(First(numbers))

}
