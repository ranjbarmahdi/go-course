/*
Problem:

Create a generic struct:

type Result[T any] struct {
    Data  T
    Error error
}

The Result represents the result of an operation.

Requirements:

- Create Result[T].
- Create a Result[int] with:
    Data: 100
    Error: nil

- Create a Result[string] with:
    Data: "Mahdi"
    Error: nil

- Print Data from both results.
- Print Error from both results.

Expected output:

100
<nil>
Mahdi
<nil>

Backend concept:

A generic Result[T] can represent different
types of successful results:

Result[User]
Result[Product]
Result[Order]

while keeping the same structure.
*/

package main

import "fmt"

type Result[T any] struct {
	Data  T
	Error error
}

func main() {
	stringResult := Result[string]{
		Data:  "Mahdi",
		Error: nil,
	}

	integerResult := Result[int]{
		Data:  100,
		Error: nil,
	}

	fmt.Println(integerResult.Data)
	fmt.Println(integerResult.Error)

	fmt.Println(stringResult.Data)
	fmt.Println(stringResult.Error)
}
