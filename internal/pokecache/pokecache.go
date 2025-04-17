package pokecache

import (
	"sync"
	"time"
)

// Cache represents a thread-safe in-memory cache that automatically removes
// stale entries after a specified interval.
type Cache struct {
	cache map[string]cacheEntry		// Map to store cached data with string keys.
	mux   *sync.Mutex				// Mutex to protect concurrent access to the map.
}

// cacheEntry represents a single entry in the cache with its creation timestamp.
type cacheEntry struct {
	createdAt time.Time				// When this entry was added to the cache.
	val       []byte
}

// NewCache creates and returns a new Cache that will automatically remove entries
// older than the specified interval.
func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
		mux:   &sync.Mutex{},
	}

	// Important and Golangs specific way of concurrent programming: Start the reaper goroutine to clean up old entries.
	go c.reapLoop(interval)

	return c
}

// Add stores a new value in the cache with the specified key.
func (c *Cache) Add(key string, value []byte) {
	c.mux.Lock()
	defer c.mux.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       value,
	}
}

// Get retrieves a value from the cache by its key and
// returns the value and a boolean indicating whether the key was found.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()
	val, ok := c.cache[key]
	return val.val, ok
}

// reapLoop runs in a separate goroutine and periodically removes stale entries
// from the cache based on the specified interval.
func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C { 						// This seems like it comes out of the blue, but look at the bottom of the file for further explaination: X
		c.reap(time.Now().UTC(), interval)
	}
}

// reap checks all cache entries and removes any that are older than the specified duration.
// This implements the cache expiration mechanism to prevent unlimited growth.
func (c *Cache) reap(now time.Time, last time.Duration) {
	c.mux.Lock()
	defer c.mux.Unlock()
	for k, v := range c.cache {
		if v.createdAt.Before(now.Add(-last)) { // Remove entry if it was created before (now - interval).
			delete(c.cache, k)
		}
	}
}

// X: "ticker.C" is a channel that's provided by the "time.Ticker" struct. In Go, the "time.Ticker" type has a field
// called "C" which is a receive-only channel of type "<-chan time.Time". This is defined in the standard library.
// When you create a ticker with "time.NewTicker(interval)", it returns a Ticker struct with this channel already set up.
// The ticker sends the current time on this channel at regular intervals (defined by the duration you passed).
// The capitalization of C means it's an exported (public) field of the Ticker struct, allowing you to access it from outside the time package.
// For more information look at the official documentation: https://pkg.go.dev/time#Ticker