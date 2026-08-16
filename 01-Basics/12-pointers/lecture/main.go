package main

import "fmt"

func main() {
	fmt.Println("1. Memory Address")
	x := 10
	fmt.Println(x)
	fmt.Println(&x)

	fmt.Println("2. Pointer Variable")
	x1 := 10

	var p1 *int
	p1 = &x1
	fmt.Println(x1)
	fmt.Println(p1)
	fmt.Println(*p1)

	fmt.Println("3. Changing Through Pointer")
	x2 := 10
	var p *int = &x2
	*p = 50
	fmt.Println(x2)
	fmt.Println(*p)
}
