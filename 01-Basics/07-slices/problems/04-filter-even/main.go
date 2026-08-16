package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	evenNumbers := make([]int, 0)

	for _, value := range numbers {
		if value%2 == 0 {
			evenNumbers = append(evenNumbers, value)
		}
	}

	fmt.Println(evenNumbers)
}
