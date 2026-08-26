package auth_transport_http

import (
	"github.com/Vlad-6894/Activity_tracking/internal/feature/auth"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

type AuthHTTPHandler struct {
	service          AuthService
	GoogleOathConfig *oauth2.Config
}

type AuthService interface{}

// Это адрес для авторизации пользователя, чтобы мы смогли получить доступ к его шагам
var stepsAddress = "https://www.googleapis.com/auth/googlehealth.activity_and_fitness.readonly"

func NewAuthHTTPHandler(
	service AuthService,
	googleConfig auth.GoogleClientAuthConfig,
) *AuthHTTPHandler {
	return &AuthHTTPHandler{
		service: service,
		GoogleOathConfig: &oauth2.Config{
			ClientID:     googleConfig.ClientID,
			ClientSecret: googleConfig.ClientSecret,
			RedirectURL:  googleConfig.RedirectURL,
			Scopes:       []string{stepsAddress},
			Endpoint:     google.Endpoint,
		},
	}
}
