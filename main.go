package main

import (
	"time"

	"github.com/Bayan2019/pokedexcli/internal/pokeapi"
)

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationAreaURL *string
	prevLoactionAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	// fmt.Println("Hello World")
	startRepl(&cfg)
}
