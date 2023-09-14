package main

import "fmt"

func callbackHelp(cfg *config, args ...string) error {
	fmt.Println("Welcome to the Pokedex. Here are the available commands:")
	for name, command := range getCommands() {
		fmt.Printf(" - %s: %s\n", name, command.description)
	}
	return nil
}
