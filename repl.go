package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startRepl(cfg *config) {
	// Get input using bufio.NewScanner, .Scan, and .Text
	scanner := bufio.NewScanner(os.Stdin)
	// Use an infinite for loop to keep the REPL running.
	for {
		// It should print a prompt
		fmt.Print("pokedex >")
		// you should block and wait for some input
		// Get input using .Scan
		scanner.Scan()
		// Once input is received
		// Get input using .Text
		text := scanner.Text()
		// you should parse it
		cleaned := cleanInput(text)
		if len(cleaned) == 0 {
			continue
		}

		commandName := cleaned[0]
		args := []string{}
		if len(cleaned) > 1 {
			args = cleaned[1:]
		}

		availableCommands := getCommands()

		command, ok := availableCommands[commandName]

		if !ok {
			fmt.Println("Invalid Command")
			continue
		}
		// and then execute a command
		err := command.callback(cfg, args...)
		if err != nil {
			fmt.Println(err)
		}
	}
}

type cliCommand struct {
	name        string
	description string
	callback    func(*config, ...string) error
}

func getCommands() map[string]cliCommand {
	// I recommend storing your commands in a map.
	// Remember, functions are first-class citizens in Go, you can pass them around as values!
	return map[string]cliCommand{
		// help: prints a help message describing how to use the REPL
		"help": {
			name:        "help",
			description: "Displays a help message",
			callback:    commandHelp,
		},
		// exit: exits the program
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"map": {
			name:        "map",
			description: "Returns the next list of some location areas",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Returns the previous list of some location areas",
			callback:    commandMapb,
		},
		"explore": {
			name:        "explore {location_area}",
			description: "Returns detailed information about particular location area",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch {pokemon_name}",
			description: "Attempts to catch Pokemon and Adds it to your pokedex",
			callback:    commandCatch,
		},
		"inspect": {
			name:        "inspect {pokemon_name}",
			description: "Inspects caught Pokemon and lists its characteristics",
			callback:    commandInspect,
		},
		"pokedex": {
			name:        "pokedex",
			description: "Returns the list of caught Pokemons",
			callback:    commandPokedex,
		},
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
