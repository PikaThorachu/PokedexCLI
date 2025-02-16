package PokeCache

import (
	"sync"
	"time"
)

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

type Cache struct {
	entries  map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

func NewCache(interval time.Duration) *Cache {
	cache := &Cache{
		entries:  make(map[string]cacheEntry),
		interval: interval,
	}
	go cache.reapLoop()
	return cache
}

func (cache *Cache) Add(key string, val []byte) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	cache.entries[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

func (cache *Cache) Get(key string) ([]byte, bool) {
	cache.mutex.Lock()
	defer cache.mutex.Unlock()
	entry, ok := cache.entries[key]
	if !ok {
		return nil, false
	}
	return entry.val, true
}

func (cache *Cache) reapLoop() {
	for {
		time.Sleep(cache.interval)

		cache.mutex.Lock()

		now := time.Now()
		for key, entry := range cache.entries {
			if now.Sub(entry.createdAt) > cache.interval {
				delete(cache.entries, key)
			}
		}
		cache.mutex.Unlock()
	}
}
