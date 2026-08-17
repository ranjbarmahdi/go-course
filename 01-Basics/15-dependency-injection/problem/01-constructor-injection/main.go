package main

import (
	"fmt"
	"math"
)

// Domain Layer
type Product struct {
	ID    int
	Name  string
	Price int
}

type ProductRepository interface {
	Create(entity Product) int
	GetById(id int) *Product
	FindMaxId() int
}

// Infra Layer
type ProductRepositoryImpl struct {
	db []Product
}

func (p *ProductRepositoryImpl) Create(entity Product) int {
	p.db = append(p.db, entity)
	return entity.ID
}

func (p *ProductRepositoryImpl) GetById(id int) *Product {
	for i := range p.db {
		if p.db[i].ID == id {
			return &p.db[i]
		}
	}
	return nil
}

func (p *ProductRepositoryImpl) FindMaxId() int {
	if len(p.db) == 0 {
		return 0
	}

	maxId := math.MinInt
	for _, v := range p.db {
		if v.ID > maxId {
			maxId = v.ID
		}
	}
	return maxId
}

func NewProductRepositoryImpl() *ProductRepositoryImpl {
	return &ProductRepositoryImpl{
		db: []Product{
			{
				ID:    1,
				Name:  "LapTop",
				Price: 100_000,
			},
			{
				ID:    2,
				Name:  "Mouse",
				Price: 10_000,
			},
			{
				ID:    3,
				Name:  "Keyboard",
				Price: 20_000,
			},
		},
	}
}

// Application Layer
type CreateProductCmd struct {
	Name  string
	Price int
}

type CreateProduct interface {
	Execute(cmd CreateProductCmd) *Product
}

type CreateProductUc struct {
	productRepo ProductRepository
}

func NewCreateProductUc(repo ProductRepository) *CreateProductUc {
	return &CreateProductUc{
		productRepo: repo,
	}
}

func (c CreateProductUc) Execute(cmd CreateProductCmd) *Product {
	maxId := c.productRepo.FindMaxId()
	newProduct := Product{
		ID:    maxId + 1,
		Name:  cmd.Name,
		Price: cmd.Price,
	}

	c.productRepo.Create(newProduct)
	return &newProduct
}

type GetProduct interface {
	Execute(id int) (*Product, error)
}

type GetProductUc struct {
	productRepo ProductRepository
}

func NewGetProductUc(repo ProductRepository) *GetProductUc {
	return &GetProductUc{
		productRepo: repo,
	}
}

func (g GetProductUc) Execute(id int) (*Product, error) {
	found := g.productRepo.GetById(id)
	if found == nil {
		return nil, fmt.Errorf("Product With %d Id Not Found", id)
	}
	return found, nil
}

func main() {
	repo := NewProductRepositoryImpl()
	createProductUc := NewCreateProductUc(repo)
	getProductUc := NewGetProductUc(repo)

	createCmd := CreateProductCmd{
		Name:  "New Product",
		Price: 99_000,
	}

	createdProduct := createProductUc.Execute(createCmd)
	fmt.Println(createdProduct)

	foundCreatedProduct, err := getProductUc.Execute(10)
	fmt.Println(foundCreatedProduct, err)
}
