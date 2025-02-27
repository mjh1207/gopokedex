package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	entries map[string]cacheEntry
	mu *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	Val []byte
}

func NewCache(interval time.Duration) *Cache{
	cache := &Cache {
		entries: make(map[string]cacheEntry),
		mu: &sync.Mutex{},
		interval: interval,
	}

	go cache.reapLoop()

	return cache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()

	c.entries[key] = cacheEntry{
		createdAt: time.Now(),
		Val: val,
	}
	c.mu.Unlock()
}

func (c *Cache) Get(key string) ([]byte, bool) {
	result := []byte{}
	c.mu.Lock()
	defer c.mu.Unlock()
	entry, ok := c.entries[key]
	if !ok {
		return result, false
	}
	return entry.Val, true
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for {
		<-ticker.C
		c.mu.Lock()
		for key := range c.entries {
			if time.Since(c.entries[key].createdAt) > c.interval {
				delete(c.entries, key)
			}
		}
		c.mu.Unlock()
	}
}