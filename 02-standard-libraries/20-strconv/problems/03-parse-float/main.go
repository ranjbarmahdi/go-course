/*
Problem:

Create a function:

func StringToFloat(value string) (float64, error)


The function converts a string number into a float64.


Requirements:

- Use strconv.ParseFloat()
- Convert with 64-bit precision.
- Return the float value.
- Return the error if conversion fails.


Examples:


Input:

value = "19.99"


Output:

19.99
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

func StringToFloat(value string) (float64, error) {
	return strconv.ParseFloat(value, 64)
}

func main() {
	res1, err1 := StringToFloat("10")
	fmt.Println(res1, err1)

	res2, err2 := StringToFloat("10.5")
	fmt.Println(res2, err2)

	res3, err3 := StringToFloat("-10.501234")
	fmt.Println(res3, err3)

	res4, err4 := StringToFloat("a")
	fmt.Println(res4, err4)

}
