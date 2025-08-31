package main

import (
	"time"

	"github.com/BestBor/pokedexcli/internal/pokeapi"
)

type config struct {
	pokeapiClient  pokeapi.Client
	nextLocURL     *string
	prevLocURL     *string
	caughtPokemons map[string]pokeapi.Pokemon
}

func main() {

	cfg := config{
		pokeapiClient:  pokeapi.NewClient(time.Hour),
		caughtPokemons: make(map[string]pokeapi.Pokemon),
	}

	startRepl(&cfg)
}
