package main

import "fmt"

func main() {
	a := 20
	b := 7

	sum := a + b
	difference := a - b
	multiplication := a * b
	division := a / b
	remainder := a % b

	fmt.Println("Sum", sum)
	fmt.Println("Difference", difference)
	fmt.Println("Multiplication", multiplication)
	fmt.Println("Division", division)
	fmt.Println("Remainder", remainder)
}
