// Package cache: provides an in-memory key-value store with time-based expiration.
package cache

import (
	"sync"
	"time"
)

// cacheEntry represents a single entry in the cache, storing the creation time and the value.
type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

// Cache represents an in-memory key-value store with time-based expiration.
type Cache struct {
	cache    map[string]cacheEntry
	mutex    sync.Mutex
	interval time.Duration
}

// Add adds a new entry to the cache with the given key and value.
func (c *Cache) Add(key string, val []byte) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.cache[key] = cacheEntry{
		createdAt: time.Now(),
		val:       val,
	}
}

// Get retrieves the value associated with the given key from the cache.
func (c *Cache) Get(key string) ([]byte, bool) {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	entry, ok := c.cache[key]
	return entry.val, ok
}

// reapLoop is a goroutine that runs indefinitely, calling reap() at regular intervals.
func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	for range ticker.C {
		c.reap()
	}
}

// reap removes all entries from the cache that have expired.
func (c *Cache) reap() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	for key, entry := range c.cache {
		if entry.createdAt.Before(time.Now().UTC().Add(-c.interval)) {
			delete(c.cache, key)
		}
	}
}

// NewCache creates a new Cache instance with a given reap interval.
func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cache:    make(map[string]cacheEntry),
		mutex:    sync.Mutex{},
		interval: interval,
	}
	go cache.reapLoop()
	return &cache
}
