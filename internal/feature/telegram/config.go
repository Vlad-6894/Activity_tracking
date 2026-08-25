package telegram

import (
	"errors"
	"fmt"
	"net/url"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	BotToken  string `envconfig:"BOT_TOKEN" required:"true"`
	WebAppURL string `envconfig:"WEBAPP_URL"`
}

func NewConfig() (Config, error) {
	var cfg Config

	if err := envconfig.Process("TELEGRAM", &cfg); err != nil {
		return Config{}, fmt.Errorf("process telegram envconfig: %w", err)
	}

	if err := cfg.Validate(); err != nil {
		return Config{}, fmt.Errorf("validate telegram config: %w", err)
	}

	return cfg, nil
}

func (cfg Config) Validate() error {
	var errs []error

	if cfg.BotToken == "" {
		errs = append(errs, errors.New("TELEGRAM_BOT_TOKEN is empty"))
	}

	switch parsed, err := url.Parse(cfg.WebAppURL); {
	case cfg.WebAppURL == "":
		errs = append(errs, errors.New("TELEGRAM_WEBAPP_URL is empty; set NGROK_DOMAIN or url itself"))
	case err != nil:
		errs = append(errs, fmt.Errorf("TELEGRAM_WEBAPP_URL is not a url: %w", err))
	case parsed.Scheme != "https":
		errs = append(errs, fmt.Errorf("TELEGRAM_WEBAPP_URL must use https, got %q", parsed.Scheme))
	case parsed.Host == "":
		errs = append(errs, errors.New("TELEGRAM_WEBAPP_URL has no host"))
	}

	return errors.Join(errs...)
}
