package utils

import (
	"log"
	"time"
)

type CacheItemsMock struct {
	Data  interface{}
	Found bool
}

type CacheMock struct {
	Items CacheItemsMock
}

func (c *CacheMock) Get(key string) (interface{}, bool) {
	return c.Items.Data, c.Items.Found
}

func (c *CacheMock) Set(key string, value interface{}, ttl time.Duration) {
	log.Printf("[Cache] success setting cache %s", key)
	items := CacheItemsMock{
		Data:  value,
		Found: true,
	}
	c.Items = items
}
