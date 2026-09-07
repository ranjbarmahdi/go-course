/*
Problem:

Create a function:

func FloatToString(value float64) string


The function converts a float64 number into a string.


Requirements:

- Use strconv.FormatFloat()
- Use format 'f'
- Keep 2 digits after decimal point.
- Use 64-bit precision.
- Return the converted string.


Examples:


Input:

value = 19.99


Output:

"19.99"



Input:

value = 10.5


Output:

"10.50"



Input:

value = 3.14159


Output:

"3.14"
*/

package main

import (
	"fmt"
	"strconv"
)

func FloatToString(value float64) string {
	return strconv.FormatFloat(value, 'f', 4, 64)
}

func main() {
	fmt.Println(FloatToString(12.51234))
	fmt.Println(FloatToString(-12.51234))
}
