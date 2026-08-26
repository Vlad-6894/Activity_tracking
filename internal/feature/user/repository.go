package user

import (
	"github.com/Vlad-6894/Activity_tracking/internal/core/db"
)

type Repository struct {
	pool db.Pool
}

func NewRepository(pool db.Pool) *Repository {
	return &Repository{pool: pool}
}
