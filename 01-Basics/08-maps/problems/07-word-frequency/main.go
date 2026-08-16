package main

import (
	"fmt"
	"strings"
)

func main() {
	text := "go is great and go is fast"

	words := strings.Split(text, " ")
	wordFreq := map[string]int{}

	for _, word := range words {
		if _, exists := wordFreq[word]; exists {
			wordFreq[word]++
		} else {
			wordFreq[word] = 1
		}
	}

	for key, value := range wordFreq {
		fmt.Printf("%s -> %d\n", key, value)
	}
}
