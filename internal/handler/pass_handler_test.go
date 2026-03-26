package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ErikDPrince/bookmark-management/internal/service/mocks"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
)

func TestPasswordHandler_GenPass(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string

		setupRequest func(c *gin.Context)
		setupMockSvc func(t *testing.T) *mocks.Password

		expectedStatus int
		expectResp     string
	}{
		{
			name: "success",

			setupRequest: func(c *gin.Context) {
				c.Request = httptest.NewRequest(http.MethodGet, "/password", nil)
			},

			setupMockSvc: func(t *testing.T) *mocks.Password {
				svcMock := mocks.NewPassword(t)
				svcMock.On("GeneratePassword").Return("1234567890", nil)
				return svcMock
			},

			expectedStatus: http.StatusOK,
			expectResp:     "{\"password\":\"1234567890\"}",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			rec := httptest.NewRecorder()
			gc, _ := gin.CreateTestContext(rec)

			tc.setupRequest(gc)
			mockSvc := tc.setupMockSvc(t)
			testHandler := NewPassword(mockSvc)

			testHandler.GenPass(gc)
			assert.Equal(t, tc.expectedStatus, rec.Code)
			assert.Equal(t, tc.expectResp, rec.Body.String())
		})
	}
}
