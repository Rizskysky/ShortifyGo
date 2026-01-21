package handler

import (
	"net/http"

	"github.com/Rizskysky/ShortifyGo/internal/entity"
	"github.com/Rizskysky/ShortifyGo/internal/usecase"
	"github.com/gin-gonic/gin"
)

type HttpHandler struct {
	shortUrlUsecase *usecase.ShortUrlUsecase
}

func New(u *usecase.ShortUrlUsecase) *HttpHandler {
	return &HttpHandler{
		shortUrlUsecase: u,
	}
}

func displayErrJson(c *gin.Context, err error) {
	c.JSON(http.StatusBadRequest, entity.BadResponse{
		Message: err.Error(),
	},
	)

}

func (h *HttpHandler) CreateShortUrl(c *gin.Context) {
	req := entity.CreateShortUrlRequest{}

	if err := c.ShouldBind(&req); err != nil {
		displayErrJson(c, err)
		return
	}

	id, err := h.shortUrlUsecase.CreateShortURL(req.URL)
	if err != nil {
		displayErrJson(c, err)
		return
	}

	c.JSON(http.StatusCreated, entity.CreateShortUrlResponse{
		Code: id,
	})
}

func (h *HttpHandler) Redirect(c *gin.Context) {
	code := c.Param("code")

	original, err := h.shortUrlUsecase.GetOriginalURL(code)
	if err != nil {
		displayErrJson(c, err)
		return
	}
	c.Redirect(http.StatusFound, original)
}
