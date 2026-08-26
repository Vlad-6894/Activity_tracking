package authTransport

import "context"

type GoogleAuthHandler struct {
	googleAuthService GoogleAuthService
}

type GoogleAuthService interface {
	Login(
		ctx context.Context,
		userID string,
	) (string, error)
	GetTokens(
		ctx context.Context,
		code string,
		state string,
	) error
}

func NewGoogleAuthHandler(
	googleAuthService GoogleAuthService,
) *GoogleAuthHandler {
	return &GoogleAuthHandler{
		googleAuthService: googleAuthService,
	}
}
