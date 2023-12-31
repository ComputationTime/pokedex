package main

import (
	"fmt"
)

func callMap(cfg *config, url *string) error {

	locations, err := cfg.pokeapiClient.ListLocationAreas(url)
	if err != nil {
		return err
	}
	cfg.nextLocationURL = locations.Next
	cfg.previousLocationURL = locations.Previous
	fmt.Println("List of some locations:")

	for _, location := range locations.Results {
		fmt.Printf(" - %s\n", location.Name)
	}
	return nil

}

func callbackMap(cfg *config, args ...string) error {
	return callMap(cfg, cfg.nextLocationURL)
}

func callbackMapb(cfg *config, args ...string) error {
	return callMap(cfg, cfg.previousLocationURL)
}
