package authRepositoryRedis

import "github.com/Vlad-6894/Activity_tracking/internal/cache"

type AuthRepositoryRedis struct {
	cache cache.Cache
}

func NewAuthRepositoryRedis(
	cache cache.Cache,
) *AuthRepositoryRedis {
	return &AuthRepositoryRedis{
		cache: cache,
	}
}
