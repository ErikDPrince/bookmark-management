package endpoint

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ErikDPrince/bookmark-management/internal/handler"
	"github.com/ErikDPrince/bookmark-management/internal/repository"
	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/alicebob/miniredis/v2"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestShortenEndpoint(t *testing.T) {
	t.Parallel()

	mock := miniredis.RunT(t)
	rdb := redis.NewClient(&redis.Options{Addr: mock.Addr()})

	repo := repository.NewURLStorage(rdb)
	codeGen := service.NewCodeGenerator()

	svc := service.NewShortenURLService(repo, codeGen)

	h := handler.NewShortenURLHandler(svc)

	app := gin.New()
	app.POST("/v1/links/shorten", h.ShortenURL)

	body := []byte(`{"url":"https://example.com","exp":3600}`)
	req := httptest.NewRequest(http.MethodPost, "/v1/links/shorten", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()
	app.ServeHTTP(rec, req)
	assert.Equal(t, http.StatusOK, rec.Code)

	var resp struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}
	assert.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	assert.NotEmpty(t, resp.Code)

	// Verify Redis saved mapping: code -> url
	val, err := rdb.Get(t.Context(), resp.Code).Result()
	assert.NoError(t, err)
	assert.Equal(t, "https://example.com", val)
}
