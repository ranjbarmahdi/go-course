package product

type Repository interface {
	Create(product Product) int

	FindMaxID() int

	GetByID(id int) *Product

	Update(product *Product) *Product
}
