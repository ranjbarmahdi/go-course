package main

import "fmt"

func main() {
	command := "start"
	switch command {
	case "start":
		fmt.Println("Starting application...")
	case "stop":
		fmt.Println("Stopping application...")
	case "restart":
		fmt.Println("Restarting application...")
	case "status":
		fmt.Println("Application is running.")
	default:
		fmt.Println("Unknown command")
	}
}
