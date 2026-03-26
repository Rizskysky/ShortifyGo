package main

import (
	"net/http"

	"github.com/Rizskysky/ShortifyGo/internal/handler"
	"github.com/Rizskysky/ShortifyGo/internal/repository"
	"github.com/Rizskysky/ShortifyGo/internal/usecase"
)

func main() {
	repo := repository.NewRepository()
	service := usecase.NewUseCase(repo)
	handling := handler.NewHandler(service)

	app := http.NewServeMux()
	app.HandleFunc("/shorten", handling.Shorten)
	app.HandleFunc("/redirect", handling.Redirect)
	app.HandleFunc("/list", handling.GetAll)
	app.HandleFunc("/find", handling.FindByCode)
	app.HandleFunc("/delete", handling.Delete)
	app.HandleFunc("/clear", handling.DeleteAll)

	if err := http.ListenAndServe(":8080", app); err != nil {
		panic(err)
	}
}
