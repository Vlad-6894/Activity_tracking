package authService

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleHealtAPI = "https://www.googleapis.com/auth/healthdata.readonly"
)

type Credentials struct {
	ClientID     string `envconfig:"CLIENT_ID" required:"true"`
	ClientSecret string `envconfig:"CLIENT_SECRET" required:"true"`
	RedirectURL  string `envconfig:"REDIRECT_URL" required:"true"`
}

func NewConfig() (*oauth2.Config, error) {
	var cred Credentials
	if err := envconfig.Process("GOOGLE", &cred); err != nil {
		return &oauth2.Config{}, fmt.Errorf("proccess envconfig: %w", err)
	}

	config := &oauth2.Config{
		ClientID:     cred.ClientID,
		ClientSecret: cred.ClientSecret,
		RedirectURL:  cred.RedirectURL,
		Scopes: []string{
			googleHealtAPI,
		},
		Endpoint: google.Endpoint,
	}

	return config, nil
}
