package main

import (
	"encoding/json"
	"fmt"
	"math"
)

// =====================
// Domain Layer
// =====================

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

// =====================
// Infrastructure Layer
// =====================

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

	maxID := math.MinInt

	for _, product := range p.db {

		if product.ID > maxID {
			maxID = product.ID
		}

	}

	return maxID
}

func NewProductRepository() *ProductRepositoryImpl {

	return &ProductRepositoryImpl{
		db: []Product{
			{
				ID:    1,
				Name:  "Laptop",
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

// =====================
// Application Layer
// =====================

type CreateProductCmd struct {
	Name  string
	Price int
}

type CreateProductUseCase interface {
	Execute(cmd CreateProductCmd) *Product
}

type CreateProductService struct {
	productRepo ProductRepository
}

func NewCreateProductUseCase(
	repo ProductRepository,
) *CreateProductService {

	return &CreateProductService{
		productRepo: repo,
	}
}

func (c *CreateProductService) Execute(
	cmd CreateProductCmd,
) *Product {

	maxID := c.productRepo.FindMaxId()

	product := Product{
		ID:    maxID + 1,
		Name:  cmd.Name,
		Price: cmd.Price,
	}

	c.productRepo.Create(product)

	return &product
}

// ---------------------

type GetProductUseCase interface {
	Execute(id int) (*Product, error)
}

type GetProductService struct {
	productRepo ProductRepository
}

func NewGetProductUseCase(
	repo ProductRepository,
) *GetProductService {

	return &GetProductService{
		productRepo: repo,
	}
}

func (g *GetProductService) Execute(
	id int,
) (*Product, error) {

	product := g.productRepo.GetById(id)

	if product == nil {
		return nil, fmt.Errorf(
			"product with id %d not found",
			id,
		)
	}

	return product, nil
}

// =====================
// Presentation Layer
// =====================

type ProductController struct {
	createProduct CreateProductUseCase

	getProduct GetProductUseCase
}

func NewProductController(
	createProduct CreateProductUseCase,
	getProduct GetProductUseCase,
) *ProductController {

	return &ProductController{
		createProduct: createProduct,
		getProduct:    getProduct,
	}
}

type ProductResponseDto struct {
	ID int `json:"id"`

	Name string `json:"name"`

	Price int `json:"price"`
}

func (p *ProductController) CreateProduct(
	cmd CreateProductCmd,
) *ProductResponseDto {

	product := p.createProduct.Execute(cmd)

	return &ProductResponseDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}
}

func (p *ProductController) GetProduct(
	id int,
) (*ProductResponseDto, error) {

	product, err := p.getProduct.Execute(id)

	if err != nil {
		return nil, err
	}

	return &ProductResponseDto{
		ID:    product.ID,
		Name:  product.Name,
		Price: product.Price,
	}, nil
}

// =====================
// Composition Root
// =====================

func main() {

	// Infrastructure
	repository := NewProductRepository()

	// Application
	createProduct := NewCreateProductUseCase(repository)

	getProduct := NewGetProductUseCase(repository)

	// Presentation
	controller := NewProductController(
		createProduct,
		getProduct,
	)

	created := controller.CreateProduct(
		CreateProductCmd{
			Name:  "Monitor",
			Price: 50_000,
		},
	)

	fmt.Println(created)

	found, err := controller.GetProduct(
		4,
	)

	fmt.Println(found, err)
	jsonData, _ := json.Marshal(created)

	fmt.Println(string(jsonData))
}
