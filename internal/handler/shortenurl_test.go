package handler

import (
	"bytes"
	"context"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type mockShortenURLService struct {
	shortenURLFunc func(ctx context.Context, url string) (string, error)
}

func (m *mockShortenURLService) ShortenURL(ctx context.Context, url string) (string, error) {
	return m.shortenURLFunc(ctx, url)
}

func (m *mockShortenURLService) GetURL(ctx context.Context, code string) (string, error) {
	return "", errors.New("not used in this test")
}

func TestShortenURLHandler_ShortenURL(t *testing.T) {
	t.Parallel()
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		body           string
		shortenFn      func(ctx context.Context, url string) (string, error)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "invalid json",
			body:           `{"url":`,
			shortenFn:      func(ctx context.Context, url string) (string, error) { return "Ab12Xyz", nil },
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"invalid request"}`,
		},
		{
			name: "service error",
			body: `{"url":"https://google.com"}`,
			shortenFn: func(ctx context.Context, url string) (string, error) {
				return "", errors.New("boom")
			},
			expectedStatus: http.StatusInternalServerError,
			expectedBody:   `{"error":"internal server error"}`,
		},
		{
			name: "success",
			body: `{"url":"https://google.com","exp":3600}`,
			shortenFn: func(ctx context.Context, url string) (string, error) {
				return "Ab12Xyz", nil
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"code":"Ab12Xyz","message":"Shorten URL generated successfully!"}`,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			r := gin.New()
			h := NewShortenURLHandler(&mockShortenURLService{shortenURLFunc: tt.shortenFn})
			r.POST("/v1/links/shorten", h.ShortenURL)
			req := httptest.NewRequest(http.MethodPost, "/v1/links/shorten", bytes.NewBufferString(tt.body))
			req.Header.Set("Content-Type", "application/json")
			w := httptest.NewRecorder()
			r.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
		})
	}
}
