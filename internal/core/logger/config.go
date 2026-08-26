package logger

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Level  string `envconfig:"LEVEL" default:"DEBUG"`
	Folder string `envconfig:"FOLDER" required:"true"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("LOGGER", &cfg); err != nil {
		return Config{}, fmt.Errorf("process logger envconfig: %w", err)
	}

	return cfg, nil
}
