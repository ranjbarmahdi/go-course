package main

import "fmt"

func changeValue(a *int, b int) {
	*a = b
}

func main() {
	number := 100
	changeValue(&number, 200)
	fmt.Println(number)
}
