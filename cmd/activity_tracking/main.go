package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"tg-echo-bot/golang_school/internal/core/config"
	"tg-echo-bot/golang_school/internal/core/db"
	"tg-echo-bot/golang_school/internal/feature/user"
)

func main() {
	// Задаем переменные окружения для envconfig
	os.Setenv("POSTGRES_HOST", "127.0.0.1")
	os.Setenv("POSTGRES_PORT", "5432")
	os.Setenv("POSTGRES_USER", "postgres")
	os.Setenv("POSTGRES_PASSWORD", "postgres")
	os.Setenv("POSTGRES_DB", "postgres")
	os.Setenv("POSTGRES_TIMEOUT", "5s")

	ctx := context.Background()

	// 1. Инициализация конфига
	cfg := config.NewPostgresConfigMust()

	// 2. Инициализация пула соединений pgxpool
	pool, err := db.New(ctx, cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL пулу: %v", err)
	}
	defer pool.Close()
	log.Println("Успешное подключение к PostgreSQL через pgxpool!")

	// 3. Создаем репозиторий через конструктор
	repo := user.NewRepository(pool)

	// 4. Создаем пользователя через конструктор NewUser
	goal := 10000
	newUser := user.NewUser("Глеб", 28, &goal, 2)

	err = repo.CreateUser(ctx, newUser)
	if err != nil {
		log.Fatalf("Ошибка создания пользователя: %v", err)
	}
	fmt.Printf("Пользователь успешно создан с ID: %d\n", newUser.ID)

	// 5. Проверяем получение пользователя
	foundUser, err := repo.GetByID(ctx, newUser.ID)
	if err != nil {
		log.Fatalf("Ошибка получения пользователя: %v", err)
	}
	fmt.Printf("Получен пользователь из базы: %+v\n", foundUser)
}
