package main

import (
	"fmt"

	PokeDex_API "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/Pokedex_API"
)

func commandExplore(cfg *PokeDex_API.Config, loc string) error {
	pokemonNames, err := PokeDex_API.GetPokemon(cfg, loc)
	if err != nil {
		return err
	}
	fmt.Printf("Exploring %s...\n", loc)
	fmt.Println("Found Pokemon:")
	for _, pokemonName := range pokemonNames {
		fmt.Println("-" + pokemonName)
	}
	return nil
}
