/*
Problem:

Create:

type Person struct {
    Name string
}

Create a method:

func (p Person) SayHello()

It should print:

Hello from Person

Then create:

type Employee struct {
    Person
}

Create another SayHello() method on Employee.

It should print:

Hello from Employee

Requirements:

- Embed Person inside Employee.
- Create an Employee with Name = "Mahdi".
- Call:

    employee.SayHello()

- Then call:

    employee.Person.SayHello()

The first call should use Employee's method.

The second call should explicitly use Person's method.

Expected output:

Hello from Employee
Hello from Person
*/

package main

import "fmt"

type Person struct {
	Name string
}

func (p Person) SayHello() {
	fmt.Println("Hello from Person")
}

type Employee struct {
	Person
}

func (e Employee) SayHello() {
	fmt.Println("Hello, from Employee")
}

func main() {
	emp := Employee{
		Person: Person{
			Name: "Mahdi",
		},
	}
	emp.SayHello()
	emp.Person.SayHello()
}
