package repository

import (
	"fmt"
	"sync"

	"github.com/Rizskysky/ShortifyGo/internal/entity"
)

type InMemoryURLRepository struct {
	memory map[string]entity.URL
	mu     sync.RWMutex
}

func New() *InMemoryURLRepository {
	return &InMemoryURLRepository{
		memory: make(map[string]entity.URL),
	}
}

func (r *InMemoryURLRepository) CreateUrl(url entity.URL) error {
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.memory[url.ID] = url
	return nil
}

func (r *InMemoryURLRepository) GetURLByID(id string) (entity.URL, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	url, ok := r.memory[id]
	if !ok {
		return entity.URL{}, fmt.Errorf("ID tidak ditemukan")
	}
	return url, nil
}
