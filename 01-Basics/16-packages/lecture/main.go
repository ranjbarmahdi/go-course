package main

import (
	"fmt"
	"go-course/01-Basics/16-packages/lecture/product"
)

func main() {

	p := product.NewProduct(
		1,
		"Laptop",
		1000,
	)

	fmt.Println(p)

}
