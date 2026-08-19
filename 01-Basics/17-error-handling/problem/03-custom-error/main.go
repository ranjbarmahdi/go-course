package main

import (
	"errors"
	"fmt"
)

type ValidationError struct {
	Field   string
	Message string
}

func (v ValidationError) Error() string {
	return fmt.Sprintf(
		"%s: %s",
		v.Field,
		v.Message,
	)
}
func CreateProduct(name string, price int) error {
	if name == "" {
		return ValidationError{
			Field:   "name",
			Message: "is required",
		}
	}

	if price <= 0 {
		return ValidationError{
			Field:   "price",
			Message: "must be positive",
		}
	}

	return nil
}

func main() {
	err := CreateProduct("", -10)

	var validationErr ValidationError
	if errors.As(err, &validationErr) {
		fmt.Println("Invalid field:", validationErr.Field)
		fmt.Println("Reason:", validationErr.Message)
	}
}
