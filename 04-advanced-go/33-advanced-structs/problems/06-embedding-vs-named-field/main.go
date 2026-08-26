/*
Problem:

Create:

type Address struct {
    City string
}

Create two different User structs.

First:

type UserWithNamedAddress struct {
    Name    string
    Address Address
}

Second:

type UserWithEmbeddedAddress struct {
    Name string
    Address
}

Requirements:

- Create one value of each User type.
- Set Name = "Mahdi".
- Set City = "Baku".

For UserWithNamedAddress:

- Access the city using:

    user.Address.City

For UserWithEmbeddedAddress:

- Access the city using field promotion:

    user.City

Print the Name and City for both users.

Expected output:

Named:
Mahdi
Baku

Embedded:
Mahdi
Baku

Important:

Named field:

    Address Address

requires:

    user.Address.City

Embedded field:

    Address

allows:

    user.City
*/

package main

import "fmt"

type Address struct {
	City string
}

type UserWithNamedAddress struct {
	Name    string
	Address Address
}

type UserWithEmbeddedAddress struct {
	Name string
	Address
}

func main() {
	user1 := UserWithNamedAddress{
		Name: "Mahdi",
		Address: Address{
			City: "Baku",
		},
	}

	user2 := UserWithEmbeddedAddress{
		Name: "Mahdi",
		Address: Address{
			City: "Baku",
		},
	}

	fmt.Println("Named:")
	fmt.Println(user1.Name)
	fmt.Println(user1.Address.City)

	fmt.Println("Embedded:")
	fmt.Println(user2.Name)
	fmt.Println(user2.City)
}
