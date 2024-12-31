package main

import (
	"fmt"
	"math/rand"
)

func commandCatch(config *Config, param string) error {

	pokemonResponse, err := config.pokeapiClient.GetPokemon(param)
	if err != nil {
		return fmt.Errorf("error getting locationsResponse: %v", err)
	}

	fmt.Printf("Throwing a Pokeball at %v...\n", param)

	experience := pokemonResponse.BaseExperience
	fmt.Printf("Experience: %v\n", experience)
	caught := rand.Intn(50) + experience

	if caught > 150 {
		fmt.Printf("Caught roll: %v\n", caught)
		fmt.Printf("%v escaped!\n", param)
	} else {
		fmt.Printf("%v was caught!\n", param)
		config.pokemon[param] = pokemonResponse
	}

	return nil
}
