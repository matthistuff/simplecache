package simplecache

import (
	"github.com/matthistuff/simpleconf"
)

type Cache struct {
	*simpleconf.Config
}

func New(fileName string) (*Cache, error) {
	config, err := simpleconf.New(fileName, make(map[string]interface{}))

	if err != nil {
		return nil, err
	}

	return &Cache{config}, nil
}

func (c Cache) Get(key string) interface{} {
	return c.Assert()[key]
}

func (c Cache) Assert() map[string]interface{} {
	return c.Data.(map[string]interface{})
}

func (c *Cache) Clear() {
	c.Data = make(map[string]interface{})
}