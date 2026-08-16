package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := [7]int{12, 45, 7, 89, 23, 56, 34}

	max := math.MinInt

	for _, value := range numbers {
		if value > max {
			max = value
		}
	}

	fmt.Println("Max =", max)
}
