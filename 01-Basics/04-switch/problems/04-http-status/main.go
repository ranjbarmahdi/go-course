package main

import "fmt"

func main() {
	status := 404

	switch status {
	case 200:
		fmt.Println("OK")
	case 201:
		fmt.Println("Created")
	case 400:
		fmt.Println("Bad Request")
	case 401:
		fmt.Println("Unauthorized")
	case 403:
		fmt.Println("Forbidden")
	case 404:
		fmt.Println("Not Found")
	case 500:
		fmt.Println("Internal Server Error")
	default:
		fmt.Println("Unknown Status.")
	}
}
