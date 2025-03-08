package PokeDex_API

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func CatchPokemon(cfg *Config, pokemon string) (int, error) {
	url := fmt.Sprintf("https://pokeapi.co/api/v2/pokemon/%s", pokemon)

	// Check if the data is already in the cache, return data if it is
	if val, ok := cfg.Cache.Get(url); ok {
		var response PokemonResponse // Unmarshal into the correct struct
		if len(val) == 0 {           // Validate empty data
			return 0, fmt.Errorf("cached data is empty for URL: %s", url)
		}
		// Unmarshal the cached JSON data into a PokemonResponse struct
		if err := json.Unmarshal(val, &response); err != nil {
			return 0, fmt.Errorf("failed to decode cached data: %w", err)
		}
		// Return the BaseExperience value
		return response.BaseExperience, nil
	}

	// If the data is not in the cache, make the API call
	response := PokemonResponse{}

	res, err := http.Get(url)
	if err != nil {
		return 0, err
	}

	defer res.Body.Close()

	// Check if HTTP response is successful (status code 200)
	if res.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("API request failed with status: %s", res.Status)
	}

	decoder := json.NewDecoder(res.Body)
	if err := decoder.Decode(&response); err != nil {
		return 0, err
	}

	// Store the data in the cache
	responseData := PokemonResponse{
		Name:           pokemon,
		BaseExperience: response.BaseExperience,
		Height:         response.Height,
		Weight:         response.Weight,
		PokemonStats:   response.PokemonStats,
	}

	jsonData, err := json.Marshal(responseData)
	if err != nil {
		return 0, err
	}
	cfg.Cache.Add(url, jsonData)

	// Return the data
	return response.BaseExperience, nil
}
