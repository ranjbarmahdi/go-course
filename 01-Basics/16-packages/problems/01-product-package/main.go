package main

import (
	"fmt"
	"go-course/01-Basics/16-packages/problems/01-product-package/product"
)

func main() {
	p := product.NewProduct(1, "Product", 200_000)
	fmt.Println(p)
}
