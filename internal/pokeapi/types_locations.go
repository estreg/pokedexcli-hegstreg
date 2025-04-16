package pokeapi

// RespShallowLocations - This represents the response from the PokeAPI when listing locations.
// "Shallow" indicates this is not the full detailed data, but a list-level response.
// The struct was build via https://mholt.github.io/json-to-go/. "Next" and "Previous" had to be changed to "*string". They are pointers to strings (*string) because they can be null in the JSON.
// If there's no next or previous page, the API returns null for these fields. Using pointers allows us to distinguish between an empty string and a null value.

type RespShallowLocations struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

// This struct matches the structure of the JSON response from the PokeAPI's location list endpoint. When we make a request to get a list of locations,
// we'll decode the JSON response into an instance of this struct, which will make it easy to access the information in a typesafe way.