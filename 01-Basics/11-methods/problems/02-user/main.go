package main

import "fmt"

type User struct {
	Name   string
	Family string
	Age    int
}

func (u User) FullName() string {
	return u.Name + " " + u.Family
}
func (u User) Greet() {
	fmt.Printf("Hello, my name is %s and I am %d years old.\n", u.FullName(), u.Age)
}

func main() {
	user := User{
		Name:   "Mahdi",
		Family: "Ranjbar",
		Age:    27,
	}

	user.Greet()
}
