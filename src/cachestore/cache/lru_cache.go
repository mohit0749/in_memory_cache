package service

import (
	"container/list"
)

type keyValue struct {
	key   string
	value interface{}
}

type lruCache struct {
	list *list.List
	mp   map[string]*list.Element
	size int
}

func (c *lruCache) Put(key string, value interface{}) {

	elem, ok := c.mp[key]
	if ok {
		c.list.MoveToBack(elem)
		return
	}

	elem = c.list.PushBack(keyValue{key, value})
	c.mp[key] = elem
	c._evict()
}

func (c *lruCache) _evict() {
	if c.list.Len() > c.size {
		elem := c.list.Front()
		if elem != nil {
			keyValueObj, ok := elem.Value.(keyValue)
			if !ok {
				return
			}
			delete(c.mp, keyValueObj.key)
		}
	}
}

func (c *lruCache) Get(key string) (interface{}, bool) {
	if value, ok := c.mp[key]; ok && value != nil {
		c.list.MoveToBack(value)
		return value.Value.(keyValue).value, true
	}
	return nil, false
}

func (c *lruCache) Delete(key string) bool {
	if value, ok := c.mp[key]; ok {
		c.list.Remove(value)
		return true
	}
	return false
}

func NewLRUCacheService(size int) *lruCache {
	var service lruCache
	service.list = list.New()
	service.mp = make(map[string]*list.Element)
	service.size = size
	return &service
}
