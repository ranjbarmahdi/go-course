package main

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15, 2}
	target := 9

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if target == numbers[i]+numbers[j] {
				fmt.Printf("%d + %d = %d\n", numbers[i], numbers[j], target)
			}
		}
	}
}
