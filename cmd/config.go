package main

import (
	"time"

	"github.com/kelseyhightower/envconfig"
	"go.uber.org/zap/zapcore"
)

// Config configs of project.
type Config struct {
	APP          string         `json:"APP_NAME" default:"streaming-user"`
	Port         string         `envconfig:"APP_PORT" default:"8080"`
	LogLevel     LogLevelConfig `envconfig:"LOG_LEVEL" default:"info"`
	DBDriver     string         `envconfig:"DB_DRIVER" required:"true"`
	DBSource     string         `envconfig:"DB_SOURCE" required:"true"`
	DBTimeout    time.Duration  `envconfig:"DB_HEALTH_TIMEOUT" default:"1s"`
	ReadTimeout  time.Duration  `envconfig:"READ_TIMEOUT" default:"2s"`
	WriteTimeout time.Duration  `envconfig:"WRITE_TIMEOUT" default:"2s"`
	SchemaPath   string         `envconfig:"SCHEMA_PATH"`
	SchemaName   string         `envconfig:"SCHEMA_NAME"`
}

func newConfig() (*Config, error) {
	cfg := &Config{}
	err := envconfig.Process("", cfg)
	if err != nil {
		return nil, err
	}
	return cfg, nil
}

// LogLevelConfig log level config.
type LogLevelConfig struct {
	Value zapcore.Level
}
