package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3, 4, 5}

	for range 3 {
		temp := numbers[len(numbers)-1]
		for i := len(numbers) - 1; i > 0; i-- {
			numbers[i] = numbers[i-1]
		}
		numbers[0] = temp
	}

	fmt.Println(numbers)
}
