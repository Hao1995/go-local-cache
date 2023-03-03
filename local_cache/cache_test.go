package local_cache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCache_GetNilIfDataNotExist(t *testing.T) {
	cache := New(1)

	value, ok := cache.Get("key")
	assert.False(t, ok)
	assert.Nil(t, value)
}

func TestCache_SetDifferentDataTypeSuccessfully(t *testing.T) {
	testCases := []struct {
		name string
		key  string
		val  interface{}
	}{
		{name: "string", key: "str", val: "hello"},
		{name: "int", key: "int", val: 42},
		{name: "bool", key: "bool", val: true},
		{name: "float", key: "float", val: 3.14},
		{name: "slice", key: "slice", val: []int{1, 2, 3}},
		{name: "struct", key: "struct", val: struct {
			Name string
			Age  int
		}{"Alice", 30}},
	}

	cache := New(1)

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			cache.Set(tc.key, tc.val)
			value, ok := cache.Get(tc.key)
			assert.True(t, ok)
			assert.Equal(t, tc.val, value)
		})
	}
}

func TestCache_CheckKeyCanBeEdited(t *testing.T) {
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

func TestCache_Expiration(t *testing.T) {
	cache := New(1)

	cache.Set("key", "value")
	time.Sleep(2 * time.Second)

	_, ok := cache.Get("key")
	assert.False(t, ok)
}
