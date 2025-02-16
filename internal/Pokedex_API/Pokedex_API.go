package PokeDex_API

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/PikaThorachu/Pokedex/PokedexCLI/internal/PokeCache"
)

// Declare config struct
type Config struct {
	Initial  string
	Next     string
	Previous string
	Cache    *PokeCache.Cache
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

func GetNextLocations(cfg *Config) ([]string, error) {
	var url string
	// Determine the URL: is it the first call or subsequent request?
	if cfg.Next == "" {
		url = "https://pokeapi.co/api/v2/location-area/"
	} else {
		url = cfg.Next
	}

	// Check if the data is already in the cache, return data if it is
	if val, ok := cfg.Cache.Get(url); ok {
		var names []string
		if err := json.Unmarshal(val, &names); err != nil {
			return nil, err
		}
		return names, nil
	}
	// If the data is not in the cache, make the API call
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

	// Update the config with the next and previous URLs from this response
	cfg.Next = response.Next
	cfg.Previous = response.Previous

	names := []string{}
	for _, loc := range response.Results {
		names = append(names, loc.Name)
	}
	// Store the data in the cache
	cachedValue, err := json.Marshal(names)
	if err != nil {
		return nil, err
	}
	cfg.Cache.Add(url, cachedValue)

	// Return the data
	return names, nil
}

func GetPreviousLocations(cfg *Config) ([]string, error) {
	var url string // IF no URL is stored yet (first call), use the base URL
	if cfg.Previous == "" {
		return nil, fmt.Errorf("you're on the first page")
	} else {
		url = cfg.Previous
	}

	// Check if the data is already in the cache, return data if it is
	if val, ok := cfg.Cache.Get(url); ok {
		var names []string
		if err := json.Unmarshal(val, &names); err != nil {
			return nil, err
		}
		return names, nil
	}

	// If the data is not in the cache, make the API call
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

	// Update the config with the next and previous URLs from this response
	cfg.Next = response.Next
	cfg.Previous = response.Previous

	names := []string{}
	for _, loc := range response.Results {
		names = append(names, loc.Name)
	}

	// Store the data in the cache
	cachedValue, err := json.Marshal(names)
	if err != nil {
		return nil, err
	}
	cfg.Cache.Add(url, cachedValue)

	// Return the data
	return names, nil
}
