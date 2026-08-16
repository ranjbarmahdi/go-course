package main

import (
	"fmt"
	"math"
)

func main() {
	number := 1234567

	count := 0
	for temp := number; temp > 0; {
		temp /= 10
		count++
	}

	reversedNumber := 0
	for number > 0 {
		lastDigit := number % 10
		number = number / 10
		count--
		reversedNumber += int(float64(lastDigit) * math.Pow(10, float64(count)))
	}

	fmt.Println("Reversed Number is:", reversedNumber)
}
