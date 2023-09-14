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
	callback    func(*config) error
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
		"map": {
			name:        "map",
			description: "show next set of up to 20 locations",
			callback:    callbackMap,
		},
		"mapb": {
			name:        "mapb",
			description: "show previous set of up to 20 locations",
			callback:    callbackMapb,
		},
	}
}

func startREPL(cfg *config) {
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
		err := command.callback(cfg)
		if err != nil {
			fmt.Println(err)
		}
	}
}

func cleanInput(str string) []string {
	lowered := strings.ToLower(str)
	words := strings.Fields(lowered)
	return words
}
