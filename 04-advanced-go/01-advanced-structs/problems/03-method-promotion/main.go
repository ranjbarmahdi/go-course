/*
Problem:

Create a struct:

type Person struct {
    Name string
}

Create a method:

func (p Person) SayHello()

The method should print:

Hello, Mahdi

Then create:

type Employee struct {
    Person
}

Requirements:

- Embed Person inside Employee.
- Create an Employee with Name = "Mahdi".
- Do NOT create a SayHello() method on Employee.
- Call SayHello() directly through Employee.
- This should work because the method is promoted
  from the embedded Person.

Expected output:

Hello, Mahdi
*/

package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) SayHello() {
	fmt.Printf("Hello, %s\n", p.Name)
}

type Employee struct {
	Person
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "Mahdi",
		},
	}
	emp.SayHello()
}
