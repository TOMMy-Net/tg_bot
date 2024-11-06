package tools

import (
	"sync"
	"tg_bot/internal/models"
)

type UserId int

type ApplicationCache struct {
	cache map[UserId]models.Application
	mu    sync.RWMutex
}

func NewCache() *ApplicationCache {
	return &ApplicationCache{
		cache: make(map[UserId]models.Application, 100),
	}
}

func (a *ApplicationCache) Read(id UserId) models.Application {
	a.mu.RLock()
	m := a.cache[id]
	a.mu.RUnlock()
	return m
}

func (a *ApplicationCache) Store(id UserId, model models.Application) {
	a.mu.Lock()
	a.cache[id] = model
	a.mu.Unlock()
}

func (a *ApplicationCache) Delete(id UserId)  {
	a.mu.Lock()
	delete(a.cache, id)
	a.mu.Unlock()
}