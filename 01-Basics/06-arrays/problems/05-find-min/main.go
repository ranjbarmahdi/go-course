package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := [7]int{12, 45, 7, 89, 23, 56, 34}

	min := math.MaxInt

	for _, value := range numbers {
		if value < min {
			min = value
		}
	}

	fmt.Println("Min =", min)
}
