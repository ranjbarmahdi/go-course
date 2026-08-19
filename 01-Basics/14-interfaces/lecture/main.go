package main

import "fmt"

// ============================================================
// 1. Interface
// ============================================================
// An interface defines a set of methods.
//
// Any type that implements all of the interface's methods
// satisfies the interface automatically.

type Speaker interface {
	Speak()
}

// ============================================================
// 2. Dog
// ============================================================

type Dog struct {
	Name string
}

func (d Dog) Speak() {
	fmt.Println("Woof")
}

// ============================================================
// 3. Cat
// ============================================================

type Cat struct {
	Name string
}

func (c Cat) Speak() {
	fmt.Println("Meow")
}

// ============================================================
// 4. Using an Interface
// ============================================================
// The function accepts any value that satisfies Speaker.
//
// Both Dog and Cat satisfy Speaker because both have:
//     Speak()

func makeSound(s Speaker) {
	s.Speak()
}

// ============================================================
// 5. Empty Interface
// ============================================================
// `interface{}` can hold a value of any type.
//
// `any` is an alias for `interface{}`.
//
// These are equivalent:
//
// interface{}
// any

func print(value interface{}) {
	fmt.Println(value)
}

// ============================================================
// 6. Type Assertion
// ============================================================
// A type assertion checks whether an interface value contains
// a specific type.
//
// Syntax:
//
// value.(Type)
//
// The safe form returns:
// - the value
// - a boolean indicating whether the assertion succeeded

// ============================================================
// 7. Type Switch
// ============================================================
// A type switch checks the concrete type stored inside
// an interface value.
//
// Syntax:
//
// switch v := value.(type) {
// case int:
//     ...
// }

func check(value any) {

	switch v := value.(type) {

	case int:
		fmt.Println("Integer:", v)

	case string:
		fmt.Println("String:", v)

	case bool:
		fmt.Println("Boolean:", v)

	case Dog:
		fmt.Println("Dog:", v)

	default:
		fmt.Println("Unknown")
	}
}

func main() {

	// ============================================================
	// 8. Basic Interface Usage
	// ============================================================

	dog := Dog{
		Name: "Buddy",
	}

	cat := Cat{
		Name: "Kitty",
	}

	makeSound(dog)
	makeSound(cat)

	// ============================================================
	// 9. Empty Interface
	// ============================================================
	// An `any` value can contain values of different types.

	print("Hello World")
	print(2)
	print(true)

	// ============================================================
	// 10. Type Assertion
	// ============================================================

	var value any

	value = 100

	number, ok := value.(int)

	if ok {
		fmt.Println(number)
	}

	value = "hello"

	number2, ok2 := value.(int)

	if ok2 {
		fmt.Println(number2)
	}

	// ok2 is false because value contains a string,
	// not an int.

	// ============================================================
	// 11. Type Switch
	// ============================================================

	check("123")
	check(1)
	check(true)

	dog2 := Dog{
		Name: "Black",
	}

	check(dog2)
}
