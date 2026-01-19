package usecase

import (
	"errors"
	"math/rand"
	"time"

	"github.com/Rizskysky/ShortifyGo/internal/repository"
)

type ShortenerUsecase struct {
	Repo *repository.URLRepository
}

func NewShortenerUsecase(repo *repository.URLRepository) *ShortenerUsecase {
	return &ShortenerUsecase{Repo: repo}
}

// CreateShortURL: Validasi -> Generate ID -> Simpan
func (u *ShortenerUsecase) CreateShortURL(original string) (string, error) {
	// TODO 3: Validasi input (misal: string kosong)
	if original == "" {
		// return error "URL cannot be empty"
	}

	// Generate ID (Sudah disediakan)
	shortCode := generateRandomString(5)

	// TODO 4: Buat object entity.URL dan panggil Repo.Save()

	return shortCode, nil
}

// GetOriginalURL: Cari data -> Return link asli
func (u *ShortenerUsecase) GetOriginalURL(id string) (string, error) {
	// TODO 5: Panggil Repo.FindByID
	// Return Original URL saja

	return "", errors.New("not implemented yet")
}

// Helper function (biarkan saja)
func generateRandomString(n int) string {
	chars := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := range b {
		b[i] = chars[rand.Intn(len(chars))]
	}
	return string(b)
}
