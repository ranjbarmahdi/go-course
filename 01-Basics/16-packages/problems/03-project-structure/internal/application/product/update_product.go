package product

import (
	"fmt"
	domainProduct "go-course/01-Basics/16-packages/problems/03-project-structure/internal/domain/product"
)

type UpdateProductCmd struct {
	ID    int
	Name  string
	Price int
}

type UpdateProduct struct {
	repo domainProduct.Repository
}

func (u *UpdateProduct) Execute(cmd UpdateProductCmd) (*domainProduct.Product, error) {
	found := u.repo.GetByID(cmd.ID)
	if found == nil {
		return nil, fmt.Errorf("product with id %d not found", cmd.ID)
	}

	found.Update(cmd.Name, cmd.Price)

	u.repo.Update(found)

	return found, nil
}

func NewUpdateProduct(repo domainProduct.Repository) *UpdateProduct {
	return &UpdateProduct{repo: repo}
}
