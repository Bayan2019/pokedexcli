package main

import "os"

// REPL
// exit: exits the program
func commandExit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
