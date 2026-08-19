/*
Problem:

Create a function:

func ReplaceWord(text string, old string, new string) string


The function replaces a word inside a text.


Requirements:

- Use strings.Replace()
- Replace only the first occurrence.
- Return the updated string.


Examples:


Input:

text = "go is good, go is fast"
old = "go"
new = "golang"


Output:

"golang is good, go is fast"



Input:

text = "hello world"
old = "world"
new = "golang"


Output:

"hello golang"
*/

package main

import (
	"fmt"
	"strings"
)

func ReplaceWord(text string, old string, new string) string {
	return strings.Replace(text, old, new, 1)
}

func main() {
	text := "go is good, go is fast"
	text = ReplaceWord(text, "go", "golang")
	fmt.Println(text)
}
