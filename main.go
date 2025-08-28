package main

import (
	"fmt"
	"log"

	"github.com/BestBor/pokedexcli/internal/pokeapi"
)

type config struct {
	Next     *string
	Previous *string
}

func main() {

	pokeapiClient := pokeapi.NewClient()
	res, err := pokeapiClient.ListLocationAreas()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res)

	// startRepl()
}
