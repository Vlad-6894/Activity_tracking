package core_config

import (
	"fmt"
	"time"

	"github.com/kelseyhightower/envconfig"
)

type PostgresConfig struct {
	Host     string        `envconfig:"HOST" required:"true"`
	Port     string        `envconfig:"PORT" default:"5432"`
	User     string        `envconfig:"USER" required:"true"`
	Password string        `envconfig:"PASSWORD" required:"true"`
	Database string        `envconfig:"DB" required:"true"`
	Timeout  time.Duration `envconfig:"TIMEOUT" required:"true"`
}

func (cfg PostgresConfig) ConnectionString() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
}

func NewPostgresConfig() (PostgresConfig, error) {
	var cfg PostgresConfig
	if err := envconfig.Process("POSTGRES", &cfg); err != nil {
		return PostgresConfig{}, fmt.Errorf("process envconfig: %w", err)
	}
	return cfg, nil
}

func NewPostgresConfigMust() PostgresConfig {
	cfg, err := NewPostgresConfig()
	if err != nil {
		panic(fmt.Errorf("get postgres connection pool config: %w", err))
	}
	return cfg
}
