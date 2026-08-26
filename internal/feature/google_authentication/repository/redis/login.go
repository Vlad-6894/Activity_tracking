package authRepositoryRedis

import (
	"context"
	"fmt"
)

func (r *AuthRepositoryRedis) SaveUUID(
	ctx context.Context,
	userID string,
	sessionUUID string,
) error {
	err := r.cache.Set(
		ctx,
		userID,
		sessionUUID,
		r.cache.OpTimeout(),
	)
	if err != nil {
		return fmt.Errorf("save session uuid for user %s: %w", userID, err)
	}

	return nil
}
