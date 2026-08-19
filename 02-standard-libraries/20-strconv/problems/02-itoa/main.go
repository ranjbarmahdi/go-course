/*
Problem:

Create a function:

func IntToString(number int) string


The function converts an integer into a string.


Requirements:

- Use strconv.Itoa()
- Return the converted string.


Examples:


Input:

number = 123


Output:

"123"



Input:

number = -50


Output:

"-50"
*/

package main

import (
	"fmt"
	"strconv"
)

func IntToString(number int) string {
	return strconv.Itoa(number)
}

func main() {
	fmt.Println(IntToString(10))
	fmt.Println(IntToString(-10))
}
