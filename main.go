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
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(time.Hour),
	}
	// fmt.Println("Hello World")
	startRepl(&cfg)
}
