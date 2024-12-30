package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	pokeapi "github.com/theblakeyg/pokedex/internal/pokeapi"
	"github.com/theblakeyg/pokedex/internal/pokecache"
)

type Config struct {
	cache         pokecache.Cache
	pokeapiClient pokeapi.Client
	previous      *string
	next          *string
}

func runPokedex(config *Config) {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := cleanInput(scanner.Text())
		commandName := text[0]

		if command, ok := getCommands()[commandName]; ok {
			command.callback(config)
		} else {
			fmt.Print("Unknown Command\n")
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*Config) error
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
		"mapb": {
			name:        "mapb",
			description: "Display previous location names",
			callback:    commandMapB,
		},
	}
}
