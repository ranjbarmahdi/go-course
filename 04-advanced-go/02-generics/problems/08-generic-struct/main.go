/*
Problem:

Create a generic struct:

type Box[T any] struct {
    Value T
}

Requirements:

- Create Box[int].
- Create Box[string].
- Initialize both.
- Print their values.

Example:

Box[int]:

Value = 100

Box[string]:

Value = "Mahdi"

Expected output:

100
Mahdi
*/

package main

import "fmt"

type Box[T any] struct {
	Value T
}

func main() {
	boxInt := Box[int]{Value: 100}
	boxString := Box[string]{Value: "Mahdi"}

	fmt.Println(boxInt)
	fmt.Println(boxString)
	fmt.Println(boxInt.Value)
	fmt.Println(boxString.Value)
}
