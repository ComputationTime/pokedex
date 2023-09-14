package main

import (
	"time"

	"github.com/computationtime/pokedex/internal/pokeapi"
)

type config struct {
	pokeapiClient       pokeapi.Client
	nextLocationURL     *string
	previousLocationURL *string
}

func main() {
	cfg := config{
		pokeapiClient: pokeapi.NewClient(30 * time.Minute),
	}
	startREPL(&cfg)
}
