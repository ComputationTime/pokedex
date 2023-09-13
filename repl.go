package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type command struct {
	name        string
	description string
	callback    func() error
}

func getCommands() map[string]command {
	return map[string]command{
		"help": {
			name:        "help",
			description: "gives you the list of all commands",
			callback:    callbackHelp,
		},
		"quit": {
			name:        "quit",
			description: "quit the Pokedex",
			callback:    callbackQuit,
		},
	}
}

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

		commandName := cleaned[0]

		command, ok := getCommands()[commandName]

		if !ok {
			fmt.Println("Invalid command. Type help to get available commands")
			continue
		}
		command.callback()
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
