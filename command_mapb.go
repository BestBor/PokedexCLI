package main

import (
	"errors"
	"fmt"
)

func callbackMapb(cfg *config, args ...string) error {
	if cfg.prevLocURL == nil {
		return errors.New("located on first page")
	}
	res, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLocURL)
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
