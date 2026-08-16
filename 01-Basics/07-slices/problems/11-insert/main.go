package main

import "fmt"

func main() {
	numbers := []int{10, 20, 40, 50}
	numbers = append(numbers, 0)
	copy(numbers[3:], numbers[2:])
	numbers[2] = 30
	fmt.Println(numbers)
}
