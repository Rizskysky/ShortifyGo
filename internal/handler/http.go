package handler

import (
	"net/http"

	"github.com/Rizskysky/ShortifyGo/internal/usecase"

	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	Usecase *usecase.ShortenerUsecase
}

func NewHttpHandler(u *usecase.ShortenerUsecase) *HttpHandler {
	return &HttpHandler{Usecase: u}
}

// Create: Menangani POST /shorten
func (h *HttpHandler) Create(c *gin.Context) {
	// var input struct {
	// 	URL string `json:"url"`
	// }

	// TODO 6: Binding JSON body ke struct input
	// Gunakan c.ShouldBindJSON(&input)

	// TODO 7: Panggil h.Usecase.CreateShortURL
	// Jika error, return c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

	// Jika sukses, return JSON
	// c.JSON(http.StatusOK, gin.H{"short_url": "http://localhost:8080/" + code})
	//DELETE: PLACEHOLDER
	c.JSON(http.StatusOK, nil)
}

// Redirect: Menangani GET /:code
func (h *HttpHandler) Redirect(c *gin.Context) {
	// code := c.Param("code") // Mengambil parameter dari URL

	// TODO 8: Panggil h.Usecase.GetOriginalURL(code)

	// Jika sukses redirect:
	// c.Redirect(http.StatusFound, originalURL)
	//DELETE: PLACEHOLDER
	c.JSON(http.StatusOK, nil)
}
