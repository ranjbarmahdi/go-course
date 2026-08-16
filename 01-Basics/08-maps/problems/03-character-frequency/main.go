package main

import "fmt"

func main() {
	text := "hello"
	frequency := map[string]int{}

	for _, char := range text {
		if _, ok := frequency[string(char)]; ok {
			frequency[string(char)]++
		} else {
			frequency[string(char)] = 1
		}
	}

	for key, value := range frequency {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
