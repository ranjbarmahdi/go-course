package main

import "fmt"

func filterEven(numbers []int) []int {
	even := []int{}
	for _, num := range numbers {
		if num%2 == 0 {
			even = append(even, num)
		}
	}
	return even
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6}
	fmt.Println(filterEven(numbers))
}
