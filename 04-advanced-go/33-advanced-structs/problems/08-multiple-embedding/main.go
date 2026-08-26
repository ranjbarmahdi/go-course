/*
Problem:

Create:

type Person struct {
    Name string
}

type Contact struct {
    Email string
}

type User struct {
    Person
    Contact
}

Requirements:

- Embed both Person and Contact inside User.
- Create a User with:

    Name:  "Mahdi"
    Email: "mahdi@example.com"

- Access Name directly using field promotion.
- Access Email directly using field promotion.

Print:

Mahdi
mahdi@example.com

Important:

A struct can embed multiple structs.

User receives promoted fields from both:

    Person → Name
    Contact → Email
*/

package main

import "fmt"

type Person struct {
	Name string
}

type Contact struct {
	Email string
}

type User struct {
	Person
	Contact
}

func main() {
	user := User{
		Person: Person{
			Name: "Mahdi",
		},
		Contact: Contact{
			Email: "mahdi@example.com",
		},
	}

	fmt.Println(user.Name)
	fmt.Println(user.Email)
}
