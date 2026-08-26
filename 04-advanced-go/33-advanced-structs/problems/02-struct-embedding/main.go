/*
Problem:

Create two structs:

type Person struct {
    Name string
    Age  int
}

type Employee struct {
    Person
    Salary int
}

Requirements:

- Embed Person inside Employee.
- Create an Employee with:

    Name:   "Mahdi"
    Age:    27
    Salary: 5000

- Access Name directly through Employee.
- Access Age directly through Employee.
- Access Salary directly through Employee.

Do NOT use:

    employee.Person.Name
    employee.Person.Age

For Name and Age, use field promotion.

Expected output:

Mahdi
27
5000
*/

package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type Employee struct {
	Person
	Salary int
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "Mahdi",
			Age:  27,
		},
		Salary: 25_000_000,
	}

	fmt.Println(emp.Name)
	fmt.Println(emp.Age)
	fmt.Println(emp.Salary)
}
