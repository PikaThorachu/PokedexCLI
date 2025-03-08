package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	PokeCache "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/PokeCache"
	PokeDex_API "github.com/PikaThorachu/Pokedex/PokedexCLI/internal/Pokedex_API"
)

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
