package pokeapi

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"
)

// ListLocations -
// This defines a method on the Client type (because of the (c *Client) receiver).
// Takes a pageURL parameter which is a pointer to a string (can be nil).
// Returns two values: a RespShallowLocations struct and an error.

func (c *Client) ListLocations(pageURL *string) (RespShallowLocations, error) {
    url := baseURL + "/location-area"
    if pageURL != nil {
        url = *pageURL 							 // This is for pagination - using "next" and "previous" URLs from prior API responses.
    }

    // When making a request, you first check if the data is already cached.
    if val, ok := c.cache.Get(url); ok {         // The URL serves as the cache key.
		locationsResp := RespShallowLocations{}  // When Dat is found in cache.
		err := json.Unmarshal(val, &locationsResp)
		if err != nil {
			return RespShallowLocations{}, err
		}

		return locationsResp, nil                // Return the data immediately, avoiding an HTTP request.
	}


    req, err := http.NewRequest("GET", url, nil) // Creates a new HTTP GET request for the specified URL. You have to do it like this,
    if err != nil {								 // because you are using a special Client and not the default one.
        return RespShallowLocations{}, err		 // Important and easiely forgotten: The third parameter "nil" indicates there's no request body.
    }

    resp, err := c.httpClient.Do(req) 			 // Executes the HTTP request using the client's HTTP client. Look it up in client.go file.
    if err != nil {								 // --> This is better than using http.Get directly because it uses the configured client with timeout.
        return RespShallowLocations{}, err
    }
    defer resp.Body.Close()						 // Ensures the response body is closed after the function finishes, preventing resource leaks.

    if resp.StatusCode != http.StatusOK {		 // Checks if the HTTP response status code is anything other than 200 OK.
        return RespShallowLocations{}, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
    }

    dat, err := io.ReadAll(resp.Body)			// io.ReadAll is a convenience function that handles reading until EOF
    if err != nil {								// and reads the entire response body into memory as a byte slice.
        return RespShallowLocations{}, err
    }

    locationsResp := RespShallowLocations{}		// Initializes an empty RespShallowLocations struct to hold the parsed JSON data. Look at types_location.go file.
    err = json.Unmarshal(dat, &locationsResp)	// Parses the JSON data in dat into the locationsResp struct. Uses the address operator & to pass a pointer to the struct.
    if err != nil {
        return RespShallowLocations{}, err
    }

    // After getting a successful HTTP response, you store it in the cache.
    c.cache.Add(url, dat)                       // The raw response body (dat) is stored as bytes in the cache.
    return locationsResp, nil
}

// This function encapsulates the API interaction logic, handling all the HTTP request/response details and error cases. It's designed to be modular and reusable,
// with a focus on separation of concerns - it only handles data retrieval, leaving display and user interaction to other parts of the program.
// We use Unmarshal over Decode because:
// - Use Unmarshal when dealing with small to medium-sized complete JSON responses.
// - Use Decode when dealing with very large JSON data, streaming applications, or when parsing JSON in chunks.

// Your cache automatically removes entries older than "cacheInterval" - Look at the pokecache package.
// The cache is thread-safe thanks to the mutex that is implemented. And the cache reaping happens in a background goroutine.