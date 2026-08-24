package user

import (
	core_db "tg-echo-bot/golang_school/internal/core/db"
)

type Repository struct {
	// [ИЗМЕНЕНО]: Используем core_db.Pool
	pool core_db.Pool
}

func NewRepository(pool core_db.Pool) *Repository {
	return &Repository{pool: pool}
}
