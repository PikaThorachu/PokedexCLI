package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"slices"
	"strings"
	"time"

	PokeCache "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/PokeCache"
	PokeDex_API "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/Pokedex_API"
)

func cleanInput(text string) []string {
	splitwords := strings.Fields(text)
	for i := range splitwords {
		splitwords[i] = strings.ToLower(splitwords[i])
	}
	return splitwords
}

func commandExit(cfg *PokeDex_API.Config, loc string) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *PokeDex_API.Config, loc string) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for command := range commands {
		fmt.Printf("%s: %s\n", commands[command].name, commands[command].description)
	}
	return nil
}

func commandMap(cfg *PokeDex_API.Config, loc string) error {
	names, err := PokeDex_API.GetNextLocations(cfg)
	if err != nil {
		return err
	}
	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}

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
	if rand.Intn(400) > exp {
		fmt.Printf("%s escaped!\n", pokemon)
	} else {
		cfg.Pokedex = append(cfg.Pokedex, pokemon) //Add Pokemon to the Pokedex upon success
		fmt.Printf("%s was caught!\n", pokemon)
	}
	return nil
}

func commandInspect(cfg *PokeDex_API.Config, pokemon string) error {
	// Step 1: Check Pokedex for selected pokemon
	if !slices.Contains(cfg.Pokedex, pokemon) {
		return fmt.Errorf("you have not caught %s", pokemon) // If not in Pokedex, return empty PR struct & error message
	} else {
		response, err := PokeDex_API.InspectPokemon(cfg, pokemon)
		if err != nil {
			return err
		} else {
			fmt.Printf("Name:  %s\n", pokemon)
			fmt.Printf("Height:  %d\n", response.Weight)
			fmt.Printf("Weight:  %d\n", response.Weight)
			fmt.Printf("Stats:\n")
			for _, stat := range response.Stats {
				fmt.Printf("\t-%s: %v\n", stat.Name, stat.BaseStat)
			}
			fmt.Printf("Types:\n")
			for _, typeInfo := range response.Types {
				fmt.Println("  -", typeInfo.Type.Name)
			}
		}
		return nil
	}
}

func commandPokedex(cfg *PokeDex_API.Config, loc string) error {
	for _, pokemon := range cfg.Pokedex {
		fmt.Println(pokemon)
	}
	return nil
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

func main() {
	cfg := &PokeDex_API.Config{
		Cache: PokeCache.NewCache(10 * time.Second),
	} //initialize this once!
	fmt.Printf("Initial config: %+v\n", cfg)
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		if len(words) == 0 {
			continue
		}

		command := words[0]
		argument := ""
		if len(words) > 1 {
			argument = words[1]
		}
		if cmd, ok := commands[command]; ok {
			err := cmd.callback(cfg, argument)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
