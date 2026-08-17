package product

type Product struct {
	ID    int
	Name  string
	Price int
}

func NewProduct(
	id int,
	name string,
	price int,
) Product {

	return Product{
		ID:    id,
		Name:  name,
		Price: price,
	}
}

func (p *Product) Update(name string, price int) {
	p.Name = name
	p.Price = price
}
