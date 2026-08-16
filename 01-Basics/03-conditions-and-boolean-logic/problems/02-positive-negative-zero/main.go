package main

import "fmt"

func main() {
	number := -10

	if number > 0 {
		fmt.Println("Positive")
	} else if number == 0 {
		fmt.Println("Zero")
	} else {
		fmt.Println("Negative")
	}
}
