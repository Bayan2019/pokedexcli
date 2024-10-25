package main

import "fmt"

// 3 REPL
// help: prints a help message describing how to use the REPL
func commandHelp(cfg *config, args ...string) error {
	fmt.Println("=================================")
	fmt.Println("Welcome to the Pokedex!")
	fmt.Println("Usage:")
	fmt.Println("")
	availableCommands := getCommands()
	for _, cmd := range availableCommands {
		fmt.Printf("- %s: %s\n", cmd.name, cmd.description)
	}
	fmt.Println("=================================")
	return nil
}
