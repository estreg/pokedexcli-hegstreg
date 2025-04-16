package main;

import (
	"time"

	"github.com/estreg/pokedexcli-hegstreg/internal/pokeapi" // The path from the root of the project to the internal package.
)

func main() {
	// Create API Client
	pokeClient := pokeapi.NewClient(5 * time.Second, time.Minute*5) // This calls a function named Newclient from your pokeapi package.
	
	// Initialize Config
	cfg := &Config{ 								 // Creates a new Config struct and initializes it with the pokeapi client,
		pokeapiClient: pokeClient,					 // returns a pointer to this config and nextLocationsURL and prevLocationsURL are left as their zero values - nil for pointers.
	}
	startRepl(cfg)									 // Calls a function named startRepl and passes it your config. Look at repl.go.
}