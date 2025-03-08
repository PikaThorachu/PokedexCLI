package main

import (
	"fmt"

	PokeDex_API "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/Pokedex_API"
)

func commandMapb(cfg *PokeDex_API.Config, loc string) error {
	names, err := PokeDex_API.GetPreviousLocations(cfg)
	if err != nil {
		return err
	}
	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}
