package main

import (
	"errors"
	"fmt"
)

var ErrProductNotFound = errors.New(
	"product not found repo",
)

type Product struct {
	ID int
}

func FindProductRepo(
	id int,
) (*Product, error) {

	return nil, ErrProductNotFound
}

func FindProductUseCase(
	id int,
) (*Product, error) {

	res, err := FindProductRepo(id)

	if err != nil {

		return nil, fmt.Errorf(
			"get product use case failed: %w",
			err,
		)

	}

	return res, nil
}

func FindProductController(
	id int,
) (*Product, error) {

	res, err := FindProductUseCase(id)

	if err != nil {

		return nil, fmt.Errorf(
			"controller failed: %w",
			err,
		)

	}

	return res, nil
}

func main() {

	res, err := FindProductController(10)

	fmt.Println(res)
	fmt.Println(err)

	if errors.Is(err, ErrProductNotFound) {

		fmt.Println(
			"HTTP 404 Not Found",
		)

	}

}
