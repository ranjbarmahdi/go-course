package main

import "fmt"

// ============================================================
// 1. Struct
// ============================================================
// A struct groups related data into a custom type.

type Rectangle struct {
	Width  float64
	Height float64
}

// ============================================================
// 2. Method
// ============================================================
// A method is a function associated with a type.
//
// `(r Rectangle)` is called the receiver.
//
// The receiver allows the method to access the Rectangle's fields.

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// ============================================================
// 3. Another Method
// ============================================================

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func main() {

	// ============================================================
	// 4. Creating a Struct
	// ============================================================

	rectangle := Rectangle{
		Width:  10,
		Height: 5,
	}

	// ============================================================
	// 5. Calling Methods
	// ============================================================

	fmt.Println("Area:", rectangle.Area())
	fmt.Println("Perimeter:", rectangle.Perimeter())
}
