package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}
	numbers = append(numbers[:2], numbers[3:]...)
	fmt.Println(numbers)
}
