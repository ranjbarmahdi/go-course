package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30}
	numbers = append(numbers, 40, 50, 60)
	fmt.Println(numbers)
}
