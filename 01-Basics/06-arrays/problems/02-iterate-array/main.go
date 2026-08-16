package main

import "fmt"

func main() {
	numbers := [5]int{10, 20, 30, 40, 50}
	for index, value := range numbers {
		fmt.Printf("Index: %d, Value: %d\n", index, value)
	}
}
