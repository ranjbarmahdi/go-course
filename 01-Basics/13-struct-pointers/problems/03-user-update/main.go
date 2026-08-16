package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func (u *User) changeName(name string) {
	u.Name = name
}

func main() {
	user := User{
		Name: "Mahdi",
		Age:  27,
	}

	user.changeName("Ali")
	fmt.Println(user)
}
