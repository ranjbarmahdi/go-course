package main

import "fmt"

func main() {
	score := -10

	if score >= 90 {
		fmt.Println("A")
	} else if score >= 80 {
		fmt.Println("B")
	} else if score >= 70 {
		fmt.Println("C")
	} else if score >= 60 {
		fmt.Println("D")
	} else if score >= 0 {
		fmt.Println("F")
	} else {
		fmt.Println("Invalid Input.")
	}
}
