/*
Problem:

Create a function:

func SplitWords(sentence string) []string

The function receives a sentence and returns each word as an element of a slice.

Requirements:

- Use strings.Split()
- The separator should be a space (" ")
- Return []string


Examples:

Input:

"golang backend developer"

Output:

[
"golang",
"backend",
"developer"
]

Input:

"hello world"

Output:

[
"hello",
"world"
]
*/

package main

import (
	"fmt"
	"strings"
)

func SplitWords(sentence string) []string {
	return strings.Split(sentence, " ")
}

func main() {
	fmt.Println(SplitWords("golang     backend developer      "))
}
