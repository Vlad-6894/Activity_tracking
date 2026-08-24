package user

import (
	"context"
	"fmt"
)

func (r *Repository) CreateUser(ctx context.Context, u User) (User, error) {
	query := `
		INSERT INTO app.users (
			full_name,
			tg_user_name,
			age,
			google_refresh_token,
			steps_goal,
			rest_days,
			streak,
			created_at
		)
		VALUES ($1, $2, $3, $4, $5, $6, $7, NOW())
		RETURNING id, full_name, tg_user_name, age, google_refresh_token, steps_goal, rest_days, streak, created_at;
	`

	var created User
	err := r.pool.QueryRow(
		ctx,
		query,
		u.FullName,
		u.TelegramUserName,
		u.Age,
		u.GoogleRefreshToken,
		u.StepsGoal,
		u.RestDays,
		u.Streak,
	).Scan(
		&created.ID,
		&created.FullName,
		&created.TelegramUserName,
		&created.Age,
		&created.GoogleRefreshToken,
		&created.StepsGoal,
		&created.RestDays,
		&created.Streak,
		&created.CreatedAt,
	)

	if err != nil {
		return User{}, fmt.Errorf("create user: %w", err)
	}

	return created, nil
}
