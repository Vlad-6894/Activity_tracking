package user

import (
	"context"
	"fmt"
)

func (r *Repository) CreateUser(ctx context.Context, u *User) error {
	query := `
		INSERT INTO app.users (full_name, age, google_refresh_token, steps_goal, rest_days, streak, created_at)
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id, full_name, age, google_refresh_token, steps_goal, rest_days, streak, created_at, updated_at;
	`

	err := r.pool.QueryRow(ctx, query,
		u.FullName,
		u.Age,
		u.GoogleRefreshToken,
		u.StepsGoal,
		u.RestDays,
		u.Streak,
	).Scan(
		&u.ID,
		&u.FullName,
		&u.Age,
		&u.GoogleRefreshToken,
		&u.StepsGoal,
		&u.RestDays,
		&u.Streak,
		&u.CreatedAt,
		&u.UpdatedAt,
	)

	if err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}
