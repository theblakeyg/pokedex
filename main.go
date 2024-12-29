package main

import (
	"time"

	pokeapi "github.com/theblakeyg/pokedex/internal"
)

func main() {
	config := &Config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
	}
	runPokedex(config)
}
