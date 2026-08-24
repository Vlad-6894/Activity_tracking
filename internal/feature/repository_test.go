package user_test

import (
	"context"
	"testing"
	"time"

	core_config "tg-echo-bot/golang_school/internal/core/config"
	core_db "tg-echo-bot/golang_school/internal/core/db"
	"tg-echo-bot/golang_school/internal/feature/user"
)

func TestUserRepository_CreateAndGet(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// 1. Конфигурация для локального тестового подключения
	cfg := core_config.PostgresConfig{
		Host:     "127.0.0.1",
		Port:     "5432",
		User:     "postgres",
		Password: "postgres",
		Database: "postgres",
		Timeout:  5 * time.Second,
	}

	pool, err := core_db.New(ctx, cfg)
	if err != nil {
		t.Fatalf("не удалось подключиться к базе: %v", err)
	}
	defer pool.Close()

	repo := user.NewRepository(pool)

	// 2. Тест создания пользователя через RegUser
	newUser := user.RegUser("Глеб Тестовый", 28, "mock_google_refresh_token_123", 10000, 2)

	createdUser, err := repo.CreateUser(ctx, newUser)
	if err != nil {
		t.Fatalf("ошибка CreateUser: %v", err)
	}

	if createdUser.ID == 0 {
		t.Errorf("ожидался сгенерированный ID > 0, получено 0")
	}

	t.Logf("Пользователь успешно создан в БД: %+v", createdUser)

	// 3. Тест получения пользователя через GetByID
	foundUser, err := repo.GetByID(ctx, createdUser.ID)
	if err != nil {
		t.Fatalf("ошибка GetByID: %v", err)
	}

	if foundUser.FullName != newUser.FullName {
		t.Errorf("ожидалось имя %s, получено %s", newUser.FullName, foundUser.FullName)
	}

	if foundUser.GoogleRefreshToken != newUser.GoogleRefreshToken {
		t.Errorf("ожидался токен %s, получен %s", newUser.GoogleRefreshToken, foundUser.GoogleRefreshToken)
	}

	t.Logf("Пользователь успешно прочитан из БД: %+v", foundUser)
}
