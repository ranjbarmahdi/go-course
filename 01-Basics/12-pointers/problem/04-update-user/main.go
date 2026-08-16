package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) birthday() {
	(*u).Age++
}

func main() {
	user := User{
		Name: "Mahdi",
		Age:  26,
	}
	user.birthday()
	fmt.Println(user)
}
