package main

import (
	"fmt"
)

func commandPokedex(cfg *Config, args ...string) error {
	fmt.Println()
	fmt.Println("Your Pokedex:")
	if len(cfg.caughtPokemon) == 0 {
		fmt.Println("hm, your Pokedex seems to be empty...")
		fmt.Println("Try to catch some Pokemon via 'catch <pokemon>'!")
	} else {
		for _, pokemon := range cfg.caughtPokemon {
			fmt.Printf("  - %s\n", pokemon.Name)
		}
		fmt.Println("You can inspect them via 'inspect <pokemon>'")
	}

	fmt.Println()
	return nil
}
