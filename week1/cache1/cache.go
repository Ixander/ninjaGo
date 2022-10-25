package cache1

type Cache struct {
	items map[string]interface{}
}

func New() *Cache {

	items := make(map[string]interface{})

	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key string, value interface{}) {
	c.items[key] = value
}

func (c *Cache) Get(key string) interface{} {
	item, found := c.items[key]

	if !found {
		return nil
	}
	return item
}
func (c *Cache) Delete(key string) {
	_, found := c.items[key]

	if !found {
		return
	}

	delete(c.items, key)

}
