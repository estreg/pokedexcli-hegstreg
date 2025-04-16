package pokeapi

import (
	"net/http"
	"time" // Not necassary, but provides functionality for measuring and displaying time --> for later requests with timeouts.

	"github.com/estreg/pokedexcli-hegstreg/internal/pokecache"
)

// Client - a client for the PokeAPI. This will be used to make all HTTP requests to the PokeAPI.
type Client struct {
	cache pokecache.Cache
	httpClient http.Client
}

// NewClient - creates a new PokeAPI client with the specified timeout.
func NewClient(timeout, cacheInterval time.Duration) Client {
	return Client{
		cache: pokecache.NewCache(cacheInterval),
		httpClient: http.Client{
			Timeout: timeout,
		},
	}
}

// The key concept here is that we're creating a reusable client that can make HTTP requests with a configurable timeout.
// This is more robust than using the default HTTP client, as it prevents our application from hanging indefinitely on slow or failed requests.