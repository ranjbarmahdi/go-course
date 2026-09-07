/*
Problem:

Create a function:

func CleanUsername(username string) string


The function removes unnecessary spaces from the beginning and end of a username.


Requirements:

- Use strings.TrimSpace()
- Return the cleaned username.


Examples:


Input:

"   mahdi   "


Output:

"mahdi"



Input:

"  golang  "


Output:

"golang"
*/

package main

import (
	"fmt"
	"strings"
)

func CleanUsername(username string) string {
	fields := strings.Fields(username)
	return strings.Join(fields, "")
}

func main() {
	username := "   mah  di   "
	fmt.Println(CleanUsername(username))
}
