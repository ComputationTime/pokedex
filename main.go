package main

import (
	"fmt"
	"log"

	"github.com/computationtime/pokedex/internal/pokeapi"
)

func main() {
	pokeapiClient := pokeapi.NewClient()

	locations, err := pokeapiClient.ListLocations()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(locations)
	// startREPL()
}
