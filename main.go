package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		fmt.Println("Your command was:", text[0])
	}

	//text, _ := scanner.('\n')

	// fmt.Println("Enter text: ")
	// text2 := ""
	// fmt.Scanln(text2)
	// fmt.Println(text2)

	// ln := ""
	// fmt.Sscanln("%v", ln)
	// fmt.Println(ln)
}

func cleanInput(text string) []string {
	//split on whitespace
	text = strings.ToLower(text)
	words := strings.Fields(text)

	return words
}
