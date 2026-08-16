package main

import "fmt"

func main() {
	numbers := [8]int{10, 25, 7, 42, 18, 99, 31, 50}
	target := 420

	index := -1
	for i, v := range numbers {
		if v == target {
			index = i
			break
		}
	}

	if index >= 0 {
		fmt.Printf("%d found at index %d", target, index)
	} else {
		fmt.Printf("%d not found", target)
	}
}
