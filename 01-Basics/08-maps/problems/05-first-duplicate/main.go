package main

import "fmt"

func main() {
	numbers := []int{5, 3, 8, 2, 3, 9, 8}
	seen := map[int]bool{}

	for _, number := range numbers {
		if _, see := seen[number]; see {
			fmt.Println("First duplication is", number)
			break
		}
		seen[number] = true
	}
}
