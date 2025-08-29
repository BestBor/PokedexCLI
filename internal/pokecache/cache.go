package pokecache

import "time"

type Cache struct {
	cache map[string]cacheEntry
}

type cacheEntry struct {
	createdAt time.Time
	val       []byte
}

func NewCache(interval time.Duration) Cache {
	c := Cache{
		cache: make(map[string]cacheEntry),
	}
	go c.reapLoop(interval)
	return c
}

func (c *Cache) Add(key string, val []byte) {
	c.cache[key] = cacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (c *Cache) Get(key string) ([]byte, bool) {
	cacheEnt, ok := c.cache[key]
	return cacheEnt.val, ok
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		c.outWithTheOld(interval)
	}
}

func (c *Cache) outWithTheOld(interval time.Duration) {
	tminusinterval := time.Now().UTC().Add(-interval)
	for k, v := range c.cache {
		if v.createdAt.Before(tminusinterval) {
			delete(c.cache, k)
		}
	}
}
