/*
Problem:

Create a function:

func JoinWords(words []string) string


The function receives a slice of words and combines them into one string.


Requirements:

- Use strings.Join()
- The separator should be a space (" ")
- Return a string


Examples:


Input:

[]string{
    "golang",
    "backend",
    "developer",
}


Output:

"golang backend developer"



Input:

[]string{
    "hello",
    "world",
}


Output:

"hello world"
*/

package main

import (
	"fmt"
	"strings"
)

func JoinWords(words []string) string {
	return strings.Join(words, " ")
}

func main() {
	words := []string{"Hello", "World"}
	fmt.Println(JoinWords(words))
}
