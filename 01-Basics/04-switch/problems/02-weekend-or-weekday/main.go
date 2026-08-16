package main

import "fmt"

func main() {
	day := 5

	switch day {
	case 1, 2, 3, 4, 5:
		fmt.Println("Weekday")
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Invalid input.")
	}
}
