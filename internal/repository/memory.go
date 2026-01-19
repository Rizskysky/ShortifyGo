package repository

import (
	"errors"
	"sync"

	"github.com/Rizskysky/ShortifyGo/internal/entity"
)

type URLRepository struct {
	Data map[string]entity.URL
	Mu   sync.RWMutex
}

func NewURLRepository() *URLRepository {
	return &URLRepository{
		Data: make(map[string]entity.URL),
	}
}

// Save menyimpan data URL
func (r *URLRepository) Save(url entity.URL) error {
	r.Mu.Lock()
	defer r.Mu.Unlock()

	// TODO 1: Simpan struct `url` ke map `r.Data`
	// Gunakan url.ID sebagai key

	return nil
}

// FindByID mencari URL berdasarkan short code
func (r *URLRepository) FindByID(id string) (entity.URL, error) {
	r.Mu.RLock()
	defer r.Mu.RUnlock()

	// TODO 2: Cek apakah id ada di map
	// Jika tidak ada, return entity.URL{} dan error baru

	// Jika ada, return datanya
	return entity.URL{}, errors.New("not implemented yet")
}
