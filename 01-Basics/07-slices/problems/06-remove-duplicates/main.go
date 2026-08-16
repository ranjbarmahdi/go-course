package main

import "fmt"

func main() {
	numbers := []int{1, 2, 2, 2, 3, 4, 4, 5, 5, 5}

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				numbers = append(numbers[:j], numbers[j+1:]...)
				j--
			}
		}
	}

	fmt.Println(numbers)
}
