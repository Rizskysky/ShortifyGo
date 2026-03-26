package usecase

import (
	"errors"
	"math/rand"
	"strings"

	"github.com/Rizskysky/ShortifyGo/internal/entity"
	"github.com/Rizskysky/ShortifyGo/internal/repository"
)

type useCaseImpl struct {
	repo repository.Repository
}

func NewUseCase(repo repository.Repository) UseCase {
	return &useCaseImpl{
		repo: repo,
	}
}

func (u *useCaseImpl) Shorten(url string) (
	entity.CreateUrlResponse,
	error,
) {
	if url == "" {
		return entity.CreateUrlResponse{}, errors.New("url is empty")
	}

	if !strings.HasPrefix(url, "http://") && !strings.HasPrefix(url, "https://") {
		return entity.CreateUrlResponse{}, errors.New("url must start with http:// or https://")
	}

	maxAttempts := 10
	for attempts := 0; attempts < maxAttempts; attempts++ {
		code := generateRandomString(6)
		_, err := u.repo.FindByCode(code)
		if err != nil {
			request := entity.Url{
				Code:      code,
				OriginUrl: url,
				Clicks:    0,
			}

			err := u.repo.Shorten(request)
			if err != nil {
				return entity.CreateUrlResponse{}, err
			}

			return entity.CreateUrlResponse{Url: url, Code: request.Code}, nil
		}
	}

	return entity.CreateUrlResponse{}, errors.New("failed to generate unique code, please try again")
}

func (u *useCaseImpl) FindByCode(code string) (
	entity.Url,
	error,
) {
	if code == "" {
		return entity.Url{}, errors.New("code is empty")
	}

	return u.repo.FindByCode(code)
}

func (u *useCaseImpl) Redirect(code string) (
	string,
	error,
) {
	if code == "" {
		return "", errors.New("code is empty")
	}

	originalUrl, err := u.repo.Redirect(code)
	if err != nil {
		return "", err
	}

	return originalUrl, nil
}

func (u *useCaseImpl) GetAll() (
	[]entity.Url,
	error,
) {
	return u.repo.GetAll()
}

func (u *useCaseImpl) Delete(code string) error {
	if code == "" {
		return errors.New("code is empty")
	}

	return u.repo.Delete(code)
}

func (u *useCaseImpl) DeleteAll() {
	u.repo.DeleteAll()
}

func generateRandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
