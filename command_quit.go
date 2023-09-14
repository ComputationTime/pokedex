package main

import "os"

func callbackQuit(cfg *config) error {
	os.Exit(0)
	return nil
}
