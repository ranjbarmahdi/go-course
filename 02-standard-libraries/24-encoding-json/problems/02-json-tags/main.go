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

- Use JSON struct tags.
- The Name field must become "user_name" in JSON.
- The Age field must become "user_age" in JSON.
- Create the User using the function arguments.
- Use json.Marshal().
- Return the JSON bytes.
- Return the error if conversion fails.
- In main(), call the function.
- Print the JSON as a string.


Example:


Input:

name = "Mahdi"
age = 27


Expected JSON:

{"user_name":"Mahdi","user_age":27}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"user_name"`
	Age  int    `json:"user_age"`
}

func UserToJSON(name string, age int) ([]byte, error) {
	user := User{
		Name: name,
		Age:  age,
	}
	return json.Marshal(user)
}

func main() {
	res, err := UserToJSON("Mahdi", 27)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)
	fmt.Println(string(res))
}
