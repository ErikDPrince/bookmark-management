package endpoint

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ErikDPrince/bookmark-management/internal/handler"
	"github.com/ErikDPrince/bookmark-management/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestPasswordEndpoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name           string
		setupTestHTTP  func(t *testing.T) *httptest.ResponseRecorder
		expectedStatus int
		expectRespLen  int
	}{
		{
			name: "success",

			setupTestHTTP: func(t *testing.T) *httptest.ResponseRecorder {
				svcMock := mocks.NewPassword(t)
				svcMock.On("GeneratePassword").Return("1234567890", nil)
				testHandler := handler.NewPassword(svcMock)

				app := gin.New()
				app.GET("/password", testHandler.GenPass)
				req := httptest.NewRequest(http.MethodGet, "/password", nil)
				respRec := httptest.NewRecorder()
				app.ServeHTTP(respRec, req)
				return respRec

			},
			expectedStatus: http.StatusOK,
			expectRespLen:  25,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := tc.setupTestHTTP(t)

			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectRespLen, rec.Body.Len())

		})
	}
}
