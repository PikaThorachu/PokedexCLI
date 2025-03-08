package PokeDex_API

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	PokeCache "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/PokeCache"
)

// Declare config struct
type Config struct {
	Initial  string
	Next     string
	Previous string
	Pokedex  []string
	Cache    *PokeCache.Cache
}

// Declare API struct
type LocationAreaResponse struct {
	Next     string
	Previous string
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
	} `json:"pokemon_encounters"`
}

type PokemonResponse struct {
	Name           string `json:"name"`
	BaseExperience int    `json:"base_experience"`
	Height         int    `json:"height"`
	Weight         int    `json:"weight"`
	PokemonStats
}

type PokemonStats struct {
	Stats []struct {
		BaseStat int `json:"base_stat"`
		Effort   int `json:"effort"`
		Stat     struct {
			Name string `json:"name"`
		} `json:"stat"`
	} `json:"Stats"`
	Types []struct {
		Type struct {
			Name string `json:"name"`
		} `json:"type"`
	} `json:"types"`
}

type ThreadSafeCache struct {
	data map[string][]byte
	mu   sync.Mutex
}

func NewThreadSafeCache() *ThreadSafeCache {
	return &ThreadSafeCache{
		data: make(map[string][]byte),
	}
}

// Adds a new entry to the cache
func (c *ThreadSafeCache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Gets an entry from the cache if it exists
func (c *ThreadSafeCache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	return val, ok
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
