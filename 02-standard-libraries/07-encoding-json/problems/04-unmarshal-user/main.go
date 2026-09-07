/*
Problem:

Create a function:

func JSONToUser(data []byte) (User, error)


The function converts JSON data into a User value
using json.Unmarshal().


Requirements:

- Create a User struct with:
    Name string
    Age  int

- Add JSON struct tags:
    Name → "name"
    Age  → "age"

- Create the JSONToUser() function.
- Use json.Unmarshal().
- Remember that json.Unmarshal() needs a pointer
  to the destination value.
- Return the User value.
- Return the error if unmarshaling fails.
- In main(), create JSON data.
- Call JSONToUser().
- Print the User.
- Print the user's Name.
- Print the user's Age.


Example:


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
	userJson := `
	{
		"name": "Mahdi",
		"age": 27
	}
	`
	user, err := JSONToUser([]byte(userJson))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(user)
	fmt.Println(user.Name)
	fmt.Println(user.Age)
}
