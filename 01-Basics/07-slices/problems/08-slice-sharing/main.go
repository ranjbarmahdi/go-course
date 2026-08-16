package main

import "fmt"

func main() {
	numbers := []int{10, 20, 30, 40, 50}

	a := numbers[1:4]

	a[0] = 200

	fmt.Println(numbers)
	fmt.Println(a)
}
