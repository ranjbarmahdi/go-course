package main

import (
	"fmt"
)

func main() {
	scores := map[string]int{
		"Ali":   85,
		"Sara":  92,
		"Reza":  78,
		"Mahdi": 95,
	}

	fmt.Println("Mahdi:", scores["Mahdi"])
	fmt.Println("Sara:", scores["Sara"])

	scores["Nima"] = 88
	scores["Ali"] = 90

	delete(scores, "Reza")

	fmt.Println(scores)
}
