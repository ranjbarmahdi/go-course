/*
Problem:

Create a function:

func WriteUserJSON(name string, age int) error


The function creates a User value and writes it
as JSON using json.NewEncoder().


Requirements:

- Create a User struct with:
    Name string
    Age  int

- Add JSON struct tags:
    Name → "name"
    Age  → "age"

- Create the User using the function arguments.
- Create an Encoder using:
    json.NewEncoder()

- Use os.Stdout as the io.Writer.

- Encode the User using:
    encoder.Encode()

- Return the error if encoding fails.
- In main(), call the function.
- Handle the returned error.


Expected output:

{"name":"Mahdi","age":27}
*/

package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

type User struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func WriteUserJson(name string, age int, writer io.Writer) error {
	user := User{
		Name: name,
		Age:  age,
	}

	encoder := json.NewEncoder(writer)

	if err := encoder.Encode(user); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := WriteUserJson("Mahdi", 27, os.Stdout); err != nil {
		fmt.Println("Error:", err)
	}

	// more train
	var buf bytes.Buffer

	writer := io.Writer(&buf)

	if err := WriteUserJson("Mahdi", 27, writer); err != nil {
		fmt.Println("Error:", err)
	}
	fmt.Println(buf.String())
}
