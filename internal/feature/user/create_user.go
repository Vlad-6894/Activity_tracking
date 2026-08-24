package user

import (
	"context"
	"fmt"
	"time"
)

func (r *Repository) CreateUser(ctx context.Context, u *User) error {
	query := `
		INSERT INTO app.users (full_name, age, google_refresh_token, steps_goal, rest_days, streak, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING id, created_at;
	`

	now := time.Now().UTC()
	err := r.pool.QueryRow(ctx, query,
		u.FullName,
		u.Age,
		u.GoogleRefreshToken,
		u.StepsGoal,
		u.RestDays,
		u.Streak,
		now,
	).Scan(&u.ID, &u.CreatedAt)

	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
