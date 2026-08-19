package main

import (
	"errors"
	"fmt"
)

// ============================================================
// 1. The error Type
// ============================================================
// `error` is a built-in interface.
//
// Its definition is:
//
// type error interface {
//     Error() string
// }
//
// Any type that implements Error() string can be used as an error.

type MyError struct {
	message string
}

func (e MyError) Error() string {
	return e.message
}

// ============================================================
// 2. Creating Simple Errors
// ============================================================
// `errors.New()` creates a simple error value.

func divide(a, b int) (int, error) {

	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	return a / b, nil
}

// ============================================================
// 3. fmt.Errorf
// ============================================================
// `fmt.Errorf()` creates a formatted error.
//
// Useful when the error message contains dynamic values.

func GetProduct(id int) error {

	return fmt.Errorf(
		"product with id %d not found",
		id,
	)
}

// ============================================================
// 4. Sentinel Errors
// ============================================================
// A sentinel error is a predefined error value.
//
// It can be compared later using errors.Is().

var ErrProductNotFound = errors.New(
	"product not found",
)

func Repository() error {
	return ErrProductNotFound
}

// ============================================================
// 5. Error Wrapping
// ============================================================
// `%w` wraps an existing error.
//
// The original error can still be detected using errors.Is().

func UseCase() error {

	err := Repository()

	if err != nil {
		return fmt.Errorf(
			"get product failed: %w",
			err,
		)
	}

	return nil
}

// ============================================================
// 6. Custom Error
// ============================================================
// A custom error can contain additional information.
//
// Because ValidationError implements Error() string,
// it satisfies the error interface.

type ValidationError struct {
	Field   string
	Value   any
	Message string
}

func (e ValidationError) Error() string {
	return e.Field + ": " + e.Message
}

func ValidatePrice(price int) error {

	if price <= 0 {
		return ValidationError{
			Field:   "price",
			Value:   price,
			Message: "must be positive",
		}
	}

	return nil
}

func main() {

	// ============================================================
	// 1. The error Type
	// ============================================================

	var err error

	err = MyError{
		message: "something went wrong",
	}

	fmt.Println(err)

	// ============================================================
	// 2. Creating Simple Errors
	// ============================================================

	err2 := errors.New("product not found")

	fmt.Println(err2)

	// ============================================================
	// 3. Returning Errors from Functions
	// ============================================================

	res, err := divide(10, 0)

	fmt.Println(res, err)

	// ============================================================
	// 4. fmt.Errorf
	// ============================================================

	fmt.Println(GetProduct(10))

	// ============================================================
	// 5. Wrap Error
	// ============================================================

	err3 := UseCase()

	fmt.Println(err3)

	if errors.Is(err3, ErrProductNotFound) {
		fmt.Println("return HTTP 404")
	}

	// ============================================================
	// 6. Custom Error
	// ============================================================

	err4 := ValidatePrice(0)

	if err4 != nil {

		fmt.Println(err4)

		// ========================================================
		// 7. errors.As
		// ========================================================
		// errors.As checks whether an error is a specific type.
		//
		// It also gives us access to the custom error's fields.

		var validationError ValidationError

		if errors.As(err4, &validationError) {

			fmt.Println(
				"Invalid field:",
				validationError.Field,
			)

			fmt.Println(
				"Value:",
				validationError.Value,
			)

			fmt.Println(
				"Reason:",
				validationError.Message,
			)
		}
	}
}
