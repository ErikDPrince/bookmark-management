package endpoint

import (
	"testing"
)

func TestPasswordEndpoint(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name string
		setupRequest func(c *gin.Context)
		setupMockSvc func(t *testing.T) *mocks.Password
		expectedStatus int
		expectResp string
	}