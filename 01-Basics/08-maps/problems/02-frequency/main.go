package main

import "fmt"

func main() {
	numbers := []int{1, 2, 2, 3, 3, 3, 4, 4, 4, 4}
	frequency := map[int]int{}
	for _, number := range numbers {
		if _, ok := frequency[number]; ok {
			frequency[number]++
		} else {
			frequency[number] = 1
		}
	}

	for key, value := range frequency {
		fmt.Printf("%d -> %d\n", key, value)
	}
}
