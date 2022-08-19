package cache

import "github.com/patrickmn/go-cache"

type Memcache struct {
	c cache.Cache
}

func NewMemcache() *Memcache {
	return &Memcache{
		c: *cache.New(cache.NoExpiration, cache.NoExpiration),
	}
}

func (m *Memcache) Get(key string) (interface{}, bool) {
	return m.c.Get(key)
}
