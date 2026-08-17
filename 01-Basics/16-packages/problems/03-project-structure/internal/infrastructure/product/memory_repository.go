package product

import domainProduct "go-course/01-Basics/16-packages/problems/03-project-structure/internal/domain/product"

type MemoryRepository struct {
	db []domainProduct.Product
}

func NewMemoryRepository() *MemoryRepository {

	return &MemoryRepository{
		db: []domainProduct.Product{},
	}

}

func (m *MemoryRepository) Create(
	product domainProduct.Product,
) int {

	m.db = append(
		m.db,
		product,
	)

	return product.ID
}

func (m *MemoryRepository) FindMaxID() int {

	maxID := 0

	for _, p := range m.db {

		if p.ID > maxID {
			maxID = p.ID
		}

	}

	return maxID
}

func (m *MemoryRepository) GetByID(
	id int,
) *domainProduct.Product {

	for i := range m.db {

		if m.db[i].ID == id {
			return &m.db[i]
		}

	}

	return nil
}

func (m *MemoryRepository) Update(product *domainProduct.Product) *domainProduct.Product {
	for i := range m.db {
		if m.db[i].ID == product.ID {
			m.db[i] = *product
			return &m.db[i]
		}
	}
	return nil
}

var _ domainProduct.Repository = (*MemoryRepository)(nil)
