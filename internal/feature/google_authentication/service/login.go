package authService

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"golang.org/x/oauth2"
)

func (s *GoogleAuthService) Login(
	ctx context.Context,
	userID string,
) (string, error) {
	sessionUUID := uuid.New().String()
	if err := s.googleAuthRepository.SaveUUID(ctx, userID, sessionUUID); err != nil {
		return "", fmt.Errorf("save to cash: %w", err)
	}

	config, err := NewConfig()
	if err != nil {
		return "", fmt.Errorf("failed to create config: %w", err)
	}

	authURL := config.AuthCodeURL(sessionUUID, oauth2.AccessTypeOffline)
	return authURL, nil
}
