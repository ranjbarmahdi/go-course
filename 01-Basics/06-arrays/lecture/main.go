package main

import "fmt"

func main() {
	fmt.Println("1. What is an array?")
	// An array stores multiple values of the same type.
	numbers := [5]int{10, 20, 30, 40, 50} // -The 5 is part of the type. [5]int
	fmt.Println(numbers)

	fmt.Println("2. Changing an element")
	scores := [5]int{5, 3, 1, 6, 8}
	scores[0] = 10000

	fmt.Println("3. Array Length")
	fmt.Println(len(scores))

	fmt.Println("4. Let Go count the elements")
	alphabets := [...]string{"A", "B", "C"}
	fmt.Println(alphabets)

	fmt.Println("5. Arrays + loops")
	arr := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}

	fmt.Println("6. range")
	arr2 := [5]int{10, 20, 30, 40, 50}
	for index, value := range arr2 {
		fmt.Println(index, value)
	}

	fmt.Println("7. You don't need the index")
	arr3 := [5]int{10, 20, 30, 40, 50}
	for _, value := range arr3 {
		fmt.Println(value)
	}

	fmt.Println("8. Arrays are values")
	a := [3]int{1, 2, 3}
	b := a // copy (no reference)
	b[0] = 100
	fmt.Println(a)
	fmt.Println(b)

	fmt.Println("9. Multi-dimensional arrays")
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
