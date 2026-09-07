package main

import (
	"errors"
	"fmt"
)

type User struct {
	ID   int
	Name string
}

// ============================================================
// 1. fmt.Print()
// ============================================================
// Prints values without automatically adding a newline.
//
// You must add "\n" yourself when needed.

func printExample() {

	fmt.Print("Hello\n")
	fmt.Print("Go\n")
}

// ============================================================
// 2. fmt.Println()
// ============================================================
// Prints values and automatically adds a newline.
//
// Multiple values are separated by a space.

func printlnExample() {

	fmt.Println("Hello")
	fmt.Println("Go")
}

// ============================================================
// 3. fmt.Printf()
// ============================================================
// Prints values using a format string.
//
// Common format verbs:
//
// %s  -> string
// %d  -> integer
// %f  -> floating-point number
// %t  -> boolean
// %T  -> type
// %v  -> default value format
// %+v -> value including struct field names

func printfExample() {

	name := "Mahdi"
	age := 25

	fmt.Printf(
		"Name: %s Age: %d\n",
		name,
		age,
	)

	price := 19.99

	// %.2f -> float with 2 decimal places
	fmt.Printf("%.2f\n", price)

	active := true

	fmt.Printf("%t\n", active)

	number := 100

	fmt.Printf("%T\n", number)

	name1 := "Mahdi"

	fmt.Printf("%v\n", name1)

	user := User{
		ID:   1,
		Name: "Mahdi",
	}

	// %v -> default struct representation
	fmt.Printf("%v\n", user)

	// %+v -> struct fields + values
	fmt.Printf("%+v\n", user)
}

// ============================================================
// 4. fmt.Sprintf()
// ============================================================
// Sprintf formats a string and RETURNS the string.
//
// Unlike Printf, it does not print anything by itself.
//
// Useful when you need to build a string.

func sprintfExample() {

	name := "Mahdi"
	age := 27

	user := User{
		ID:   1,
		Name: "Mahdi",
	}

	message := fmt.Sprintf(
		"name: %s\nage: %d\nuser: %+v",
		name,
		age,
		user,
	)

	fmt.Println(message)
}

// ============================================================
// 5. Formatting Structs
// ============================================================

func structFormatting() {

	user := User{
		ID:   1,
		Name: "Mahdi",
	}

	// Default representation
	fmt.Printf("%v\n", user)

	// Field names + values
	fmt.Printf("%+v\n", user)
}

// ============================================================
// 6. Formatting Values
// ============================================================

func valueFormatting() {

	name := "Mahdi"
	age := 27
	price := 19.99
	active := true

	fmt.Printf("Name: %s\n", name)
	fmt.Printf("Age: %d\n", age)
	fmt.Printf("Price: %.2f\n", price)
	fmt.Printf("Active: %t\n", active)

	// Print the type of a value
	fmt.Printf("Type: %T\n", age)

	// Default formatting
	fmt.Printf("Value: %v\n", name)
}

// ============================================================
// 7. Building Error Messages with Sprintf
// ============================================================
// Sprintf is useful when creating dynamic strings.
//
// It is commonly used when constructing error messages.

func UserNotFound(id int) error {

	message := fmt.Sprintf(
		"user %d not found",
		id,
	)

	return errors.New(message)
}

func main() {

	// ============================================================
	// 1. fmt.Print()
	// ============================================================

	printExample()

	// ============================================================
	// 2. fmt.Println()
	// ============================================================

	printlnExample()

	// ============================================================
	// 3. fmt.Printf()
	// ============================================================

	printfExample()

	// ============================================================
	// 4. fmt.Sprintf()
	// ============================================================

	sprintfExample()

	// ============================================================
	// 5. Formatting Structs
	// ============================================================

	structFormatting()

	// ============================================================
	// 6. Formatting Values
	// ============================================================

	valueFormatting()

	// ============================================================
	// 7. Sprintf + errors
	// ============================================================

	fmt.Println(UserNotFound(1))
}
