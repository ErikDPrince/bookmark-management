package handler

import (
	"errors"
	"net/http"

	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

type shortenURLHandler struct {
	shortenURLService service.ShortenURL
}

type ShortenURLHandler interface {
	ShortenURL(c *gin.Context)
	GetURL(c *gin.Context)
}

func NewShortenURLHandler(shortenURLService service.ShortenURL) ShortenURLHandler {
	return &shortenURLHandler{
		shortenURLService: shortenURLService,
	}
}

type ShortenURLRequest struct {
	URL string `json:"url" binding:"required,url"`
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
		log.Error().Err(err).Msg("failed to shorten URL")
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    code,
		"message": "Shorten URL generated successfully!",
	})
}

// GetURL resolves a short code and redirects the client to the stored URL.
// @Summary Resolve short link
// @Description Looks up the code in Redis and responds with HTTP 301 Moved Permanently; the Location header is the original URL.
// @Tags Links
// @Param code path string true "Short code"
// @Success 301 {string} string "Redirect (follow Location header)"
// @Header 301 {string} Location "Original URL"
// @Failure 400 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Router /v1/links/{code} [get]
func (s *shortenURLHandler) GetURL(c *gin.Context) {
	// lay input
	code := c.Param("code")

	if code == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid input"})
		return
	}

	// call service <input:code> --> url --> output
	url, err := s.shortenURLService.GetURL(c, code)
	if err != nil {

		if errors.Is(err, service.ErrCodeNotExist) {
			c.JSON(http.StatusNotFound, gin.H{"error": "code not exists"})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	//redirect client to url
	c.Redirect(http.StatusMovedPermanently, url)
}
