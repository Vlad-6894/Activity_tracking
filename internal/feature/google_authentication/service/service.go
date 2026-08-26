package authService

import (
	"context"

	"golang.org/x/oauth2"
)

type GoogleAuthService struct {
	googleAuthRepository GoogleAuthRepository
}

type GoogleAuthRepository interface {
	SaveUUID(
		ctx context.Context,
		userID string,
		sessionUUID string,
	) error

	GetTelegramIDFromUUID(
		ctx context.Context,
		state string,
	) (string, error)

	SaveToken(
		ctx context.Context,
		userID string,
		token *oauth2.Token,
	) error
}

func NewGoogleAuthService(
	googleAuthRepository GoogleAuthRepository,
) *GoogleAuthService {
	return &GoogleAuthService{
		googleAuthRepository: googleAuthRepository,
	}
}
