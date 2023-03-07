package localcache

import (
	"testing"
	"time"

	"github.com/stretchr/testify/suite"
)

const (
	EXPIRATION_TTL = 1
)

type localCacheSuite struct {
	suite.Suite
	localcache *localCache
}

func (s *localCacheSuite) SetupSuite() {}

func (s *localCacheSuite) SetupTest() {
	s.localcache = New(EXPIRATION_TTL).(*localCache)
	s.localcache.data = make(map[string]interface{})
}

func (s *localCacheSuite) TearDownSuite() {}

func (s *localCacheSuite) TearDownTest() {}

func (s *localCacheSuite) TestCacheGetNilIfDataNotExist() {
	value, ok := s.localcache.Get("key")

	s.False(ok)
	s.Nil(value)
}

func (s *localCacheSuite) TestCacheSetDifferentDataTypeSuccessfully() {
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

	for _, tc := range testCases {
		s.localcache.Set(tc.key, tc.val)

		value, ok := s.localcache.data[tc.key]
		s.True(ok)
		s.Equal(tc.val, value)
	}
}

func (s *localCacheSuite) TestCacheCheckKeyCanBeEdited() {
	s.localcache.data = map[string]interface{}{
		"key": "v1 value",
	}

	s.localcache.Set("key", "v2 value")

	s.Equal("v2 value", s.localcache.data["key"])
}

func (s *localCacheSuite) TestCacheExpiration() {
	s.localcache.Set("key", "value")

	s.Equal("value", s.localcache.data["key"])

	time.Sleep((EXPIRATION_TTL + 1) * time.Second)

	_, ok := s.localcache.data["key"]
	s.False(ok)
}

func TestLocalcacheSuite(t *testing.T) {
	suite.Run(t, new(localCacheSuite))
}
