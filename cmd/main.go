package main

import (
	"github.com/Rizskysky/ShortifyGo/internal/handler"
	"github.com/Rizskysky/ShortifyGo/internal/repository"
	"github.com/Rizskysky/ShortifyGo/internal/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	//wiring
	repo := repository.New()

	usecase := usecase.New(repo)

	h := handler.New(usecase)

	g := gin.Default()

	g.POST("/shorten", h.CreateShortUrl)
	g.GET("/:code", h.Redirect)
	g.Run()
}
