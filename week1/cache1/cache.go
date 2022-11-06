package cache1

import (
	"errors"
	"sync"
	"time"
)

type Cache struct {
	cache map[string]Value
	sync.RWMutex
}

type Value struct {
	Value       interface{}
	CreatedTime time.Time
	Expiration  int64
}

func New() *Cache {

	values := make(map[string]Value)

	cache := Cache{
		cache: values,
	}

	//go cache.StartCleaner()

	return &cache
}

func (c *Cache) Set(key string, value interface{}, ttl time.Duration) {

	expiration := time.Now().Add(ttl).UnixNano()

	c.Lock()
	defer c.Unlock()

	c.cache[key] = Value{
		Value:       value,
		CreatedTime: time.Now(),
		Expiration:  expiration,
	}
}

func (c *Cache) Get(key string) (interface{}, error) {
	c.RLock()
	defer c.RUnlock()

	item, exists := c.cache[key]
	if !exists {
		return nil, errors.New("no values for the key " + key + ".")
	}

	if time.Now().UnixNano() > item.Expiration {
		return nil, errors.New("value for the key  " + key + " outdated")
	}

	return item.Value, nil
}

func (c *Cache) Delete(key string) error {
	c.Lock()
	defer c.Unlock()

	if _, found := c.cache[key]; !found {
		return errors.New("no values for the key " + key + ".")
	}

	delete(c.cache, key)
	return nil
}

//func (c *Cache) StartCleaner() {
//	for {
//		<-time.After(time.Second * 10)
//
//		if keys := c.expiredKeys(); len(keys) != 0 {
//			c.clearValues(keys)
//		}
//	}
//}

//func (c *Cache) expiredKeys() (keys []string) {
//
//	c.RLock()
//	defer c.RUnlock()
//
//	for k, i := range c.cache {
//		if time.Now().UnixNano() > i.Expiration && i.Expiration > 0 {
//			keys = append(keys, k)
//		}
//	}
//
//	return
//}

//func (c *Cache) clearValues(keys []string) {
//
//	c.Lock()
//	defer c.Unlock()
//
//	for _, k := range keys {
//		delete(c.cache, k)
//	}
//}
