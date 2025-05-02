package main

import (
	"time"

	"github.com/Bayan2019/pokedexcli/internal/pokeapi"
)

// import (
//
//	"bufio"
//	"fmt"
//	"os"
//)

// PokeAPI
// Your commands should now accept a pointer to a "config" struct as a parameter.
type config struct {
	pokeapiClient pokeapi.Client
	// This struct will contain the Next and Previous URLs
	// that you'll need to paginate through location areas.
	nextLocationAreaURL *string
	prevLoactionAreaURL *string
	caughtPokemon       map[string]pokeapi.Pokemon
}

func main() {
	// Your commands should now accept a pointer to a "config" struct as a parameter.
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
		caughtPokemon: make(map[string]pokeapi.Pokemon),
	}
	// start up an interactive REPL
	startRepl(&cfg)
}
