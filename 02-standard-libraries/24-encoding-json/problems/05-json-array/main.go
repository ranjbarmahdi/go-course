/*
Problem:

Create a function:

func JSONToUsers(data []byte) ([]User, error)


The function converts a JSON array into a slice of User
using json.Unmarshal().


Requirements:

- Create a User struct with:
    Name string
    Age  int

- Add JSON struct tags:
    Name → "name"
    Age  → "age"

- Create the JSONToUsers() function.
- Create a []User variable.
- Use json.Unmarshal().
- Remember that json.Unmarshal() needs a pointer
  to the destination slice.
- Return the users slice.
- Return the error if unmarshaling fails.
- In main(), create JSON containing multiple users.
- Call JSONToUsers().
- Print the users slice.
- Print each user's Name and Age.


Input:

[
    {
        "name": "Mahdi",
        "age": 27
    },
    {
        "name": "Ali",
        "age": 30
    },
    {
        "name": "Sara",
        "age": 22
    }
]


Expected output:

[{Mahdi 27} {Ali 30} {Sara 22}]
Mahdi 27
Ali 30
Sara 22
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

func JSONToUsers(data []byte) ([]User, error) {
	var users []User = []User{}
	err := json.Unmarshal(data, &users)
	return users, err
}

func main() {
	jsonData := `
	[
		{
			"name": "Mahdi",
			"age": 27
		},
		{
			"name": "Reza",
			"age": 20
		}
	]`

	users, err := JSONToUsers([]byte(jsonData))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(users)
	for _, v := range users {
		fmt.Printf("%s\t%d\n", v.Name, v.Age)
	}
}
