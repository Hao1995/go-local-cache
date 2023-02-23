// Package localcache provides the functionality of storing any data in memory
package localcache

import (
	"sync"
	"time"
)

// Cache defines the interface for a local cache.
type Cache interface {
	Get(key string) (interface{}, bool)
	Set(key string, value interface{})
}

// localCache implements the Cache interface using a map.
type localCache struct {
	data map[string]interface{}
	ttl  int
	mu   sync.Mutex
}

// Get retrieves a value from the cache by key.
func (c *localCache) Get(key string) (interface{}, bool) {
	c.mu.Lock()
	defer c.mu.Unlock()

	value, ok := c.data[key]
	return value, ok
}

// Set sets a key/value pair in the cache with expiration in ttl.
func (c *localCache) Set(key string, value interface{}) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
	time.AfterFunc(time.Duration(c.ttl)*time.Second, func() {
		c.mu.Lock()
		defer c.mu.Unlock()

		delete(c.data, key)
	})
}

// New creates a new instance of the localCache.
func New(ttl int) Cache {
	return &localCache{
		data: make(map[string]interface{}),
		ttl:  ttl,
	}
}
