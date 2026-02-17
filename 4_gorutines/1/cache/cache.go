package cache

import "sync"

type Cache struct {
	mu   sync.RWMutex
	data map[string]string
}

func NewCache() *Cache {
	return &Cache{
		data: make(map[string]string),
	}
}

// Get чтение из мапы.
func (c *Cache) Get(key string) (string, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()

	// Проверяем значение в мапе
	// Если есть, будет true
	val, ok := c.data[key]

	return val, ok
}

// Set запись в мапу.
func (c *Cache) Set(key string, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	c.data[key] = value
}

// Delete удаление из мапы.
func (c *Cache) Delete(key string) {
	c.mu.Lock()
	defer c.mu.Unlock()

	delete(c.data, key)
}
