package main

import (
	"fmt"
)

func commandExplore(config *Config, param string) error {

	pokemonResponse, err := config.pokeapiClient.GetPokemonByLocation(param)
	if err != nil {
		return fmt.Errorf("error getting locationsResponse: %v", err)
	}

	for _, encounter := range pokemonResponse.PokemonEncounters {
		fmt.Println(encounter.Pokemon.Name)
	}

	return nil
}
