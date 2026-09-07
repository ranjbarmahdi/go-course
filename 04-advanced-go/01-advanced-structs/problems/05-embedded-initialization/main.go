/*
Problem:

Create:

type Contact struct {
    Email string
    Phone string
}

Create:

type Customer struct {
    Name string
    Contact
}

Requirements:

- Embed Contact inside Customer.
- Create a Customer with:

    Name:  "Mahdi"
    Email: "mahdi@example.com"
    Phone: "123456789"

- Initialize Contact using its type name.
- Access Email directly through Customer.
- Access Phone directly through Customer.

Use field promotion.

Expected output:

Mahdi
mahdi@example.com
123456789
*/

package main

import "fmt"

type Contact struct {
	Email string
	Phone string
}

type Customer struct {
	Name string
	Contact
}

func main() {
	customer := Customer{
		Name: "Mahdi",
		Contact: Contact{
			Phone: "1234567",
			Email: "mahdi@example.com",
		},
	}
	fmt.Println(customer.Name)
	fmt.Println(customer.Email)
	fmt.Println(customer.Phone)
}
