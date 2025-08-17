package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var cliCommands map[string]cliCommand

type cliCommand struct {
	name        string
	description string
	callback    func() error
}

func commandExit() error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil
}

func commandHelp() error {
	fmt.Println(`Welcome to the Pokedex!
Usage:
	`)
	for _, cmd := range cliCommands {
		fmt.Printf("  %-10s %s\n", cmd.name, cmd.description)
	}
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
	}
}

func main() {
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
			if err := cmd.callback(); err != nil {
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
