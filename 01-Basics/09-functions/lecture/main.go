package main

import "fmt"

// ============================================================
// 1. Basic Function
// ============================================================

func sayHello() {
	fmt.Println("Hello")
}

// ============================================================
// 2. Function with Parameters
// ============================================================

func greet(name string) {
	fmt.Println("Hello", name)
}

// ============================================================
// 3. Multiple Parameters
// ============================================================
// Parameters with the same type can be grouped:
//
// func add(a, b int)

func add(a, b int) {
	fmt.Println(a + b)
}

// ============================================================
// 4. Return Values
// ============================================================
// A function can return a value.
//
// The return type is written after the parameters.

func add2(a, b int) int {
	return a + b
}

// ============================================================
// 5. Multiple Return Values
// ============================================================
// A Go function can return multiple values.
//
// This is very common in Go, especially for:
// - result + error
// - result + success
// - value + exists

func divide(a, b float64) (float64, bool) {

	if b == 0 {
		return 0, false
	}

	return a / b, true
}

// ============================================================
// 6. Named Return Values
// ============================================================
// Return values can have names.
//
// The named variable is automatically created by Go.

func calculate(a, b int) (sum int) {

	sum = a + b

	return
}

// ============================================================
// 7. Variadic Functions
// ============================================================
// `...int` means the function can receive zero or more ints.
//
// Inside the function, `numbers` is a []int slice.

func sum(numbers ...int) int {

	total := 0

	for _, number := range numbers {
		total += number
	}

	return total
}

// ============================================================
// 8. Functions as Values
// ============================================================
// Functions can be assigned to variables.
//
// `myFunc` has the same function type as `add2`.
//
// Function type:
// func(int, int) int

func add3(a, b int) int {
	return a + b
}

// ============================================================
// 9. Anonymous Functions
// ============================================================
// A function can be created without a name.
//
// This is called an anonymous function.

// ============================================================
// 10. defer
// ============================================================
// defer delays a function call until the surrounding function
// is about to return.
//
// Commonly used for:
// - closing files
// - unlocking mutexes
// - cleanup operations

// ============================================================
// 11. Returning Errors
// ============================================================
// Go commonly returns an error as one of the return values.
//
// nil means there is no error.

func divide2(a, b float64) (float64, error) {

	if b == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}

	return a / b, nil
}

func main() {

	// ============================================================
	// 1. Basic Function
	// ============================================================

	sayHello()

	// ============================================================
	// 2. Function with Parameters
	// ============================================================

	greet("Mahdi")

	// ============================================================
	// 3. Multiple Parameters
	// ============================================================

	add(1, 2)

	// ============================================================
	// 4. Return Values
	// ============================================================

	res := add2(1, 2)

	fmt.Println(res)

	// ============================================================
	// 5. Multiple Return Values
	// ============================================================

	result, ok := divide(10, 2)

	fmt.Println(result)
	fmt.Println(ok)

	// ============================================================
	// 6. Named Return Values
	// ============================================================

	res2 := calculate(1, 1)

	fmt.Println(res2)

	// ============================================================
	// 7. Variadic Functions
	// ============================================================

	fmt.Println(sum(1, 2, 3, 4))

	fmt.Println(sum(10, 20))

	fmt.Println(sum())

	// ============================================================
	// 8. Functions as Values
	// ============================================================

	myFunc := add3

	fmt.Println(myFunc(1, 2))

	// ============================================================
	// 9. Anonymous Functions
	// ============================================================

	addAnonymous := func(a, b int) int {
		return a + b
	}

	fmt.Println(addAnonymous(1, 2))

	// ============================================================
	// 10. defer
	// ============================================================

	defer fmt.Println("Finished")

	fmt.Println("Running")

	// ============================================================
	// 11. Errors
	// ============================================================

	result, err := divide2(10, 0)

	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(result)
	}
}
