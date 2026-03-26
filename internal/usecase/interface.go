package usecase

import "github.com/Rizskysky/ShortifyGo/internal/entity"

type UseCase interface {
	Shorten(url string) (
		entity.CreateUrlResponse,
		error,
	)
	FindByCode(code string) (
		entity.Url,
		error,
	)
	Redirect(code string) (
		string,
		error,
	)
	GetAll() (
		[]entity.Url,
		error,
	)
	Delete(code string) error
	DeleteAll()
}
