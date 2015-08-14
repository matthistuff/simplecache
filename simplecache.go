package simplecache

import (
	"github.com/matthistuff/simpleconf"
)

type Cache struct {
	*simpleconf.Config
}

type CacheMap map[string]interface{}

func New(fileName string) (*Cache, error) {
	config, err := simpleconf.New(fileName, make(CacheMap))

	if err != nil {
		return nil, err
	}

	return &Cache{config}, nil
}

func (c Cache) Get(key string) interface{} {
	return c.Assert()[key]
}

func (c Cache) Assert() CacheMap {
	return c.Data.(CacheMap)
}

func (c *Cache) Clear() {
	c.Data = make(CacheMap)
}