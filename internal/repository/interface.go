package repository

import "github.com/Rizskysky/ShortifyGo/internal/entity"

type Repository interface {
	Shorten(url entity.Url) error
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
