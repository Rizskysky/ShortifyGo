package usecase

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/Rizskysky/ShortifyGo/internal/entity"
	"github.com/Rizskysky/ShortifyGo/internal/repository"
)

const MaximumLenRandID = 5

type ShortUrlUsecase struct {
	repo *repository.InMemoryURLRepository
}

func New(r *repository.InMemoryURLRepository) *ShortUrlUsecase {
	return &ShortUrlUsecase{
		repo: r,
	}
}

func (u *ShortUrlUsecase) CreateShortURL(url string) (string, error) {

	if url == "" {
		return "", fmt.Errorf("URL is mandatory!")
	}

	id := generateRandomString(MaximumLenRandID)

	urlObj := entity.URL{
		ID:          id,
		OriginalURL: url,
	}
	if err := u.repo.CreateUrl(urlObj); err != nil {
		return "", err
	}

	return id, nil
}

func (u *ShortUrlUsecase) GetOriginalURL(id string) (string, error) {
	url, err := u.repo.GetURLByID(id)
	if err != nil {
		return "", err
	}

	return url.OriginalURL, nil
}

func generateRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	rand.NewSource(time.Now().UnixNano())
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
