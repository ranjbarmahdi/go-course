package main

import "fmt"

func main() {
	numbers := [...]int{10, 20, 30, 40, 50}
	fmt.Println("First:", numbers[0])
	fmt.Println("Last:", numbers[len(numbers)-1])
	fmt.Println("Length:", len(numbers))

	numbers[0] = 100
	fmt.Println(numbers)
}
