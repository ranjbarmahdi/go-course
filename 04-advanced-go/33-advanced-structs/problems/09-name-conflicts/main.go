/*
Problem:

Create:

type Person struct {
    Name string
}

type Company struct {
    Name string
}

Create:

type Employee struct {
    Person
    Company
}

Requirements:

- Embed both Person and Company.
- Create an Employee with:

    Person.Name  = "Mahdi"
    Company.Name = "Vardast"

- Try to access:

    employee.Name

You will find that this is ambiguous because both
embedded structs have a Name field.

Therefore:

- Access the person's name using:

    employee.Person.Name

- Access the company's name using:

    employee.Company.Name

Print:

Mahdi
Vardast

Important:

When multiple embedded structs contain the same field name,
Go does NOT choose one automatically.

You must explicitly specify which embedded struct
you want to access.
*/

package main

import "fmt"

type Person struct {
	Name string
}

type Company struct {
	Name string
}

type Employee struct {
	Person
	Company
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "Mahdi",
		},
		Company: Company{
			Name: "Npath",
		},
	}

	fmt.Println(emp.Company.Name)
	fmt.Println(emp.Person.Name)
}
