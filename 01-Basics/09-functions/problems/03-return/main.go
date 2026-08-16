package main

import "fmt"

func multiply(a, b int) int {
	return a * b
}

func main() {
	result := multiply(2, 5)
	fmt.Println(result)
}
