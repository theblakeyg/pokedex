package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		commandName := text[0]

		if command, ok := getCommands()[commandName]; ok {
			command.callback()
		} else {
			fmt.Print("Unknown Command\n")
		}
	}
}

func cleanInput(text string) []string {
	//make lowercase and split on whitespace
	text = strings.ToLower(text)
	words := strings.Fields(text)

	return words
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Display location names",
			callback:    commandMap,
		},
	}
}
