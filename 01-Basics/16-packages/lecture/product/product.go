// ============================================================
// product/product.go
// ============================================================

package product

// ============================================================
// 1. Product Struct
// ============================================================
// A package can define and expose its own types.
//
// Fields starting with an uppercase letter are exported.
// Fields starting with a lowercase letter are unexported.

type Product struct {
	ID    int
	Name  string
	price int
}

// ============================================================
// 2. Constructor Function
// ============================================================
// Go does not have constructors like some other languages.
//
// A common pattern is to create a function such as NewProduct.
//
// The function can initialize the struct and return it.

func NewProduct(id int, name string, price int) Product {

	return Product{
		ID:    id,
		Name:  name,
		price: price,
	}
}
