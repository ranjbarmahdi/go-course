package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i <= 100; i++ {
		sum += i
	}
	fmt.Println("Sum of 1 to 100 =", sum)
}
