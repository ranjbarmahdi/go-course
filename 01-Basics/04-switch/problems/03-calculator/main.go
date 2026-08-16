package main

import "fmt"

func main() {
	a := 10
	b := 2
	operator := "%"

	switch operator {
	case "+":
		sum := a + b
		fmt.Println("a + b =", a, "+", b, "=", sum)
	case "-":
		difference := a - b
		fmt.Println("a - b =", a, "-", b, "=", difference)

	case "*":
		multiplication := a * b
		fmt.Println("a * b =", a, "*", b, "=", multiplication)
	case "/":
		division := float64(a) / float64(b)
		fmt.Println("a / b =", a, "/", b, "=", division)
	case "%":
		remainder := a % b
		fmt.Println("a % b =", a, "%", b, "=", remainder)
	default:
		fmt.Println("Invalid operator.")
	}
}
