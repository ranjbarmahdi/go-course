package main

import "fmt"

func main() {

	// ============================================================
	// 1. What is an Array?
	// ============================================================
	// An array stores multiple values of the same type.
	//
	// The length is part of the array's type.
	// [5]int and [3]int are different types.

	numbers := [5]int{10, 20, 30, 40, 50}

	fmt.Println(numbers)

	// ============================================================
	// 2. Changing an Element
	// ============================================================
	// Access an element using its index.
	// Array indexes start from 0.

	scores := [5]int{5, 3, 1, 6, 8}

	scores[0] = 10000

	fmt.Println(scores)

	// ============================================================
	// 3. Array Length
	// ============================================================
	// len() returns the number of elements.

	fmt.Println(len(scores))

	// ============================================================
	// 4. Let Go Count the Elements
	// ============================================================
	// [...] tells Go to calculate the array length
	// from the number of elements provided.

	alphabets := [...]string{"A", "B", "C"}

	fmt.Println(alphabets)
	fmt.Println(len(alphabets))

	// ============================================================
	// 5. Arrays + for Loop
	// ============================================================
	// Use the index to access each element.

	arr := [5]int{10, 20, 30, 40, 50}

	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	// ============================================================
	// 6. range
	// ============================================================
	// range gives both the index and the value.

	arr2 := [5]int{10, 20, 30, 40, 50}

	for index, value := range arr2 {
		fmt.Println(index, value)
	}

	// ============================================================
	// 7. range Without the Index
	// ============================================================
	// Use `_` when you don't need the index.

	arr3 := [5]int{10, 20, 30, 40, 50}

	for _, value := range arr3 {
		fmt.Println(value)
	}

	// ============================================================
	// 8. Arrays Are Values
	// ============================================================
	// Assigning an array to another array creates a copy.
	//
	// Changing the copy does NOT change the original array.

	a := [3]int{1, 2, 3}

	b := a

	b[0] = 100

	fmt.Println("a:", a)
	fmt.Println("b:", b)

	// ============================================================
	// 9. Multi-dimensional Arrays
	// ============================================================
	// An array can contain other arrays.
	//
	// [2][3]int means:
	// 2 rows
	// 3 columns

	matrix := [2][3]int{
		{1, 1, 1},
		{2, 2, 2},
	}

	for i := 0; i < len(matrix); i++ {

		for j := 0; j < len(matrix[i]); j++ {
			fmt.Printf("%d\t", matrix[i][j])
		}

		fmt.Println()
	}
}
