package main

import (
	"fmt"
)

func callbackHelp(cfg *config) error {
	fmt.Println("Welcome to the Pokedex help menu!")
	fmt.Println("Available commands:")
	commandList := getCommands()
	for _, cmd := range commandList {
		fmt.Printf("  %-10s %s\n", cmd.name, cmd.description)
	}
	return nil
}
