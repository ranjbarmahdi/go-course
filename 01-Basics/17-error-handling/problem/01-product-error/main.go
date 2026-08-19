package main

import (
	"errors"
	"fmt"
)

type Product struct {
	ID   int
	Name string
}

var ErrProductNotFound = errors.New("product not found")

func FindProduct(id int) (*Product, error) {
	if id == 1 {
		return &Product{
			ID:   1,
			Name: "Product",
		}, nil
	}
	return nil, ErrProductNotFound
}

func main() {
	product, err := FindProduct(2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(product)
}
