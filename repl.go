package main

import (
	"strings"

	PokeDex_API "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/Pokedex_API"
)

func cleanInput(text string) []string {
	splitwords := strings.Fields(text)
	for i := range splitwords {
		splitwords[i] = strings.ToLower(splitwords[i])
	}
	return splitwords
}

type cliCommand struct {
	name        string
	description string
	callback    func(*PokeDex_API.Config, string) error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Displays the map",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays previous 20 map items",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore <location>",
			description: "Explore the location",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch <pokemon>",
			description: "Catch the pokemon",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect <pokemon>",
			description: "Inspect the pokemon",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Display all pokemon in pokedex",
			callback:    commandPokedex,
		},
	}
}
