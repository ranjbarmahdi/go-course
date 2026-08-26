/*
Problem:

Create a constraint:

type Number interface {
    int | int64 | float64
}

Then create:

func Add[T Number](a T, b T) T

The function should return a + b.

Requirements:

- Use the Number constraint.
- Support int.
- Support int64.
- Support float64.

Test:

Add(10, 20)

Add(int64(10), int64(20))

Add(1.5, 2.5)

Expected output:

30
30
4
*/

package main

import "fmt"

type Number interface {
	int | int64 | float64
}

func Add[T Number](a T, b T) T {
	return a + b
}

func main() {
	fmt.Println(Add(10, 20))

	fmt.Println(Add(int64(10), int64(20)))

	fmt.Println(Add(1.5, 2.5))
}
