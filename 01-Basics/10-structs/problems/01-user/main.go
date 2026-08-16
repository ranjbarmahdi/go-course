package main

import "fmt"

type User struct {
	ID     int
	Name   string
	Age    int
	Email  string
	Active bool
}

func main() {
	user := User{
		ID:     1,
		Name:   "Mahdi",
		Age:    28,
		Email:  "mahdi@example.com",
		Active: true,
	}
	fmt.Println("ID ->", user.ID)
	fmt.Println("Name ->", user.Name)
	fmt.Println("Age ->", user.Age)
	fmt.Println("Email ->", user.Email)
	fmt.Println("Active ->", user.Active)
}
