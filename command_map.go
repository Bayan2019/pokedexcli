package main

import (
	"errors"
	"fmt"
	// "github.com/Bayan2019/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	// pokeapiClient := pokeapi.NewClient()

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLoactionAreaURL = resp.Previous

	return nil
}

func commandMapb(cfg *config) error {
	// pokeapiClient := pokeapi.NewClient()
	if cfg.prevLoactionAreaURL == nil {
		return errors.New("You are on first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLoactionAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLoactionAreaURL = resp.Previous

	return nil
}
