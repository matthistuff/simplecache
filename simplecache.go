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

func (c Cache) FromCache(key string) interface{} {
	return c.GetData()[key]
}

func (c Cache) GetData() CacheMap {
	return c.Data.(CacheMap)
}