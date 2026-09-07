/*
Problem:

Create two structs:

type Address struct {
    City    string
    Country string
}

type User struct {
    Name    string
    Age     int
    Address Address
}

Create a User with:

Name:    "Mahdi"
Age:     27
City:    "Baku"
Country: "Azerbaijan"

Requirements:

- Use struct composition.
- Address must be a named field, not embedded.
- Print the user's Name.
- Print the user's Age.
- Print the user's City.
- Print the user's Country.

Expected output:

Mahdi
27
Baku
Azerbaijan
*/

package main

import "fmt"

type Address struct {
	City    string
	Country string
}

type User struct {
	Name    string
	Age     int
	Address Address
}

func main() {
	user := User{
		Name: "Mahdi",
		Age:  27,
		Address: Address{
			City:    "Baku",
			Country: "Azerbaijan",
		},
	}

	fmt.Println(user.Name)
	fmt.Println(user.Age)
	fmt.Println(user.Address.Country)
	fmt.Println(user.Address.City)
}
