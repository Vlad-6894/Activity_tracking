package user

import (
	"context"
	"fmt"
)

func (r *Repository) GetByID(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT id, full_name, age, google_refresh_token, steps_goal, rest_days, streak, created_at, updated_at
		FROM app.users
		WHERE id = $1;
	`

	u := &User{}
	err := r.pool.QueryRow(ctx, query, id).Scan(
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
		return nil, fmt.Errorf("get user by id %d: %w", id, err)
	}

	return u, nil
}
