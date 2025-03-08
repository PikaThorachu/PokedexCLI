package PokeDex_API

import (
	"encoding/json"
	"fmt"
	"net/http"
)

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

	// Check if HTTP response is successful (status code 200)
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %s", res.Status)
	}

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
