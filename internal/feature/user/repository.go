package user

import (
	core_db "github.com/Vlad-6894/Activity_tracking/internal/core/db"
)

type Repository struct {
	pool core_db.Pool
}

func NewRepository(pool core_db.Pool) *Repository {
	return &Repository{pool: pool}
}
