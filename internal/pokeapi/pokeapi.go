package pokeapi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Client struct {
	client  http.Client
	Timeout time.Duration
	baseUrl string
}

func NewClient(timeout time.Duration) Client {
	return Client{
		client:  http.Client{},
		Timeout: timeout,
		baseUrl: "https://pokeapi.co/api/v2",
	}
}

func (c *Client) GetLocations(pageURL *string) (locationsResponse, error) {
	url := c.baseUrl + "/location-area"
	if pageURL != nil {
		url = *pageURL
	}
	//create Request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationsResponse{}, fmt.Errorf("error creating request: %v", err)
	}

	//Client Do(Request)
	res, err := c.client.Do(request)

	if err != nil {
		errorText := fmt.Errorf("error getting locations: %v", err)
		fmt.Println(errorText)
		return locationsResponse{}, errorText
	}
	defer res.Body.Close()

	var locations locationsResponse

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&locations); err != nil {
		errorText := fmt.Errorf("error decoding response body: %v", err)
		fmt.Println(errorText)
		return locationsResponse{}, errorText
	}

	return locations, nil
}

func (c *Client) GetPokemonByLocation(location string) (locationDetailsResponse, error) {
	url := c.baseUrl + "/location-area/" + location

	//create Request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return locationDetailsResponse{}, fmt.Errorf("error creating request: %v", err)
	}

	//Client Do(Request)
	res, err := c.client.Do(request)

	if err != nil {
		errorText := fmt.Errorf("error getting locations: %v", err)
		fmt.Println(errorText)
		return locationDetailsResponse{}, errorText
	}
	defer res.Body.Close()

	var pokemon locationDetailsResponse

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&pokemon); err != nil {
		errorText := fmt.Errorf("error decoding response body: %v", err)
		fmt.Println(errorText)
		return locationDetailsResponse{}, errorText
	}

	return pokemon, nil
}

func (c *Client) GetPokemon(pokemon string) (PokemonResponse, error) {
	url := c.baseUrl + "/pokemon/" + pokemon

	//create Request
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return PokemonResponse{}, fmt.Errorf("error creating request: %v", err)
	}

	//Client Do(Request)
	res, err := c.client.Do(request)

	if err != nil {
		errorText := fmt.Errorf("error getting locations: %v", err)
		fmt.Println(errorText)
		return PokemonResponse{}, errorText
	}
	defer res.Body.Close()

	var pokemonDetails PokemonResponse

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&pokemonDetails); err != nil {
		errorText := fmt.Errorf("error decoding response body: %v", err)
		fmt.Println(errorText)
		return PokemonResponse{}, errorText
	}

	return pokemonDetails, nil
}
