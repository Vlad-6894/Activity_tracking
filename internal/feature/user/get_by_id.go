package user

import (
	"context"
	"fmt"
)

func (r *Repository) GetByID(ctx context.Context, id int) (User, error) {
	query := `
		SELECT 
			id,
			full_name,
			tg_user_name,
			age,
			google_refresh_token,
			steps_goal,
			rest_days,
			streak,
			created_at
		FROM app.users
		WHERE id = $1;
	`

	var u User
	err := r.pool.QueryRow(ctx, query, id).Scan(
		&u.ID,
		&u.FullName,
		&u.TelegramUserName,
		&u.Age,
		&u.GoogleRefreshToken,
		&u.StepsGoal,
		&u.RestDays,
		&u.Streak,
		&u.CreatedAt,
	)

	if err != nil {
		return User{}, fmt.Errorf("get user by id: %w", err)
	}

	return u, nil
}
