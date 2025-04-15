package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type LocationAreaResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func commandMap(cfg *Config) error {
	url := "https://pokeapi.co/api/v2/location-area?limit=20"

	if cfg.Next != "" { // If we have a Next URL stored in config, than this is used instead.
		url = cfg.Next
	}
	return fetchLocationAreas(url, cfg)
}

func commandMapb(cfg *Config) error {
    if cfg.Previous == "" {
        fmt.Println("you're on the first page")
        return nil
    }
    return fetchLocationAreas(cfg.Previous, cfg)
}

func fetchLocationAreas(url string, cfg *Config) error {
	resp, err := http.Get(url) // Getting Step: put the json string in it.
	if err != nil {
		return err
	}
	defer resp.Body.Close() // Makes shure to close at the end, if something goes wrong.

	if resp.StatusCode != http.StatusOK { // Not necessaryn but makes the code more robust, in case something happens with the API. Makes Error Dectition more easy.
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	body, err := io.ReadAll(resp.Body) // Reading Step: Read the response body (completely into memory).
	if err != nil {
		return err
	}
	
	var locationResp LocationAreaResponse // Parsing Step: Parse the response json via unmarshall, because we use not so much data. Decode would be unnecessary overhead.
	if err := json.Unmarshal(body, &locationResp); err != nil {
		return err
	}

	if locationResp.Next != nil { // Updating Step A: Update the config with Next URL.
		cfg.Next = *locationResp.Next
	} else {
		cfg.Next = ""
	}
	if locationResp.Previous != nil { // Updating Step B: Update the config with Previous URL.
        cfg.Previous = *locationResp.Previous
    } else {
        cfg.Previous = ""
    }

	// Now we print the location names to the consol!
	for _, location := range locationResp.Results {
		fmt.Println(location.Name)
	}

	return nil
}