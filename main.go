package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func cleanInput(text string) []string {
	splitwords := strings.Fields(text)
	for i := range splitwords {
		splitwords[i] = strings.ToLower(splitwords[i])
	}
	return splitwords
}

func main() {
	bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		text := scanner.Text()
		words := cleanInput(text)
		fmt.Printf("Your command was: %s\n", words[0])
	}
}
