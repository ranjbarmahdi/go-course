package main

import "fmt"

func main() {

	// ============================================================
	// 1. Variables with var
	// ============================================================

	var age int = 25

	fmt.Println(age)

	// ============================================================
	// 2. Short Variable Declaration :=
	// ============================================================

	// Go can infer the type from the assigned value.

	name := "Mahdi"
	height := 1.80
	developer := true

	fmt.Println(name)
	fmt.Println(height)
	fmt.Println(developer)

	// ============================================================
	// 3. Constants
	// ============================================================

	// A constant cannot be changed after it is declared.

	const pi = 3.14159

	fmt.Println(pi)

	// ============================================================
	// 4. Basic Data Types
	// ============================================================

	// string
	username := "Mahdi"

	// int
	score := 95

	// float64
	price := 19.99

	// bool
	active := true

	fmt.Println(username)
	fmt.Println(score)
	fmt.Println(price)
	fmt.Println(active)

	// ============================================================
	// 5. Type Inference
	// ============================================================

	// Go automatically determines the type
	// based on the assigned value.

	message := "Hello"
	count := 10
	ratio := 1.5
	passed := true

	fmt.Println(message)
	fmt.Println(count)
	fmt.Println(ratio)
	fmt.Println(passed)

	// ============================================================
	// 6. Type Conversion
	// ============================================================

	// Go does not automatically convert between numeric types.

	var number int = 10
	var decimal float64 = float64(number)

	fmt.Println(number)
	fmt.Println(decimal)

	// ============================================================
	// 7. Arithmetic
	// ============================================================

	a := 10
	b := 3

	fmt.Println("Addition:", a+b)
	fmt.Println("Subtraction:", a-b)
	fmt.Println("Multiplication:", a*b)
	fmt.Println("Division:", a/b)
	fmt.Println("Remainder:", a%b)
}
