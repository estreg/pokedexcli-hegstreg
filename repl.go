package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/estreg/pokedexcli-hegstreg/internal/pokeapi"
)

type Config struct {
	pokeapiClient    pokeapi.Client // "pokeapi.Client" indicates this is a custom type called Client from a package called pokeapi.
	nextLocationsURL *string
	prevLocationsURL *string
	caughtPokemon 	 map[string]pokeapi.Pokemon
}
// Note that this is storing the actual client value, not a pointer to the client. Whether this is optimal depends on how the pokeapi.Client type is designed and how large it is.

func startRepl(cfg *Config) {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInput(scanner.Text())
		if len(words) == 0 {
			continue
		}

		command := words[0]
		args := []string{}
		if len(words) > 1 {
			args = words[1:]
		}
		
		cmd, exists := getCommands()[command]
		if exists {
			err := cmd.callback(cfg, args...)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Unknown command")
			continue
		}
	}
}

func cleanInput(text string) []string {
	output := strings.ToLower(text)
	words := strings.Fields(output)
	return words
}

type cliCommand struct { // Registry of commands. More an abstraction Layer.
	name		string
	description string
	callback 	func(*Config, ...string) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{ // Adding XXX Command to the command registery or register XXX command means to place it this map. You have the name of the command and the cliCommand Struc like Structure.
		"catch": {
			name:        "catch <pokemon>",
			description: "try to catch a pokemon",
			callback:    commandCatch,
		},
		"explore": {
			name:        "explore <location_name>",
			description: "Explore a location",
			callback:    commandExplore,
		},
		"exit" : {
			name:		 "exit",
			description: "Exit the Pokedex.",
			callback:	 commandExit,
		},
		"help": {
			name:		 "help",
			description: "Displays a help message.",
			callback:	 commandHelp,
		},
		"map" : {
			name:		 "map",
			description: "Displays the names of 20 locations. Each subsequent call displays the next 20.",
			callback:	 commandMap,
		},
		"mapb" : {
			name:		 "mapb",
			description: "Displays the previous 20 locations.",
			callback:	 commandMapb,
		},
		
	}
}
