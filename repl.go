package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

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
		cmd, exists := getCommands()[command]
		if exists {
			err := cmd.callback(cfg)
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
type Config struct {
	Next 	 string
	Previous string
}

type cliCommand struct { // Registry of commands. More an abstraction Layer.
	name		string
	description string
	callback 	func(cfg *Config) error
}

func getCommands() map[string]cliCommand {
	return map[string]cliCommand{ // Adding XXX Command to the command registery or register XXX command means to place it this map. You have the name of the command and the cliCommand Struc like Structure.
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
