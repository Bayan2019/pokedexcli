package main

import (
	"errors"
	"fmt"
	// "github.com/Bayan2019/pokedexcli/internal/pokeapi"
)

// 4. PokeAPI
func commandMap(cfg *config, args ...string) error {
	// pokeapiClient := pokeapi.NewClient()

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		return err
	}

	fmt.Println("=================================================================")
	fmt.Println("")
	// The map command displays the names of 20 location areas in the Pokemon world.
	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	fmt.Println("")
	fmt.Println("=================================================================")
	// Each subsequent call to map should display the next 20 locations, and so on.
	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLoactionAreaURL = resp.Previous

	return nil
}

// 4. PokeAPI
func commandMapb(cfg *config, args ...string) error {
	// pokeapiClient := pokeapi.NewClient()
	if cfg.prevLoactionAreaURL == nil {
		return errors.New("You are on first page")
	}
	// Similar to the map command, however, instead of displaying the next 20 locations,
	// it displays the previous 20 locations
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
