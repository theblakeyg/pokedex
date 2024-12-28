package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type location struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

type locationsResponse struct {
	Count    int        `json:"count"`
	Next     string     `json:"next"`
	Previous *string    `json:"previous"`
	Results  []location `json:"results"`
}

func commandMap() error {
	//if next is available....we actually use that url
	defaultUrl := "https://pokeapi.co/api/v2/location-area/"
	res, err := http.Get(defaultUrl)

	if err != nil {
		errorText := fmt.Errorf("error getting locations: %v", err)
		fmt.Println(errorText)
		return errorText
	}
	defer res.Body.Close()

	var locations locationsResponse

	decoder := json.NewDecoder(res.Body)
	if err = decoder.Decode(&locations); err != nil {
		errorText := fmt.Errorf("error decoding response body: %v", err)
		fmt.Println(errorText)
		return errorText
	}

	for _, location := range locations.Results {
		fmt.Println(location.Name)
	}

	return nil
}
