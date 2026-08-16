package main

import "fmt"

func main() {
	number := 2

	isPrime := true
	i := 2
	for isPrime && i < number/2 {
		if number%i == 0 {
			isPrime = false
		}
		i++
	}

	if isPrime {
		fmt.Printf("Number %d is prime.", number)
	} else {
		fmt.Printf("Number %d is not prime.", number)
	}
}
