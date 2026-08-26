package server

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Addr            string        `envconfig:"ADDR" default:":8080"`
	ShutdownTimeout time.Duration `envconfig:"SHUTDOWN_TIMEOUT" default:"10s"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("HTTP", &cfg); err != nil {
		return Config{}, fmt.Errorf("process http envconfig: %w", err)
	}

	return cfg, nil
}
