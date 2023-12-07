package main

import (
	"errors"
	"fmt"
	// "github.com/Bayan2019/pokedexcli/internal/pokeapi"
)

func commandExplore(cfg *config, args ...string) error {
	if len(args) != 1 {
		return errors.New("No location area is provided")
	}

	locationAreaName := args[0]

	locationArea, err := cfg.pokeapiClient.GetLocationArea(locationAreaName)
	if err != nil {
		return err
	}

	fmt.Println("=================================================================")
	fmt.Println("")

	fmt.Printf("Exploring %s...\n", locationArea.Name)
	fmt.Println("Found Pokemon:")
	for _, pokemon := range locationArea.PokemonEncounters {
		fmt.Printf(" - %s\n", pokemon.Pokemon.Name)
	}
	fmt.Println("")
	fmt.Println("=================================================================")

	return nil
}
