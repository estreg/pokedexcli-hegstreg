package main

import (
	"errors"
	"fmt"
)

func commandMap(cfg *Config) error {
	// Get locations data using the API client
	locationsResp, err := cfg.pokeapiClient.ListLocations(cfg.nextLocationsURL) // Calls the "ListLocations" method from your PokeAPI client
	if err != nil {																// and Passes the current "nextLocationsURL" from the config (which could be nil on first call).
		return err																
	}

	// Update pagination URLs in the config: This ensures future calls to map and mapb have the correct URLs.
	cfg.nextLocationsURL = locationsResp.Next
	cfg.prevLocationsURL = locationsResp.Previous
	
	// Print the location names
	for _, loc := range locationsResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}

func commandMapb(cfg *Config) error {
	//Simple Check, if we are on the first page.
	if cfg.prevLocationsURL == nil {
		return errors.New("you're on the first page")
	}

	locationResp, err := cfg.pokeapiClient.ListLocations(cfg.prevLocationsURL)
	if err != nil {
		return err
	}

	cfg.nextLocationsURL = locationResp.Next
	cfg.prevLocationsURL = locationResp.Previous

	for _, loc := range locationResp.Results {
		fmt.Println(loc.Name)
	}
	return nil
}