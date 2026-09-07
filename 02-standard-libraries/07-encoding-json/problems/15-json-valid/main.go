/*
Problem:

Create a function:

func IsValidJSON(data []byte) bool


The function checks whether the given bytes
contain valid JSON.


Requirements:

- Create IsValidJSON().

- Use:

    json.Valid()

- Return the result.

- In main():

    - Create valid JSON.

    - Create invalid JSON.

    - Call IsValidJSON() for both.

    - Print the results.


Valid JSON:

{
    "name": "Mahdi",
    "age": 27
}


Invalid JSON:

{
    "name": "Mahdi",
    "age": 27,
}


Expected output:

Valid JSON: true

Invalid JSON: false


Important:

json.Valid() only checks if the JSON syntax
is correct.

It does NOT decode the JSON.

Example:

json.Unmarshal()

does:

JSON
 ↓
Go value


json.Valid()

does:

JSON
 ↓
true / false
*/

package main

import (
	"encoding/json"
	"fmt"
)

func IsValidJSON(data []byte) bool {
	return json.Valid(data)
}

func main() {
	validJSON := `
		{
			"name": "Mahdi",
			"age": 27
		}
		`

	invalidJSON := `
		{
			"name": "Mahdi",
			"age": 27,
		}
		`

	fmt.Println("Valid JSON:", IsValidJSON([]byte(validJSON)))
	fmt.Println("Invalid JSON:", IsValidJSON([]byte(invalidJSON)))
}
