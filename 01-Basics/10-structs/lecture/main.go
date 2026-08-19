package main

import "fmt"

// ============================================================
// 1. Struct
// ============================================================
// A struct groups related data into a custom type.
//
// Each field has its own name and type.

type User struct {
	Name  string
	Age   int
	Email string
}

// ============================================================
// 2. Creating a Struct
// ============================================================
// Create a value of the User struct by providing field values.

func main() {

	user := User{
		Name:  "Mahdi",
		Age:   27,
		Email: "mahdi@gmail.com",
	}

	// ============================================================
	// 3. Accessing Struct Fields
	// ============================================================
	// Use `.` to access a struct field.

	fmt.Println(user)

	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Email)
}
