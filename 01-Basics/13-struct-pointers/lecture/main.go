package main

import "fmt"

// ============================================================
// 1. Struct
// ============================================================

type User struct {
	Name string
	Age  int
}

// ============================================================
// 2. Pointer Receiver
// ============================================================
// A pointer receiver allows a method to modify the original
// struct.
//
// `*User` means the receiver is a pointer to User.

func (u *User) Birthday1() {
	u.Age++
}

// ============================================================
// 3. Value Receiver
// ============================================================
// A value receiver receives a copy of the struct.
//
// Changes made inside the method affect only the copy.

func (u User) Birthday2() {
	u.Age++
}

func main() {

	// ============================================================
	// 4. Pointer to Struct
	// ============================================================
	// `&user` gets the memory address of the struct.

	user := User{
		Name: "Mahdi",
		Age:  27,
	}

	userPtr := &user

	fmt.Println(userPtr.Name)
	fmt.Println((*userPtr).Age)

	// Go automatically dereferences struct pointers when
	// accessing fields:
	//
	// userPtr.Name
	// is equivalent to:
	// (*userPtr).Name

	// ============================================================
	// 5. Creating a Struct with a Pointer
	// ============================================================
	// `&User{...}` creates a User value and returns its address.

	user2 := &User{
		Name: "Mahdi",
		Age:  27,
	}

	fmt.Println(user2.Name)
	fmt.Println((*user2).Age)

	// ============================================================
	// 6. Pointer Receiver vs Value Receiver
	// ============================================================

	user3 := User{
		Name: "Mahdi",
		Age:  27,
	}

	// Pointer receiver:
	// The original struct is modified.

	user3.Birthday1()

	fmt.Println(user3)

	// Value receiver:
	// A copy of the struct is modified.
	// The original struct does not change.

	user3.Birthday2()

	fmt.Println(user3)
}
