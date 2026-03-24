package api

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)
func TestNewConfig_UseEnvInstanceID(t *testing.T) {
	t.Parallel()
	t.Setenv("APP_PORT", "9090")
	t.Setenv("SERVICE_NAME", "bookmark_service")
	t.Setenv("INSTANCE_ID", "fixed-instance-id")
	cfg, err := NewConfig()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, "9090", cfg.AppPort)
	assert.Equal(t, "bookmark_service", cfg.ServiceName)
	assert.Equal(t, "fixed-instance-id", cfg.InstanceID)
}
func TestNewConfig_GenerateUUIDWhenInstanceIDEmpty(t *testing.T) {
	t.Parallel()
	t.Setenv("APP_PORT", "8081")
	t.Setenv("SERVICE_NAME", "bookmark_service")
	t.Setenv("INSTANCE_ID", "")
	cfg, err := NewConfig()
	require.NoError(t, err)
	require.NotNil(t, cfg)
	assert.Equal(t, "8081", cfg.AppPort)
	assert.Equal(t, "bookmark_service", cfg.ServiceName)
	assert.NotEmpty(t, cfg.InstanceID)
	_, parseErr := uuid.Parse(cfg.InstanceID)
	assert.NoError(t, parseErr)
}
