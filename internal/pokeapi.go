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
