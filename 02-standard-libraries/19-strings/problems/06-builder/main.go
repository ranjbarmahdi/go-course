/*
Problem:

Create a function:

func BuildMessage(names []string) string


The function receives a list of usernames and creates one formatted message.


Requirements:

- Use strings.Builder
- Add "Users: " at the beginning.
- Add each name separated by ", ".


Examples:


Input:

[]string{
    "Ali",
    "Reza",
    "Mahdi",
}


Output:

"Users: Ali, Reza, Mahdi"
*/

package main

import (
	"fmt"
	"strings"
)

func BuildMessage(names []string) string {
	var builder strings.Builder

	builder.WriteString("Users: ")
	builder.WriteString(strings.Join(names, ", "))

	return builder.String()
}

func main() {
	users := []string{
		"Ali",
		"Reza",
		"Mahdi",
	}

	fmt.Println(BuildMessage(users))
}
