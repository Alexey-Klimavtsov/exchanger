package cache

import (
	"sync"
	"time"
)

type item struct {
	value     any
	expiresAt time.Time
}

type Cache struct {
	mu    sync.Mutex
	items map[string]item
}

func New() *Cache {
	return &Cache{
		items: make(map[string]item),
	}
}

func (c *Cache) Get(key string) (any, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	it, ok := c.items[key]
	if !ok || time.Now().After(it.expiresAt) {
		return nil, false
	}
	return it.value, true
}

func (c *Cache) Set(key string, value any, ttl time.Duration) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.items[key] = item{
		value:     value,
		expiresAt: time.Now().Add(ttl),
	}
}
