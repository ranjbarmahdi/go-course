package product

import domainProduct "go-course/01-Basics/16-packages/problems/03-project-structure/internal/domain/product"

type CreateProduct struct {
	repo domainProduct.Repository
}

func NewCreateProduct(
	repo domainProduct.Repository,
) *CreateProduct {

	return &CreateProduct{
		repo: repo,
	}

}

type CreateProductCommand struct {
	Name string

	Price int
}

func (c *CreateProduct) Execute(
	cmd CreateProductCommand,
) *domainProduct.Product {

	id := c.repo.FindMaxID() + 1

	product := domainProduct.NewProduct(
		id,
		cmd.Name,
		cmd.Price,
	)

	c.repo.Create(product)

	return &product

}
