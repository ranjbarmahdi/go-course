package main

import (
	"fmt"
	"math"
)

func findMax(numbers []int) (int, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("numbers is empty")
	}

	max := math.MinInt
	for _, num := range numbers {
		if num > max {
			max = num
		}
	}

	return max, nil
}

func main() {
	numbers := []int{10, 50, 20, 90, 30}
	fmt.Println(findMax(numbers))
}
