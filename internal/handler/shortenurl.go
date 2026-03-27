package handler

import (
	"net/http"

	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
)

type shortenURLHandler struct {
	shortenURLService service.ShortenURL
}

type ShortenURLHandler interface {
	ShortenURL(c *gin.Context)
}

func NewShortenURLHandler(shortenURLService service.ShortenURL) ShortenURLHandler {
	return &shortenURLHandler{
		shortenURLService: shortenURLService,
	}
}

type shortenURLRequest struct {
	URL string `json:"url" `
}

func (s *shortenURLHandler) ShortenURL(c *gin.Context) {
	input := &shortenURLRequest{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(400, gin.H{"error": "invalid request"})
		return
	}

	key, err := s.shortenURLService.ShortenURL(c, input.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"key": key,
	})
}
