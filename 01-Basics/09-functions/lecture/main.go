package main

import (
	"fmt"
)

func sayHello() {
	fmt.Println("Hello")
}

func greet(name string) {
	fmt.Println("Hello", name)
}

func add(a, b int) {
	fmt.Println(a + b)
}

func add2(a, b int) int {
	return a + b
}

func divide(a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}

	return a / b, true
}

func calculate(a, b int) (sum int) {
	sum = a + b
	return
}

func sum(numbers ...int) int {
	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

func divide2(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func main() {
	fmt.Println("1. Basic Function")
	sayHello()

	fmt.Println("2. Function with Parameters")
	greet("Mahdi")

	fmt.Println("3. Multiple Parameters")
	add(1, 2)

	fmt.Println("4. Return Values")
	res := add2(1, 2)
	fmt.Println(res)

	fmt.Println("5. Multiple Return Values")
	result, ok := divide(10, 2)
	fmt.Println(result)
	fmt.Println(ok)

	fmt.Println("6. Named Return Values")
	res2 := calculate(1, 1)
	fmt.Println(res2)

	fmt.Println("7. Variadic Functions")
	fmt.Println(sum(1, 2, 3, 4))

	fmt.Println("8. Functions as Values")
	myFunc := add2
	fmt.Println(myFunc(1, 2))

	fmt.Println("9. Anonymous Functions")
	sum := func(a, b int) int {
		return a + b
	}
	fmt.Println(sum(1, 2))

	fmt.Println("9. Defer")
	defer fmt.Println("Finished")
	fmt.Println("Running")

	fmt.Println("10. Errors")
	result, err := divide2(10, 0)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
