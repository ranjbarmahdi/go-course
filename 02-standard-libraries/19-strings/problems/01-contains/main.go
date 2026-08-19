/*
Problem:

Create a function:

func ContainsWord(text string, word string) bool

The function checks if "word" exists inside "text".

Requirements:
- Use strings.Contains()
- Return true when the word exists.
- Return false otherwise.

Examples:

Input:
text = "golang backend developer"
word = "backend"

Output:
true


Input:
text = "golang backend developer"
word = "python"

Output:
false
*/

package main

import (
	"fmt"
	"strings"
)

func ContainsWord(text string, word string) bool {
	return strings.Contains(text, word)
}

func main() {
	text := "    golang              backend  developer     "
	word := "backend"
	fmt.Println(ContainsWord((text), strings.TrimSpace(word)))
}
