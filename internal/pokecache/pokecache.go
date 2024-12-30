package pokecache

import (
	"sync"
	"time"
)

type Cache struct {
	cache    map[string]cacheEntry
	mu       *sync.Mutex
	interval time.Duration
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	newCache := Cache{
		cache:    map[string]cacheEntry{},
		mu:       &sync.Mutex{},
		interval: interval,
	}

	go newCache.reapLoop()

	return newCache
}

func (c *Cache) Add(key string, val []byte) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.cache[key] = cacheEntry{createdAt: time.Now(), val: val}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()
	cacheEntry, exists := c.cache[key]
	return cacheEntry.val, exists
}

func (c *Cache) reapLoop() {
	ticker := time.NewTicker(c.interval)
	defer ticker.Stop()

	for range ticker.C {
		reap(c)
	}
}

func reap(c *Cache) {
	for cacheEntry, val := range c.cache {
		c.mu.Lock()
		defer c.mu.Unlock()
		if time.Since(val.createdAt) > c.interval {

			delete(c.cache, cacheEntry)
		}
	}
}
