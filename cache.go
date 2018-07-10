package golru

import (
	"fmt"

	"github.com/hashicorp/golang-lru"
)

// Cache is our cache object
type Cache struct {
	Store *lru.Cache
}

// New ...
func New() *Cache {
	cache, err := lru.New(100)
	if err != nil {
		fmt.Println(err)
	}

	return &Cache{Store: cache}
}

// Add adds things to the cache
func (c *Cache) Add(key, value string) error {
	c.Store.Add(key, value)
	return nil
}
