package repository

import (
	"errors"

	"github.com/Rizskysky/ShortifyGo/internal/entity"
)

type repositoryImpl struct {
	memory map[string]entity.Url
}

func NewRepository() Repository {
	return &repositoryImpl{
		memory: make(map[string]entity.Url),
	}
}

func (r *repositoryImpl) Shorten(url entity.Url) error {
	r.memory[url.Code] = url
	return nil
}

func (r *repositoryImpl) FindByCode(code string) (
	entity.Url,
	error,
) {
	url, exists := r.memory[code]
	if !exists {
		return entity.Url{}, errors.New("url not found")
	}

	return url, nil
}

func (r *repositoryImpl) Redirect(code string) (
	string,
	error,
) {
	url, exists := r.memory[code]
	if !exists {
		return "", errors.New("url not found")
	}

	url.Clicks++
	r.memory[code] = url

	return url.OriginUrl, nil
}

func (r *repositoryImpl) GetAll() (
	[]entity.Url,
	error,
) {
	var urls []entity.Url
	for _, url := range r.memory {
		urls = append(urls, url)
	}
	return urls, nil
}

func (r *repositoryImpl) Delete(code string) error {
	delete(r.memory, code)
	return nil
}

func (r *repositoryImpl) DeleteAll() {
	r.memory = make(map[string]entity.Url)
}
