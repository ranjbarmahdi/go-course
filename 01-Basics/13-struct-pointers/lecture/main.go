package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) Birthday1() {
	u.Age++
}

func (u User) Birthday2() {
	u.Age++
}

func main() {

	fmt.Println("1. Pointer to Struct")
	user := User{
		Name: "Mahdi",
		Age:  27,
	}

	userPtr := &user

	fmt.Println(userPtr.Name)
	fmt.Println((*userPtr).Age)

	fmt.Println("2. Creating Struct With Pointer")
	user2 := &User{
		Name: "Mahdi",
		Age:  27,
	}

	fmt.Println(user2.Name)
	fmt.Println((*user2).Age)

	fmt.Println("3. Pointer Receiver vs Value Receiver")
	user3 := User{
		Name: "Mahdi",
		Age:  27,
	}
	user3.Birthday1()
	fmt.Println(user3)

	user3.Birthday2()
	fmt.Println(user3)
}
