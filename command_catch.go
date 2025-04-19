package main

import (
	"errors"
	"fmt"
	"math/rand"
)

func commandCatch(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a name of a pokemon from this area to catch")
	}
	//maybe implement a statemant that only a pokemon could be catched, that is in this area.


	name := args[0]
	pokemon, err := cfg.pokeapiClient.GetPokemon(name)
	if err != nil {
		return err
	}
	fmt.Println()
	fmt.Printf("Throwing a Pokeball at %s...\n", pokemon.Name) // check if ".Name" works
	
	randomNum := rand.Intn(43)
	catchChance := 42 - pokemon.BaseExperience/10
	if catchChance < 2 {									  
		catchChance = 3						
	}
	
	if randomNum < catchChance {
		fmt.Printf("%s was caught!\n", pokemon.Name)
        cfg.caughtPokemon[pokemon.Name] = pokemon
		fmt.Println("You may now inspect it with the 'inspect' command.")
	} else {
		fmt.Printf("%s escaped!\n", pokemon.Name)
		fmt.Println("You can try it again!")
	}
	fmt.Println()
	
	return nil
}