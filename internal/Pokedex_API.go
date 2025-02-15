package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Declare config struct
type Config struct {
	Initial  string
	Next     string
	Previous string
}

// Declare API struct
type LocationAreaResponse struct {
	Count    int    `json:"count"`
	Next     string `json:"next"`
	Previous string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func GetNextLocations(c *Config) ([]string, error) {
	var url string // IF no URL is stored yet (first call), use the base URL
	if c.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = c.Next
	}

	response := LocationAreaResponse{}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	c.Next = response.Next
	c.Previous = response.Previous

	names := []string{}
	for _, loc := range response.Results {
		names = append(names, loc.Name)
	}

	return names, nil
}

func GetPreviousLocations(c *Config) ([]string, error) {
	var url string // IF no URL is stored yet (first call), use the base URL
	if c.Previous == "" {
		return nil, fmt.Errorf("you're on the first page")
	} else {
		url = c.Previous
	}

	response := LocationAreaResponse{}

	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return nil, err
	}

	c.Next = response.Next
	c.Previous = response.Previous

	names := []string{}
	for _, loc := range response.Results {
		names = append(names, loc.Name)
	}

	return names, nil
}
