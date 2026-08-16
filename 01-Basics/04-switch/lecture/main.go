package main

import "fmt"

func main() {
	fmt.Println("1. Basic Switch")
	day := 1

	switch day {
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednesday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid day")
	}

	fmt.Print("2. Multiple values in one case")
	switch day {
	case 6, 7:
		fmt.Printf("Weekend")
	default:
		fmt.Println("Weekday")
	}

	fmt.Print("3. Switch without an expression")
	age := 27
	switch {
	case age < 13:
		fmt.Println("Child")
	case age < 18:
		fmt.Println("Teenager")
	case age < 65:
		fmt.Println("Adult")
	default:
		fmt.Println("Senior")
	}

	fmt.Print("4. Switch with strings")
	command := "start"
	switch command {
	case "start":
		fmt.Println("Starting...")
	case "stop":
		fmt.Println("Stopping...")
	case "restart":
		fmt.Println("Restarting...")
	default:
		fmt.Println("Unknown command")
	}
}
