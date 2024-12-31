package main

import (
	"fmt"
)

func commandMap(config *Config, param string) error {

	// data, exists := config.cache.Get(*config.next)
	// if exists {

	// }

	locationsResponse, err := config.pokeapiClient.GetLocations(config.next)
	if err != nil {
		return fmt.Errorf("error getting locationsResponse: %v", err)
	}

	config.next = &locationsResponse.Next
	config.previous = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}

func commandMapB(config *Config, param string) error {
	if config.previous == nil {
		fmt.Println("you're on page one")
		return nil
	}

	locationsResponse, err := config.pokeapiClient.GetLocations(config.previous)
	if err != nil {
		return fmt.Errorf("error getting locationsResponse: %v", err)
	}

	config.next = &locationsResponse.Next
	config.previous = locationsResponse.Previous

	for _, location := range locationsResponse.Results {
		fmt.Println(location.Name)
	}

	return nil
}
