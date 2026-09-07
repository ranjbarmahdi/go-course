/*
Problem:

Create a function:

func UserToJSON(name string, age int) ([]byte, error)


The function creates a User value and converts it
into JSON using json.Marshal().


Requirements:

- Create a User struct with:
    Name string
    Age  int

- Create the User using the function arguments.
- Use json.Marshal()
- Return the JSON bytes.
- Return the error if marshaling fails.
- In main(), call the function.
- Print the returned []byte.
- Print the JSON as a string.


Example:


Input:

name = "Mahdi"
age = 27


Output:

[123 34 78 ...]
{"Name":"Mahdi","Age":27}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string
	Age  int
}

func UserToJSON(name string, age int) ([]byte, error) {
	user := User{
		Name: name,
		Age:  age,
	}
	return json.Marshal(user)
}

func main() {
	result, err := UserToJSON("Mahdi", 22)
	if err != nil {
		fmt.Println("Error: ", err)
		panic(err)
	}
	fmt.Println("Json Of User (bytes): ", result)
	fmt.Println("Json Of User: ", string(result))
}
