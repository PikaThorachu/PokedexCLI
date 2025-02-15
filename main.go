package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/PikaThorachu/Pokedex/PokedexCLI/internal"
)

func cleanInput(text string) []string {
	splitwords := strings.Fields(text)
	for i := range splitwords {
		splitwords[i] = strings.ToLower(splitwords[i])
	}
	return splitwords
}

func commandExit(c *internal.Config) error {
	fmt.Printf("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(c *internal.Config) error {
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	for command := range commands {
		fmt.Printf("%s: %s\n", commands[command].name, commands[command].description)
	}
	return nil
}

func commandMap(c *internal.Config) error {
	names, err := internal.GetNextLocations(c)
	if err != nil {
		return err
	}
	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}

func commandMapb(c *internal.Config) error {
	names, err := internal.GetPreviousLocations(c)
	if err != nil {
		return err
	}
	for _, name := range names {
		fmt.Println(name)
	}
	return nil
}

type cliCommand struct {
	name        string
	description string
	callback    func(*internal.Config) error
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
	}
}

func main() {
	cfg := &internal.Config{} //initialize this once!
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
		if cmd, ok := commands[command]; ok {
			err := cmd.callback(cfg)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Unknown command")
		}
	}
}
