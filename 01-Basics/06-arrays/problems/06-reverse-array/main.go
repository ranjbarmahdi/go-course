package main

import "fmt"

func main() {

	numbers := [5]int{10, 20, 30, 40, 50}
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - 1 - i
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	fmt.Println("Reversed:", numbers)
}
