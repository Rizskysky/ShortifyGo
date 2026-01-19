package main

import (
	"github.com/Rizskysky/ShortifyGo/internal/handler"
	"github.com/Rizskysky/ShortifyGo/internal/repository"
	"github.com/Rizskysky/ShortifyGo/internal/usecase"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Setup Layer Data (Repository)
	repo := repository.NewURLRepository()

	// 2. Setup Layer Bisnis (Usecase) - butuh repo
	logic := usecase.NewShortenerUsecase(repo)

	// 3. Setup Layer Handler (Gin) - butuh usecase
	h := handler.NewHttpHandler(logic)

	// 4. Setup Framework Gin
	r := gin.Default()

	// Define Routes
	r.POST("/shorten", h.Create) // Buat short link
	r.GET("/:code", h.Redirect)  // Redirect link

	// Jalankan Server
	r.Run(":8080")
}
