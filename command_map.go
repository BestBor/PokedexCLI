package main

import (
	"fmt"
)

func callbackMap(cfg *config, args ...string) error {
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range res.Results {
		fmt.Printf(" -%s\n", area.Name)
	}

	cfg.nextLocURL = res.Next
	cfg.prevLocURL = res.Previous

	return nil
}
