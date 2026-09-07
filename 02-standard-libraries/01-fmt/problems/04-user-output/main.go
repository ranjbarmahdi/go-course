package main

import "fmt"

type User struct {
	ID    int
	Name  string
	Email string
}

func PrintUser(user User) {
	fmt.Println("ID:", user.ID)
	fmt.Println("Name:", user.Name)
	fmt.Println("Email:", user.Email)
}

func main() {
	user := User{
		ID:    1,
		Name:  "Mahdi",
		Email: "mehdi@test.com",
	}
	PrintUser(user)
}
