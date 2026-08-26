/*
Problem:

Create a function:

func JSONToUser(data []byte) (User, error)


The function converts JSON data into a User struct
using json.Unmarshal().


Requirements:

- Create a User struct with:
    Name *string
    Age  *int

- Add JSON tags:
    Name → "name"
    Age  → "age"

- Create JSONToUser().

- Use json.Unmarshal().

- Remember that json.Unmarshal() needs a pointer
  to the destination User.

- Return the User.
- Return the error.

- In main(), create this JSON:

{
    "name": null,
    "age": null
}

- Call JSONToUser().

- Print the User.
- Print User.Name.
- Print User.Age.

Important:

JSON null can be decoded into pointer fields
as nil.

For example:

"name": null

becomes:

Name == nil

And:

"age": null

becomes:

Age == nil


Expected output:

{<nil> <nil>}
<nil>
<nil>
*/

package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name *string `json:"name"`
	Age  *int    `json:"age"`
}

func JSONToUser(data []byte) (User, error) {
	user := User{}
	err := json.Unmarshal(data, &user)
	return user, err
}

func main() {
	jsonString :=
		`
		{
			"name": null,
			"age": null
		}
	`
	if res, err := JSONToUser([]byte(jsonString)); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println(res)
		fmt.Println(res.Name)
		fmt.Println(res.Age)
	}
}
