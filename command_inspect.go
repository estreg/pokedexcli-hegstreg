package main

import (
	"errors"
	"fmt"
)

func commandInspect(cfg *Config, args ...string) error {
	if len(args) != 1 {
		return errors.New("you must provide a pokemon name to inspect")
	}

	name := args[0]
    pokemon, ok := cfg.caughtPokemon[name]
    if !ok {
        return errors.New("you have not caught that pokemon")
    }
    fmt.Println()
    fmt.Println("Name:", pokemon.Name)
    fmt.Println("Height:", pokemon.Height)
    fmt.Println("Weight:", pokemon.Weight)
    fmt.Println("Base Experience:", pokemon.BaseExperience)

    fmt.Println("Stats:")
    for _, stat := range pokemon.Stats {
        fmt.Printf("  -%s: %d\n", stat.Stat.Name, stat.BaseStat)
    }
    
    fmt.Println("Types:")
    for _, typeInfo := range pokemon.Types {
        fmt.Printf("  - %s\n", typeInfo.Type.Name)
    }
	fmt.Println()
    
    return nil
}