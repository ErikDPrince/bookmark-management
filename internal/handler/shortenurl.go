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

type ShortenURLRequest struct {
	URL string `json:"url"`
	Exp int64  `json:"exp" default:"3600"`
}

// ShortenURL godoc
// @Summary Shorten URL
// @Description Generate random short code and store mapping URL
// @Tags Links
// @Accept json
// @Produce json
// @Param request body ShortenURLRequest true "Shorten URL request"
// @Success 200 {object} map[string]string
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/links/shorten [post]
func (s *shortenURLHandler) ShortenURL(c *gin.Context) {

	// layer input
	// call service <input:code> --> url --> output

	// redirect client to url
	input := &ShortenURLRequest{}
	if err := c.ShouldBindJSON(input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
		return
	}

	code, err := s.shortenURLService.ShortenURL(c, input.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "Shorten URL generated successfully!",
	})
}
func (s *shortenURLHandler) GetURL(c *gin.Context) {
	// lay input
	code := c.Param("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// call service <input:code> --> url --> output
}
