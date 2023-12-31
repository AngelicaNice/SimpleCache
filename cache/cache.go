package cache

type Cache struct {
	storage map[string]interface{}
}

func (c Cache) Get(key string) interface{} {
	value, ok := c.storage[key]
	if ok {
		return value
	}
	return ok
}

func (c Cache) Set(key string, value interface{}) {
	c.storage[key] = value
}

func (c Cache) Delete(key string) {
	delete(c.storage, key)
}

func New() Cache {
	return Cache{
		storage: make(map[string]interface{}),
	}
}
