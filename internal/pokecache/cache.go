package pokecache

import (
	"sync"
	"time"
)

var CacheInst *Cache

func init() {
	CacheInst = NewCache(10 * time.Second)
}

type cacheEntry struct {
	val       []byte
	createdAt time.Time
}

type Cache struct {
	cacheEntries map[string]cacheEntry
	mutex        sync.Mutex
}

func (c *Cache) Add(key string, val []byte) {
	go func() {
		c.mutex.Lock()
		c.cacheEntries[key] = cacheEntry{
			val:       val,
			createdAt: time.Now(),
		}
		c.mutex.Unlock()
	}()

}

func (c *Cache) Get(key string) (bool, []byte) {

	cacheEntry, ok := c.cacheEntries[key]

	if !ok {
		return ok, nil
	}
	return ok, cacheEntry.val
}

func (c *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)

	go func() {
		for {
			<-ticker.C
			for key, entry := range c.cacheEntries {
				if time.Since(entry.createdAt) > interval {
					delete(c.cacheEntries, key)
				}
			}
		}
	}()
}

func NewCache(interval time.Duration) *Cache {
	cache := Cache{
		cacheEntries: make(map[string]cacheEntry),
	}
	cache.reapLoop(interval)
	return &cache
}
