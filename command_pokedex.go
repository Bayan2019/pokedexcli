package main

import (
	// "errors"
	"fmt"
	// "math/rand"
	// "github.com/Bayan2019/pokedexcli/internal/pokeapi"
)

func commandPokedex(cfg *config, args ...string) error {
	fmt.Println("")
	fmt.Println("=================================================================")
	fmt.Println("Your Pokedex: ")

	for _, pokemon := range cfg.caughtPokemon {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	fmt.Println("=================================================================")
	fmt.Println("")

	// fmt.Printf("Name: %s\n", pokemon.Name)
	// fmt.Printf("Height: %v\n", pokemon.Height)
	// fmt.Printf("Weight: %v\n", pokemon.Weight)
	// fmt.Println("Stats:")
	// for _, stat := range pokemon.Stats {
	// 	fmt.Printf("- %s: %v\n", stat.Stat.Name, stat.BaseStat)
	// }
	// fmt.Println("Types:")
	// for _, typ := range pokemon.Types {
	// 	fmt.Printf("- %s\n", typ.Type.Name)
	// }

	return nil
}
