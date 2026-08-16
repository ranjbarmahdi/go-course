package main

import "fmt"

func main() {
	fmt.Println("1. Create a Map")
	users := map[string]int{
		"Mahdi": 27,
		"Ali":   30,
		"Sara":  25,
	}

	fmt.Println(users)
	fmt.Println(users["Mahdi"])

	delete(users, "Mahdi")
	fmt.Println(users)
}
