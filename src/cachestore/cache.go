package cachestore

import (
	"errors"
	cache "inmemcache/src/cachestore/cache"
)

const (
	LRU = iota
	LFU
	TTL
)

type Cache interface {
	//Get fetches the value from cache
	Get(key string) (interface{}, bool)

	//Put store the value in cache
	Put(key string, value interface{})

	//delete the key and value from cache
	Delete(key string) bool
}

func NewCache(size, evictionPolicy int) (Cache, error) {
	var service Cache
	switch evictionPolicy {
	case LRU:
		service = cache.NewLRUCacheService(size)
	case LFU:
		return nil, errors.New("LFU eviction policy is not implemented")
	case TTL:
		return nil, errors.New("TTL eviction policy is not implemented")
	}
	return service, nil
}
