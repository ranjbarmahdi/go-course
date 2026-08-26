/*
Problem:

Create a function:

func ReadUserJSON(reader io.Reader) (User, error)


The function reads JSON from an io.Reader
and converts it into a User using json.NewDecoder().


Requirements:

- Create a User struct with:
    Name string
    Age  int

- Add JSON struct tags:
    Name → "name"
    Age  → "age"

- Create the ReadUserJSON() function.
- Accept an io.Reader as the input.
- Create a decoder using:
    json.NewDecoder()

- Use decoder.Decode() to decode the JSON.
- Remember that Decode() needs a pointer
  to the destination User.

- Return the User.
- Return the error if decoding fails.

- In main():
    - Create JSON data using strings.NewReader().
    - Pass the reader to ReadUserJSON().
    - Print the User.
    - Print Name.
    - Print Age.


Input:

{
    "name": "Mahdi",
    "age": 27
}


Expected output:

{Mahdi 27}
Mahdi
27
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func ReadUserJson(reader io.Reader) (User, error) {
	user := User{}
	decoder := json.NewDecoder(reader)
	err := decoder.Decode(&user)
	return user, err
}

func main() {
	jsonData := `
	{
	   "name": "Mahdi",
	   "age": 27
	}
	`
	buf := bytes.NewBufferString(jsonData)
	reader := io.Reader(buf)

	if user, err := ReadUserJson(reader); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(user)
	}
}
