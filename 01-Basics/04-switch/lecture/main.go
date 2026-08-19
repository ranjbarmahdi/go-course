package main

import "fmt"

func main() {

	// ============================================================
	// 1. Basic Switch
	// ============================================================

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

	// ============================================================
	// 2. Multiple Values in One Case
	// ============================================================

	switch day {
	case 6, 7:
		fmt.Println("Weekend")
	default:
		fmt.Println("Weekday")
	}

	// ============================================================
	// 3. Switch Without an Expression
	// ============================================================

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

	// ============================================================
	// 4. Switch with Strings
	// ============================================================

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

	// ============================================================
	// 5. Switch with Multiple Conditions
	// ============================================================

	score := 85

	switch {
	case score >= 90:
		fmt.Println("Excellent")
	case score >= 70:
		fmt.Println("Good")
	case score >= 50:
		fmt.Println("Pass")
	default:
		fmt.Println("Fail")
	}

	// ============================================================
	// 6. Switch with Initialization
	// ============================================================

	switch status := 200; status {
	case 200:
		fmt.Println("OK")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Internal Server Error")
	default:
		fmt.Println("Unknown status")
	}
}
