package tools

import (
	"fmt"
	"sync"
	"tg_bot/internal/models"
)


type Cache struct {
	cache map[string]any
	mu    sync.RWMutex
}

func NewCache() *Cache {
	return &Cache{
		cache: make(map[string]any, 100),
	}
}

func (c *Cache) read(id string) any {
	c.mu.RLock()
	m := c.cache[id]
	c.mu.RUnlock()
	return m
}

func (c *Cache) store(id string, model any) {
	c.mu.Lock()
	c.cache[id] = model
	c.mu.Unlock()
}

func (c *Cache) delete(id string) {
	c.mu.Lock()
	delete(c.cache, id)
	c.mu.Unlock()
}

func CreateSupportKey(id int64) string {
	return fmt.Sprintf("%d:support", id)
}

func CreateApplicationKey(id int64) string {
	return fmt.Sprintf("%d:application", id)
}

func (c *Cache) ReadApplication(id int64) models.Application{
	a := c.read(CreateApplicationKey(id))
	if v, ok := a.(models.Application); ok{
		return v
	}else{
		return models.Application{}
	}
}

func (c *Cache) DeleteApplication(id int64)  {
	c.delete(CreateApplicationKey(id))
}

func (c *Cache) StoreApplication(id int64, m models.Application)  {
	c.store(CreateApplicationKey(id), m)
}

func (c *Cache) ReadSupport(id int64) models.Support {
	a := c.read(CreateSupportKey(id))
	if v, ok := a.(models.Support); ok{
		return v
	}else{
		return models.Support{}
	}
}

func (c *Cache) DeleteSupport(id int64)  {
	c.delete(CreateSupportKey(id))
}

func (c *Cache) StoreSupport(id int64, m models.Support)  {
	c.store(CreateSupportKey(id), m)
}