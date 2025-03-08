package main

import (
	"fmt"
	"math/rand"

	PokeDex_API "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/Pokedex_API"
)

func commandCatch(cfg *PokeDex_API.Config, pokemon string) error {
	// Check if the Pokémon is already in the Pokedex
	for _, p := range cfg.Pokedex {
		if p == pokemon {
			fmt.Printf("You already caught %s!\n", pokemon)
			return nil
		}
	}

	// Fetch data about the Pokemon from cache or API
	exp, err := PokeDex_API.CatchPokemon(cfg, pokemon)
	if err != nil {
		return err
	}

	// Simulate throwing the pokeball
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon)
	if rand.Intn(200) > exp {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		cfg.Pokedex = append(cfg.Pokedex, pokemon) //Add Pokemon to the Pokedex upon success
		fmt.Printf("%s was caught!\n", pokemon)
	}
	return nil
}
