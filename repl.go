package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    callbackExit,
		},
		"help": {
			name:        "help",
			description: "Describes how to use the REPL",
			callback:    callbackHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 location areas in the Pokemon world",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 location areas in the Pokemon world",
			callback:    callbackMapb,
		},
		"explore": {
			name:        "explore <location_area>",
			description: "Displays the pokemon available in a specific area",
			callback:    callbackExplore,
		},
		"catch": {
			name:        "catch <pokemon_name>",
			description: "Makes an attempt to capture the given pokemon",
			callback:    callbackCatch,
		},
		"inspect": {
			name:        "inspect <pokemon_name>",
			description: "View information about given pokemon",
			callback:    callbackInspect,
		},
	}
}

func startRepl(cfg *config) {
	reader := bufio.NewScanner(os.Stdin)
	commandList := getCommands()
	for {
		fmt.Print("Pokedex > ")

		reader.Scan()
		text := reader.Text()
		cleanTxt := cleanInput(text)
		if len(cleanTxt) == 0 {
			continue
		}

		commandChosen := cleanTxt[0]
		args := []string{}
		if len(cleanTxt) > 1 {
			args = cleanTxt[1:]
		}

		cmd, ok := commandList[commandChosen]
		if !ok {
			fmt.Println("Invalid or unknown command")
			continue
		}
		err := cmd.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}

	}
}

func cleanInput(text string) []string {
	trimmed := strings.TrimSpace(text)
	lowered := strings.ToLower(trimmed)
	words := strings.Fields(lowered)

	return words
}
