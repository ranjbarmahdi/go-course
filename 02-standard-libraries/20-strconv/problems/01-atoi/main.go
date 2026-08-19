/*
Problem:

Create a function:

func StringToInt(value string) (int, error)


The function converts a string number into an integer.


Requirements:

- Use strconv.Atoi()
- Return the integer value.
- Return the error if conversion fails.


Examples:


Input:

value = "123"


Output:

123
nil



Input:

value = "abc"


Output:

0
error
*/

package main

import (
	"fmt"
	"strconv"
)

func StringToInt(value string) (int, error) {
	return strconv.Atoi(value)
}

func main() {

	number, err := StringToInt("123")

	fmt.Println(number)
	fmt.Println(err)

}
