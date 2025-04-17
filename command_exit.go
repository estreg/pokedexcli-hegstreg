package main

import (
	"fmt"
	"os"
)

func commandExit(cfg *Config, args ...string) error {
	fmt.Println("Closing the Pokedex... Goodbye!")
	os.Exit(0)
	return nil // This line will never execude due to os.Exit
}