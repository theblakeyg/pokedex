package main

import (
	"time"

	pokeapi "github.com/theblakeyg/pokedex/internal/pokeapi"
	"github.com/theblakeyg/pokedex/internal/pokecache"
)

func main() {
	config := &Config{
		pokeapiClient: pokeapi.NewClient(5 * time.Second),
		cache:         pokecache.NewCache(5 * time.Second),
		pokemon:       make(map[string]pokeapi.PokemonResponse),
	}
	runPokedex(config)
}
