/*
Problem:

Create a function:

func JSONToUser(data []byte) (User, error)


The function converts JSON data into a User struct
using json.Unmarshal().


Requirements:

- Create a User struct with:
    Name string
    Age  int

- Add JSON tags:
    Name → "name"
    Age  → "age"

- Create JSONToUser().
- Use json.Unmarshal().
- Return the User and error.

- In main(), create JSON containing:
    name
    age
    city

- The User struct does NOT contain a City field.

- Unmarshal the JSON into User.

- Print the User.

Important:

By default, json.Unmarshal() ignores JSON fields
that do not exist in the destination struct.


Input:

{
    "name": "Mahdi",
    "age": 27,
    "city": "Baku"
}


Expected output:

{Mahdi 27}

The "city" field should be ignored.
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

func JSONToUser(data []byte) (User, error) {
	user := User{}
	err := json.Unmarshal(data, &user)
	return user, err
}

func main() {
	jsonData := `
		{
			"name": "Mahdi",
			"age": 27,
			"city": "Baku"
		}
	`

	if res, err := JSONToUser([]byte(jsonData)); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}
}
