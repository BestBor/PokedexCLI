package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/BestBor/pokedexcli/pokeapi"
)

var cliCommands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func(*config) error
}

type config struct {
	Next     *string
	Previous *string
}

func commandExit(cfg *config) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp(cfg *config) error {
	fmt.Println(`Welcome to the Pokedex!
Usage:
	`)
	for _, cmd := range cliCommands {
		fmt.Printf("  %-10s %s\n", cmd.name, cmd.description)
	}
	return nil
}

func commandMap(cfg *config) error {

	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Next != nil && *cfg.Next != "" {
		url = *cfg.Next
	}

	locations, err := pokeapi.Fetch(url)
	if err != nil {
		return fmt.Errorf("error getting the locations: %w", err)
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}

func commandMapb(cfg *config) error {

	url := "https://pokeapi.co/api/v2/location-area/"
	if cfg.Previous != nil && *cfg.Previous != "" {
		url = *cfg.Previous
	} else {
		fmt.Println("you're on the first page")
		return nil
	}

	locations, err := pokeapi.Fetch(url)
	if err != nil {
		return fmt.Errorf("error getting the locations: %w", err)
	}

	for _, loc := range locations.Results {
		fmt.Println(loc.Name)
	}

	cfg.Next = locations.Next
	cfg.Previous = locations.Previous

	return nil
}

func initCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"help": {
			name:        "help",
			description: "Describes how to use the REPL",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    commandMapb,
		},
	}
}

func main() {

	cfg := &config{}

	reader := bufio.NewScanner(os.Stdin)
	cliCommands = initCommands()
	for {
		fmt.Print("Pokedex > ")
		if !reader.Scan() {
			break
		}
		input := strings.TrimSpace(reader.Text())
		components := strings.Fields(input)
		command := strings.ToLower(components[0])

		if cmd, ok := cliCommands[command]; ok {
			if err := cmd.callback(cfg); err != nil {
				fmt.Println("Error:", err)
			}
		} else {
			fmt.Println("Unknown command:", command)
		}

	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)

	return words
}
