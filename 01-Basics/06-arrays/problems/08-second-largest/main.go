package main

import (
	"fmt"
	"math"
)

func main() {
	numbers := [8]int{10, 25, 7, 42, 18, 99, 31, 50}

	max1 := math.MinInt
	max2 := math.MinInt

	for _, value := range numbers {
		if value > max1 {
			max2 = max1
			max1 = value
		} else if value > max2 {
			max2 = value
		}
	}

	fmt.Println("Second Max =", max2)
}
