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
	callback    func(*config) error
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

		cmd, ok := commandList[commandChosen]
		if !ok {
			fmt.Println("Invalid or unknown command")
			continue
		}
		err := cmd.callback(cfg)
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
