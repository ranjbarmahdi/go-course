package main

import (
	"bufio"
	"fmt"
	"strings"
)

// Used for reading data token by token.

func main() {
	fmt.Println("1. bufio.Scanner ⭐⭐⭐")

	text := `
Hello Go
Backend Developer
Learning bufio
	`

	reader := strings.NewReader(text)
	scanner := bufio.NewScanner(reader)

	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}

	err := scanner.Err()

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("2. bufio.Reader ⭐⭐⭐")
	text2 := `
		Hello Go
		Backend Developer
		Learning bufio
	`

	reader2 := strings.NewReader(text2)
	bufferedReader := bufio.NewReader(reader2)
	fmt.Println(bufferedReader)
}
