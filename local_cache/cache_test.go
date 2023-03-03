package local_cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCacheSetAndGet(t *testing.T) {
	cache := New(1)

	cache.Set("key", "value")
	value, ok := cache.Get("key")
	assert.True(t, ok)
	assert.Equal(t, "value", value)

	cache.Set("key", "new value")
	value, ok = cache.Get("key")
	assert.True(t, ok)
	assert.Equal(t, "new value", value)
}

func TestCacheExpiration(t *testing.T) {
	cache := New(1)

	cache.Set("key", "value")
	time.Sleep(2 * time.Second)

	_, ok := cache.Get("key")
	assert.False(t, ok)
}
