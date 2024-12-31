package main

import "fmt"

func commandInspect(config *Config, param string) error {
	pokemon, exists := config.pokemon[param]
	if !exists {
		fmt.Printf("You have not caught a %v yet!\n", param)
	}

	if exists {
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Println("Stats:")
		for _, stat := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", stat.Stat.Name, stat.BaseStat)
		}
		fmt.Println("Types:")
		for _, pokemonType := range pokemon.Types {
			fmt.Printf("  -%v\n", pokemonType.Type.Name)
		}
	}

	return nil
}
