package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println("Hello, World!")
}

func cleanInput(text string) []string {
	//split on whitespace
	words := strings.Fields(text)

	//lowercase word and save back to slice
	for i, word := range words {
		words[i] = strings.ToLower(word)
	}

	return words
}
