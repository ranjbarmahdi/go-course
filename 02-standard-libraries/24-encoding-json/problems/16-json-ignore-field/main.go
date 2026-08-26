/*
Problem:

Create a function:

func UserToJSON(user User) ([]byte, error)


The function converts a User struct into JSON
using json.Marshal().


Requirements:

- Create a User struct with:

    Name     string
    Age      int
    Password string

- Add JSON tags:

    Name     → "name"
    Age      → "age"
    Password → "-"

- Create UserToJSON().

- Use json.Marshal().

- Return the JSON bytes.
- Return the error.

- In main():

    - Create a User:

        Name: "Mahdi"
        Age:  27
        Password: "123456"

    - Call UserToJSON().

    - Print the JSON string.


Important:

The JSON "-" tag tells the encoder:

"Do not include this field in JSON."


Example:

Go struct:

{
    Name: "Mahdi",
    Age: 27,
    Password: "123456"
}


JSON:

{
    "name": "Mahdi",
    "age": 27
}


Password should NOT appear.


Expected output:

{"name":"Mahdi","age":27}
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name     string `json:"name"`
	Age      int    `json:"age"`
	Password string `json:"-"`
}

func UserToJSON(user User) ([]byte, error) {
	return json.Marshal(user)
}

func main() {
	user := User{
		Name:     "Mahdi",
		Age:      27,
		Password: "12345",
	}

	if res, err := UserToJSON(user); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(string(res))
	}
}
