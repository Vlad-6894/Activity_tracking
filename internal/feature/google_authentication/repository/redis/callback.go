package authRepositoryRedis

import (
	"context"
	"fmt"

	core_errors "github.com/Vlad-6894/Activity_tracking/internal/core/errors"
)

func (r *AuthRepositoryRedis) GetTelegramIDFromUUID(
	ctx context.Context,
	state string,
) (string, error) {
	userID, err := r.cache.Get(ctx, state)
	if err != nil {
		return "", fmt.Errorf(
			"get userID from cache: %v: %w",
			err,
			core_errors.ErrNotFound,
		)
	}

	return string(userID), nil
}
