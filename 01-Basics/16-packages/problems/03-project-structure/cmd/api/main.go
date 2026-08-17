package main

import (
	"fmt"

	appProduct "go-course/01-Basics/16-packages/problems/03-project-structure/internal/application/product"

	infraProduct "go-course/01-Basics/16-packages/problems/03-project-structure/internal/infrastructure/product"
)

func main() {

	repository := infraProduct.NewMemoryRepository()

	createProduct := appProduct.NewCreateProduct(
		repository,
	)

	product := createProduct.Execute(
		appProduct.CreateProductCommand{
			Name:  "Keyboard",
			Price: 20000,
		},
	)

	getProduct := appProduct.NewGetProduct(repository)
	found, err := getProduct.Execute(product.ID)
	fmt.Println(found, err)

	updateProduct := appProduct.NewUpdateProduct(repository)
	updateCmd := appProduct.UpdateProductCmd{
		ID:    product.ID,
		Name:  "Updated Product",
		Price: 900_000,
	}
	updated, err := updateProduct.Execute(updateCmd)

	fmt.Println(updated, err)

}
