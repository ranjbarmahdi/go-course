package main

import (
	"fmt"
	"strings"
)

// ============================================================
// strings package
// ============================================================
//
// Common functions:
//
// strings.Contains()
// strings.Split()
// strings.Join()
// strings.Replace()
// strings.ReplaceAll()
// strings.TrimSpace()
// strings.Fields()
// strings.Builder
//

func main() {

	// ============================================================
	// 1. strings.Contains()
	// ============================================================
	// Checks whether a substring exists inside a string.
	//
	// Returns:
	// true  -> substring exists
	// false -> substring does not exist

	mainString := "Hi, i am mahdi ranjbar."
	searchString := "mahdi"

	fmt.Println(strings.Contains(mainString, searchString))

	// ============================================================
	// 2. strings.Split()
	// ============================================================
	// Splits a string into a []string using a separator.
	//
	// Example:
	//
	// "Hello World"
	//      ↓
	// ["Hello", "World"]

	test := "Welcome Mahdi To This Course"
	separator := " "

	words := strings.Split(test, separator)

	fmt.Println(words)

	// ============================================================
	// 3. strings.Join()
	// ============================================================
	// Combines a []string into one string.
	//
	// The second argument is inserted between elements.

	words2 := []string{
		"Welcome",
		"Mahdi",
		"To",
		"This",
		"Course",
	}

	sentence := strings.Join(words2, " ")

	fmt.Println(sentence)

	// ============================================================
	// 4. strings.Replace()
	// ============================================================
	// Replaces part of a string.
	//
	// The fourth argument specifies the maximum number
	// of replacements.
	//
	// 1  -> replace first occurrence
	// 2  -> replace first two occurrences
	// -1 -> replace all occurrences

	message := "hello java java"

	result := strings.Replace(
		message,
		"java",
		"golang",
		1,
	)

	fmt.Println(result)

	// ============================================================
	// 5. strings.ReplaceAll()
	// ============================================================
	// Replaces ALL occurrences.
	//
	// It is equivalent to:
	//
	// strings.Replace(text, old, new, -1)

	result2 := strings.ReplaceAll(
		message,
		"java",
		"golang",
	)

	fmt.Println(result2)

	// ============================================================
	// 6. strings.TrimSpace()
	// ============================================================
	// Removes whitespace from the beginning and end
	// of a string.
	//
	// It does NOT remove spaces between words.

	name := "   Mahdi   "

	clean := strings.TrimSpace(name)

	fmt.Println(clean)

	// ============================================================
	// 7. strings.Fields()
	// ============================================================
	// Splits a string into words.
	//
	// Unlike strings.Split(), Fields automatically handles
	// multiple spaces and other whitespace characters.

	sentence4 := "   Mahdi    Ranjbar Hello    World     "

	fields := strings.Fields(sentence4)

	fmt.Println(fields)

	// Example:
	//
	// "  Mahdi    Ranjbar  "
	//          ↓
	// ["Mahdi", "Ranjbar"]

	// ============================================================
	// 8. strings.Builder
	// ============================================================
	// Builder is used to efficiently build a string.
	//
	// Useful when repeatedly adding strings.

	var builder strings.Builder

	builder.WriteString("Hello ")
	builder.WriteString("World")

	result3 := builder.String()

	fmt.Println(result3)

	// ============================================================
	// 9. strings.Builder with multiple values
	// ============================================================

	var builder2 strings.Builder

	builder2.WriteString("Go")
	builder2.WriteString(" ")
	builder2.WriteString("Backend")
	builder2.WriteString(" ")
	builder2.WriteString("Developer")

	fmt.Println(builder2.String())
}
