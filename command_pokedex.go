package main

import "fmt"

func commandPokedex(config *Config, param string) error {
	if len(config.pokemon) == 0 {
		fmt.Println("You haven't caught any Pokemon yet!")
		return nil
	}

	for pokemonName, _ := range config.pokemon {
		fmt.Printf("  - %v\n", pokemonName)
	}

	return nil
}
