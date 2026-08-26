package authRepositoryPostgres

import core_db "github.com/Vlad-6894/Activity_tracking/internal/core/db"

type AuthRepositoryPostgres struct {
	pool core_db.Pool
}

func NewAuthRepositoryPostgres(
	pool core_db.Pool,
) *AuthRepositoryPostgres {
	return &AuthRepositoryPostgres{
		pool: pool,
	}
}
