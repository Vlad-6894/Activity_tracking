package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{db: db}
}

// CreateUser создает нового пользователя в схеме app.users
func (r *Repository) CreateUser(ctx context.Context, u *User) error {
	query := `
		INSERT INTO app.users (
			full_name, 
			age, 
			google_refresh_token, 
			steps_goal, 
			rest_days, 
			streak, 
			created_at
		) 
		VALUES ($1, $2, $3, $4, $5, $6, NOW())
		RETURNING id, created_at;
	`

	err := r.db.QueryRowContext(
		ctx,
		query,
		u.FullName,
		u.Age,
		u.GoogleRefreshToken,
		u.StepsGoal,
		u.RestDays,
		u.Streak,
	).Scan(&u.ID, &u.CreatedAt)

	if err != nil {
		return fmt.Errorf("ошибка создания пользователя: %w", err)
	}

	return nil
}

// GetByID возвращает пользователя по первичному ключу ID
func (r *Repository) GetByID(ctx context.Context, id int64) (*User, error) {
	query := `
		SELECT 
			id, 
			full_name, 
			age, 
			google_refresh_token, 
			steps_goal, 
			rest_days, 
			streak, 
			created_at, 
			updated_at
		FROM app.users
		WHERE id = $1;
	`

	var u User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, fmt.Errorf("ошибка получения пользователя: %w", err)
	}

	return &u, nil
}
