package authRepository

import (
	"github.com/Vlad-6894/Activity_tracking/internal/cache"
	authRepositoryRedis "github.com/Vlad-6894/Activity_tracking/internal/feature/google_authentication/repository/redis"
)

type AuthRepository struct {
	authRepositoryRedis.AuthRepositoryRedis
}

func NewAuthRepository(
	cache cache.Cache,
) *AuthRepository {
	return &AuthRepository{
		*authRepositoryRedis.NewAuthRepositoryRedis(cache),
	}
}
