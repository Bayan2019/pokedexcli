package main

import (
	"errors"
	"fmt"
	// "github.com/Bayan2019/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config, args ...string) error {
	// pokeapiClient := pokeapi.NewClient()

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("=================================================================")
	fmt.Println("")

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	fmt.Println("")
	fmt.Println("=================================================================")

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLoactionAreaURL = resp.Previous

	return nil
}

func commandMapb(cfg *config, args ...string) error {
	// pokeapiClient := pokeapi.NewClient()
	if cfg.prevLoactionAreaURL == nil {
		return errors.New("You are on first page")
	}
	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.prevLoactionAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("=================================================================")
	fmt.Println("")

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	fmt.Println("")
	fmt.Println("=================================================================")

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLoactionAreaURL = resp.Previous

	return nil
}
