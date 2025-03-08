package PokeDex_API

import (
	"encoding/json"
	"fmt"
)

func InspectPokemon(cfg *Config, pokemon string) (PokemonResponse, error) {
	var response PokemonResponse // Define the struct to unmarshal json into here to prevent scope errors in the return line.
	// Step 2: Pull & Unmarshal json data from cache
	url := "https://pokeapi.co/api/v2/pokemon/" + pokemon
	if val, ok := cfg.Cache.Get(url); ok {
		if len(val) == 0 { // Validate empty data
			return PokemonResponse{}, fmt.Errorf("cached data is empty for URL: %s", url)
		}
		// Unmarshal the cached JSON data into a PokemonResponse struct
		if err := json.Unmarshal(val, &response); err != nil {
			return PokemonResponse{}, fmt.Errorf("failed to decode cached data: %w", err)
		}
	}

	// Step 3: Return PokemonResponse{}, nil error
	return response, nil
}
