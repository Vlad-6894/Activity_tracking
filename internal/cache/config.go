package cache

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	Host       string        `envconfig:"HOST" required:"true"`
	Port       string        `envconfig:"PORT" default:"6379"`
	Password   string        `envconfig:"PASSWORD" required:"true"`
	DB         int           `envconfig:"DB" default:"0"`
	PoolSize   int           `envconfig:"POOL_SIZE" default:"10"`
	Timeout    time.Duration `envconfig:"TIMEOUT" required:"true"`
	SessionTTL time.Duration `envconfig:"SESSION_TTL" default:"1h"`
}

func (c Config) Address() string {
	return fmt.Sprintf("%s:%s", c.Host, c.Port)
}

func NewConfig() (Config, error) {
	var cfg Config
	if err := envconfig.Process("REDIS", &cfg); err != nil {
		return Config{}, fmt.Errorf("process redis envconfig: %w", err)
	}

	return cfg, nil
}
