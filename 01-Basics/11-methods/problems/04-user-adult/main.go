package main

import "fmt"

type User struct {
	Name   string
	Family string
	Age    int
}

func (u User) IsAdult() bool {
	return u.Age >= 18
}

func main() {
	user := User{
		Name:   "Mahdi",
		Family: "Ranjbar",
		Age:    27,
	}

	fmt.Println(user.IsAdult())
}
