package main

import "os"

func callbackQuit(cfg *config, args ...string) error {
	os.Exit(0)
	return nil
}
