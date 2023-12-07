package main

import (
	"fmt"
	"log"
	// "github.com/Bayan2019/pokedexcli/internal/pokeapi"
)

func commandMap(cfg *config) error {
	// pokeapiClient := pokeapi.NewClient()

	resp, err := cfg.pokeapiClient.ListLocationAreas(cfg.nextLocationAreaURL)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Location areas:")
	for _, area := range resp.Results {
		fmt.Printf(" - %s\n", area.Name)
	}

	cfg.nextLocationAreaURL = resp.Next
	cfg.prevLoactionAreaURL = resp.Previous

	return nil
}
