package main

import "fmt"

type Product struct {
	ID        int
	Name      string
	Price     float64
	Available bool
}

func main() {
	product := Product{
		ID:        1,
		Name:      "Laptop",
		Price:     1200.50,
		Available: true,
	}
	fmt.Printf("%v\n", product)
	fmt.Printf("%+v\n", product)
	fmt.Printf("Product %s costs %.2f\n", product.Name, product.Price)
}
