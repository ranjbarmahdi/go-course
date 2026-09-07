/*
Problem:

Create a constraint:

type Number interface {
	~int | ~int64 | ~float64
}

The `~` allows not only the exact types,
but also custom types whose underlying type
is one of these types.

Create:

type UserID int
type Price float64

Create:

func Add[T Number](a T, b T) T

Requirements:

- Use `~int`.
- Use `~int64`.
- Use `~float64`.
- Create UserID based on int.
- Create Price based on float64.
- Use Add() with UserID.
- Use Add() with Price.
- Print the results.

Expected output:

30
30.5

Important:

Without `~`:

    int

would allow only int.

With:

    ~int

it also allows types such as:

    type UserID int
*/

package main

import "fmt"

type Number interface {
	~int | ~int64 | ~float64
}

type UserID int

type Price float64

func Add[T Number](a T, b T) T {
	return a + b
}

func main() {
	var userID UserID = 10
	var userID2 UserID = 20

	var price Price = 10.5
	var price2 Price = 20.0

	fmt.Println(Add(userID, userID2))
	fmt.Println(Add(price, price2))

}
