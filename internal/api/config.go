package api

import (
	"strings"

	"github.com/google/uuid"
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	AppPort     string `envconfig:"APP_PORT" default:"8080"`
	ServiceName string `envconfig:"SERVICE_NAME" default:"bookmark-management"`
	InstanceID  string `envconfig:"INSTANCE_ID" default:""`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}

	if strings.TrimSpace(cfg.InstanceID) == "" {
		cfg.InstanceID = uuid.New().String()
	}

	return cfg, nil

}
