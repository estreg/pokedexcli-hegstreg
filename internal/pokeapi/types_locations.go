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

// Location
// The struct was build via https://mholt.github.io/json-to-go/.
// --> I used this JSON string https://pokeapi.co/api/v2/location-area/canalave-city-area.
// You need not worry about pointers, not like above, because the name of the location always "points to something".
type Location struct {
	EncounterMethodRates []struct {
		EncounterMethod struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"encounter_method"`
		VersionDetails []struct {
			Rate    int `json:"rate"`
			Version struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"encounter_method_rates"`
	GameIndex int `json:"game_index"`
	ID        int `json:"id"`
	Location  struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"location"`
	Name  string `json:"name"`
	Names []struct {
		Language struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"language"`
		Name string `json:"name"`
	} `json:"names"`
	PokemonEncounters []struct {
		Pokemon struct {
			Name string `json:"name"`
			URL  string `json:"url"`
		} `json:"pokemon"`
		VersionDetails []struct {
			EncounterDetails []struct {
				Chance          int           `json:"chance"`
				ConditionValues []interface{} `json:"condition_values"`
				MaxLevel        int           `json:"max_level"`
				Method          struct {
					Name string `json:"name"`
					URL  string `json:"url"`
				} `json:"method"`
				MinLevel int `json:"min_level"`
			} `json:"encounter_details"`
			MaxChance int `json:"max_chance"`
			Version   struct {
				Name string `json:"name"`
				URL  string `json:"url"`
			} `json:"version"`
		} `json:"version_details"`
	} `json:"pokemon_encounters"`
}