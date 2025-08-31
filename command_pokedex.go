package main

import (
	"fmt"
)

func callbackPokedex(cfg *config, args ...string) error {
	fmt.Println("Pokemon Captured & in the Pokedex:")
	for _, pokemon := range cfg.caughtPokemons {
		fmt.Printf(" - %s\n", pokemon.Name)
	}

	return nil
}
