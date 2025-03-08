package main

import (
	"fmt"

	PokeDex_API "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/Pokedex_API"
)

func commandPokedex(cfg *PokeDex_API.Config, loc string) error {
	for _, pokemon := range cfg.Pokedex {
		fmt.Println(" - ", pokemon)
	}
	return nil
}
