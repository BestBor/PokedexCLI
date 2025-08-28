package main

import "github.com/BestBor/pokedexcli/internal/pokeapi"

type config struct {
	pokeapiClient pokeapi.Client
	nextLocURL    *string
	prevLocURL    *string
}

func main() {

	cfg := config{
		pokeapiClient: pokeapi.NewClient(),
	}

	startRepl(&cfg)
}
