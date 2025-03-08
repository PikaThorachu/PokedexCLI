package PokeDex_API

import (
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
