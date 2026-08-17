package product

type Product struct {
	ID    int
	Name  string
	Price int
}

func NewProduct(id int, name string, price int) *Product {
	return &Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}
