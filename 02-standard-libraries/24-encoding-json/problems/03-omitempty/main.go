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

- Add JSON tags:
    Name → "name"
    Age  → "age"

- Add `omitempty` to the Age field.
- Create the User using the function arguments.
- Use json.Marshal().
- Return the JSON bytes.
- Return the error if conversion fails.
- In main(), call the function.
- Print the JSON as a string.


Example 1:


Input:

name = "Mahdi"
age = 27


Expected output:

{"name":"Mahdi","age":27}


Example 2:


Input:

name = "Mahdi"
age = 0


Expected output:

{"name":"Mahdi"}


Why?

Because Age has the `omitempty` tag and
0 is the zero value of int.
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age,omitempty"`
}

func UserToJSON(name string, age int) ([]byte, error) {
	user := User{
		Name: name,
		Age:  age,
	}
	return json.Marshal(user)
}

func main() {
	res, err := UserToJSON("Mahdi", 0)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(res)
	fmt.Println(string(res))
}
