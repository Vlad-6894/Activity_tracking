package authRepositoryPostgres

import (
	"context"

	"golang.org/x/oauth2"
)

func (r AuthRepositoryPostgres) SaveToken(
	ctx context.Context,
	userID string,
	token *oauth2.Token,
) error {
	query := `
		
	`
}
