package main

import "fmt"

type Product struct {
	Name  string
	Price float64
	Stock int
}

func (p Product) TotalValue() float64 {
	return p.Price * float64(p.Stock)
}

func main() {
	p := Product{
		Name:  "Mahdi",
		Price: 150,
		Stock: 10,
	}
	fmt.Println(p.TotalValue())
}
