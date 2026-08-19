package main

import "fmt"

func main() {

	// ============================================================
	// 1. if / else if / else
	// ============================================================

	age := 27

	if age < 13 {
		fmt.Println("Child")
	} else if age < 18 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Adult")
	}

	// ============================================================
	// 2. Boolean Operators
	// ============================================================

	// AND (&&)
	// True only when both conditions are true.
	adultWorkingAge := age >= 18 && age <= 65

	fmt.Println("Adult working age:", adultWorkingAge)

	// OR (||)
	// True when at least one condition is true.
	childOrSenior := age < 18 || age > 65

	fmt.Println("Child or senior:", childOrSenior)

	// NOT (!)
	// Reverses a boolean value.

	isDeveloper := true

	fmt.Println("Is not developer:", !isDeveloper)

	// ============================================================
	// 3. Comparison Operators
	// ============================================================

	a := 10
	b := 20

	fmt.Println("a == b:", a == b)
	fmt.Println("a != b:", a != b)
	fmt.Println("a > b:", a > b)
	fmt.Println("a < b:", a < b)
	fmt.Println("a >= b:", a >= b)
	fmt.Println("a <= b:", a <= b)

	// ============================================================
	// 4. Boolean Values
	// ============================================================

	isLoggedIn := true
	isAdmin := false

	if isLoggedIn && isAdmin {
		fmt.Println("Admin is logged in")
	}

	if isLoggedIn || isAdmin {
		fmt.Println("User has access")
	}

	if !isAdmin {
		fmt.Println("User is not an admin")
	}
}
