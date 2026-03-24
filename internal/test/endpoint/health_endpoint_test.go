package endpoint

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/ErikDPrince/bookmark-management/internal/handler"
	"github.com/ErikDPrince/bookmark-management/internal/service"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/assert/v2"
	"github.com/stretchr/testify/require"
	"github.com/google/uuid"
)


func TestHealthEndpoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		serviceName string
		instanceID string
		expectedStatus int
		expectedMsg string 
	}{
		{
			name: "success",
			serviceName: "bookmark_service",
			instanceID: uuid.New().String(),
			expectedStatus: http.StatusOK,
			expectedMsg: "ok",
		},
	}


for _, tc := range testCases {
	t.Run(tc.name, func(t *testing.T){
		t.Parallel()

		healthSvc := service.NewHealthService(tc.serviceName, tc.instanceID)
		healthHandler := handler.NewHealthHandler(healthSvc)

		app := gin.New()
		app.GET("/health-check", healthHandler.Check)
		req := httptest.NewRequest(http.MethodGet, "/health-check", nil)
		respRec := httptest.NewRecorder()
		app.ServeHTTP(respRec, req)



		require.Equal(t, tc.expectedStatus, respRec.Code)
		var body map[string]interface{}
		err := json.NewDecoder(respRec.Body).Decode(&body)
		require.NoError(t, err)

		assert.Equal(t, tc.expectedMsg, body["message"])
		assert.Equal(t, tc.serviceName, body["service_name"])
		assert.Equal(t, tc.instanceID, body["instance_id"])
	})
}
}