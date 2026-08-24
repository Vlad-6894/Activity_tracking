package user

import "tg-echo-bot/golang_school/internal/core/db"

type Repository struct {
	pool db.Pool
}

func NewRepository(pool db.Pool) *Repository {
	return &Repository{
		pool: pool,
	}
}
