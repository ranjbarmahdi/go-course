package main

import "fmt"

func swap(a *int, b *int) {
	tempAdd := *a
	*a = *b
	*b = tempAdd
}

func main() {
	a := 10
	b := 20
	swap(&a, &b)
	fmt.Println("a:", a)
	fmt.Println("b:", b)
}
