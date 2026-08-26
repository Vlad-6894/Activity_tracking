package authService

import (
	"context"
	"fmt"

	core_errors "github.com/Vlad-6894/Activity_tracking/internal/core/errors"
)

func (s *GoogleAuthService) GetTokens(
	ctx context.Context,
	code string,
	state string,
) error {
	userID, err := s.googleAuthRepository.GetTelegramIDFromUUID(ctx, state)
	if err != nil {
		return fmt.Errorf("get userID from cache: %v: %w", err, core_errors.ErrNotFound)
	}

	config, err := NewConfig()
	if err != nil {
		return fmt.Errorf("create config: %w", err)
	}

	token, err := config.Exchange(ctx, code)
	if err != nil {
		return fmt.Errorf("exchange code: %w", err)
	}

	if err := s.googleAuthRepository.SaveToken(ctx, userID, token); err != nil {
		return fmt.Errorf("save token: %w", err)
	}

	//maybe create delete state from cache?

	return nil
}
