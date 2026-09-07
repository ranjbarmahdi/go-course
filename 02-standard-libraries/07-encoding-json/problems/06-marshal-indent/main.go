/*
Problem:

Create a function:

func UserToPrettyJSON(name string, age int) ([]byte, error)


The function creates a User value and converts it
into formatted JSON using json.MarshalIndent().


Requirements:

- Create a User struct with:
    Name string
    Age  int

- Add JSON struct tags:
    Name → "name"
    Age  → "age"

- Create the User using the function arguments.
- Use json.MarshalIndent().
- Use:
    ""      as the prefix
    "    "  as the indent

- Return the JSON bytes.
- Return the error if marshaling fails.
- In main(), call the function.
- Print the JSON as a string.


Input:

name = "Mahdi"
age = 27


Expected output:

{
    "name": "Mahdi",
    "age": 27
}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func UserToPrettyJSON(name string, age int) ([]byte, error) {
	user := User{
		Name: name,
		Age:  age,
	}

	return json.MarshalIndent(user, "", "    ")
}

func main() {

	res, err := UserToPrettyJSON("Mahdi", 27)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(res)
	fmt.Println(string(res))
}
