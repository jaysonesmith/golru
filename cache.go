package golru

import (
	"github.com/hashicorp/golang-lru"
)

// Cache is our cache object
type Cache struct {
	Store *lru.Cache
}

// Collection is a collection of individual items
type Collection struct {
	Items []Item
}

// Item is a data describing an item
type Item struct {
	ID    string
	Type  string
	Dur   string
	Group string
	URL   string
}

// New sets up a cache
func New(size int) *Cache {
	cache, _ := lru.New(size)
	return &Cache{Store: cache}
}

// AddItems adds a group of items to the cache
func (c *Cache) AddItems(collection Collection) error {
	for _, item := range collection.Items {
		c.Store.Add(item.ID, item)
	}
	return nil
}
