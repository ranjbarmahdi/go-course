package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price int
	Stock int
}

func main() {
	lapTop := Product{
		ID:    1,
		Name:  "Laptop",
		Price: 2_000,
		Stock: 5,
	}

	mouse := Product{
		ID:    2,
		Name:  "Mouse",
		Price: 120,
		Stock: 20,
	}

	keyboard := Product{
		ID:    3,
		Name:  "Keyboard",
		Price: 40,
		Stock: 10,
	}

	lapTop.Stock = 4

	fmt.Println(lapTop)
	fmt.Println(mouse)
	fmt.Println(keyboard)
}
