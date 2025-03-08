package PokeDex_API

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func GetPokemon(cfg *Config, loc string) ([]string, error) {
	// Add the location to the base URL
	url := fmt.Sprintf("https://pokeapi.co/api/v2/location-area/%s", loc)

	// Check if the data is already in the cache, return data if it is
	if val, ok := cfg.Cache.Get(url); ok {
		var pokemonName []string
		if err := json.Unmarshal(val, &pokemonName); err != nil {
			return nil, err
		}
		return pokemonName, nil
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

	pokemonNames := []string{}
	for _, pokemon := range response.PokemonEncounters {
		pokemonNames = append(pokemonNames, pokemon.Pokemon.Name)
	}

	// Store the data in the cache
	cachedValue, err := json.Marshal(pokemonNames)
	if err != nil {
		return nil, err
	}
	cfg.Cache.Add(url, cachedValue)

	// Return the data
	return pokemonNames, nil
}
