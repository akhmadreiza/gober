package utils

import (
	"log"
	"sync"
	"time"
)

type CacheItems struct {
	Data      interface{}
	ExpiresAt time.Time
}

type Cache struct {
	items map[string]CacheItems
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{items: make(map[string]CacheItems)}
}

func (c *Cache) Get(key string) (interface{}, bool) {
	log.Printf("[Cache] getting cache %s", key)
	c.mu.RLock()
	defer c.mu.RUnlock()
	item, found := c.items[key]

	if !found {
		log.Printf("[Cache] cache not found %s", key)
		return nil, false
	}

	if time.Now().After(item.ExpiresAt) {
		log.Printf("[Cache] cache %s is expired", key)
		return nil, false
	}

	log.Printf("[Cache] success getting cache")
	return item.Data, true
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {
	log.Printf("[Cache] setting cache %s", key)
	c.mu.Lock()
	defer c.mu.Unlock()
	c.items[key] = CacheItems{
		Data:      value,
		ExpiresAt: time.Now().Add(ttl),
	}
	log.Printf("[Cache] success setting cache %s. expired at: %s", key, c.items[key].ExpiresAt)
}
