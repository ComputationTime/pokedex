package main

import "os"

func callbackQuit() error {
	os.Exit(0)
	return nil
}
