package main

import "fmt"

func main() {
	source := []int{10, 20, 30, 40, 50}
	destination := make([]int, len(source))

	copy(destination, source)

	fmt.Println(destination)
}
