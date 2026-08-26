package authRepositoryPostgres

import (
	"context"
	"fmt"

	"golang.org/x/oauth2"
)

func (r AuthRepositoryPostgres) SaveToken(
	ctx context.Context,
	userID string,
	token *oauth2.Token,
) error {
	ctxWithTime, cancel := context.WithTimeout(ctx, r.pool.GetTimeot())
	defer cancel()

	query := `
		UPDATE app.users 
		SET google_refresh_token = $1
		WHERE id = $2;
	`

	if _, err := r.pool.Exec(ctxWithTime, query, token.RefreshToken, userID); err != nil {
		return fmt.Errorf("fail to save refresh token: %w", err)
	}

	return nil
}
