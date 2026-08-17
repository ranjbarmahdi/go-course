package main

import (
	"fmt"
	"go-course/01-Basics/16-packages/problems/02-private-public/user"
)

func main() {
	user := user.NewUser("Mahdi", 27)
	fmt.Println(user)
}
