package product

import (
	"fmt"
	domainProduct "go-course/01-Basics/16-packages/problems/03-project-structure/internal/domain/product"
)

type GetProduct struct {
	repo domainProduct.Repository
}

func (g *GetProduct) Execute(id int) (*domainProduct.Product, error) {
	found := g.repo.GetByID(id)
	if found == nil {
		return nil, fmt.Errorf("product with %d id not found.", id)
	}

	return found, nil
}

func NewGetProduct(repo domainProduct.Repository) *GetProduct {
	return &GetProduct{repo: repo}
}
