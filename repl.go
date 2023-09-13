package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Printf("Pokedex > ")
		scanner.Scan()
		text := scanner.Text()
		cleaned := cleanInput(text)

		if len(cleaned) == 0 {
			continue
		}

		command := cleaned[0]

		switch command {
		case "quit":
			os.Exit(0)
		case "help":
			fmt.Println("Welcome to the Pokedex. Here are the available commands:")
			fmt.Println(" - help")
			fmt.Println(" - exit")
		default:
			fmt.Println("Invalid command")
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
