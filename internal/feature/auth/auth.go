package auth

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

type GoogleClientAuthConfig struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
	RedirectURL  string `envconfig:"REDIRECT_URL" required:"true"`
}

func New() (GoogleClientAuthConfig, error) {
	var config GoogleClientAuthConfig
	if err := envconfig.Process("GOOGLE", &config); err != nil {
		return GoogleClientAuthConfig{}, fmt.Errorf("fail process google config: %w", err)
	}

	return config, nil
}
