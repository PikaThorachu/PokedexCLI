package PokeCache

import (
	"sync"
)

type ThreadSafeCache struct {
	data map[string][]byte
	mu   sync.Mutex
}

func NewThreadSafeCache() *ThreadSafeCache {
	return &ThreadSafeCache{
		data: make(map[string][]byte),
	}
}

// Adds a new entry to the cache
func (c *ThreadSafeCache) Add(key string, value []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

// Gets an entry from the cache if it exists
func (c *ThreadSafeCache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	val, ok := c.data[key]
	return val, ok
}
