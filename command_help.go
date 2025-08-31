package main

import (
	"fmt"
)

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Available commands:")
	commandList := getCommands()
	for _, cmd := range commandList {
		fmt.Printf("  %-25s %s\n", cmd.name, cmd.description)
	}
	return nil
}
