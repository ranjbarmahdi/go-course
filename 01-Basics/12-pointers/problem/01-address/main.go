package main

import "fmt"

func main() {
	x := 100

	fmt.Println("Value:", x)
	fmt.Println("Address:", &x)
	fmt.Println("Pointer value:", *&x)
}
