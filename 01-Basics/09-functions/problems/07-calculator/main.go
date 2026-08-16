package main

import "fmt"

func calculate(a, b float64, operator string) (float64, bool) {
	switch operator {
	case "+":
		return a + b, true
	case "-":
		return a - b, true
	case "*":
		return a * b, true
	case "/":
		if b == 0 {
			return 0, false
		}
		return a / b, true
	default:
		return 0, false
	}
}

func main() {
	result, ok := calculate(10, 2, "/")
	fmt.Println(result, ok)
}
