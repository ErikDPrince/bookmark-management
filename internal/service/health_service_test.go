package service

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestHealthService_Check(t *testing.T) {

	t.Parallel()
	instanceID := uuid.New().String()
	svc := NewHealthService("bookmark-management", instanceID)
	assert.NotNil(t, svc)

	resp := svc.Check()
	assert.NotNil(t, resp)

	assert.Equal(t, "ok", resp.Message)
	assert.Equal(t, "bookmark-management", resp.ServiceName)
	assert.Equal(t, instanceID, resp.InstanceID)

}
