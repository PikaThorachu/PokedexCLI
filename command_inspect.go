package main

import (
	"fmt"
	"slices"

	PokeDex_API "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/Pokedex_API"
)

func commandInspect(cfg *PokeDex_API.Config, pokemon string) error {
	// Step 1: Check Pokedex for selected pokemon
	if !slices.Contains(cfg.Pokedex, pokemon) {
		return fmt.Errorf("you have not caught %s", pokemon) // If not in Pokedex, error message
	} else {
		response, err := PokeDex_API.InspectPokemon(cfg, pokemon)
		if err != nil {
			return err
		} else {
			fmt.Printf("Name:  %s\n", pokemon)
			fmt.Printf("Height:  %d\n", response.Height)
			fmt.Printf("Weight:  %d\n", response.Weight)
			fmt.Printf("Stats:\n")
			for _, stat := range response.Stats {
				fmt.Printf("  - %s: %v\n", stat.Stat.Name, stat.BaseStat)
			}
			fmt.Printf("Types:\n")
			for _, typeInfo := range response.Types {
				fmt.Println("  - ", typeInfo.Type.Name)
			}
		}
		return nil
	}
}
