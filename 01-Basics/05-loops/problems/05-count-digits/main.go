package main

import "fmt"

func main() {
	digit := 123456
	count := 0
	for digit > 0 {
		digit = digit / 10
		count++
	}
	fmt.Printf("%d digits.", count)
}
