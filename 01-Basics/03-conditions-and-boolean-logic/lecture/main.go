package main

import "fmt"

func main() {
	fmt.Println("1. if")
	age := 27
	if age < 13 {
		fmt.Println("Child")
	} else if age < 18 {
		fmt.Println("Teenager")
	} else {
		fmt.Println("Adult")
	}

	fmt.Println("2. Boolean operators")
	// AND: age >= 18 && age <= 65
	// OR: age < 18 || age > 65
	// NOT: !isDeveloper
}
