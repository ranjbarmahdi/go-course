/*
Problem:

Create a function:

func JSONToUserStrict(data []byte) (User, error)


The function converts JSON data into a User struct
using json.NewDecoder() with DisallowUnknownFields().


Requirements:

- Create a User struct with:
    Name string
    Age  int

- Add JSON tags:
    Name → "name"
    Age  → "age"

- Create JSONToUserStrict().

- The function must:
    - Create a strings.Reader from data.
    - Create a json.Decoder.
    - Call DisallowUnknownFields().
    - Decode the JSON into User.
    - Return the User.
    - Return the error.

- In main(), create JSON containing:
    name
    age
    city

- The User struct does NOT contain a City field.

- Call JSONToUserStrict().

- Print the User.

- Print the error.


Important:

Normally:

json.Unmarshal()

ignores unknown fields.

But:

decoder.DisallowUnknownFields()

causes decoding to fail when the JSON contains
a field that does not exist in the destination struct.


Input:

{
    "name": "Mahdi",
    "age": 27,
    "city": "Baku"
}


Expected output:

{Mahdi 27}

Error:

json: unknown field "city"
*/

package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func JSONToUserStrict(data []byte) (User, error) {
	user := User{}
	reader := strings.NewReader(string(data))
	decoder := json.NewDecoder(reader)
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&user)
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

	if res, err := JSONToUserStrict([]byte(jsonData)); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
	}
}
