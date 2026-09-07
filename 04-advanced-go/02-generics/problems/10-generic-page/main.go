/*
Problem:

Create a generic struct:

type Page[T any] struct {
    Items []T
    Total int
}

The Page represents a paginated result.

Requirements:

- Create Page[T].
- Create Page[int].
- Store:
    Items: []int{10, 20, 30}
    Total: 100

- Create Page[string].
- Store:
    Items: []string{"Ali", "Mahdi", "Sara"}
    Total: 50

- Print the items and total of both pages.

Expected output:

[10 20 30]
100
[Ali Mahdi Sara]
50

Backend concept:

A generic Page[T] can be reused for:

Page[User]
Page[Product]
Page[Order]

without creating separate Page structs for each type.
*/

package main

import "fmt"

type Page[T any] struct {
	Items []T
	Total int
}

func main() {
	paginatedStringResponse := Page[string]{
		Items: []string{"Ali", "Mahdi", "Sara"},
		Total: 100,
	}

	paginatedIntegerResponse := Page[int]{
		Items: []int{10, 20, 30},
		Total: 50,
	}

	fmt.Println(paginatedStringResponse.Items)
	fmt.Println(paginatedStringResponse.Total)

	fmt.Println(paginatedIntegerResponse.Items)
	fmt.Println(paginatedIntegerResponse.Total)
}
