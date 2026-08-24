package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	"tg-echo-bot/golang_school/internal/core/config"
	"tg-echo-bot/golang_school/internal/core/db"
	"tg-echo-bot/golang_school/internal/feature/user"
)

func main() {
	// Контекст с отслеживанием системных сигналов завершения (Ctrl+C, SIGTERM)
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 1. Инициализация конфига
	cfg := config.NewPostgresConfigMust()

	// 2. Инициализация пула соединений pgxpool
	pool, err := db.New(ctx, cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL пулу: %v", err)
	}
	defer pool.Close()
	log.Println("Успешное подключение к PostgreSQL через pgxpool!")

	// 3. Создаем репозиторий через конструктор (принимает интерфейс db.Pool)
	repo := user.NewRepository(pool)

	// 4. Создаем пользователя через конструктор регистрации RegUser
	goal := 10000
	newUser := user.RegUser("Глеб", 28, &goal, 2)

	err = repo.CreateUser(ctx, newUser)
	if err != nil {
		log.Fatalf("Ошибка создания пользователя: %v", err)
	}
	fmt.Printf("Пользователь успешно создан: %+v\n", newUser)

	// 5. Проверяем получение пользователя
	foundUser, err := repo.GetByID(ctx, newUser.ID)
	if err != nil {
		log.Fatalf("Ошибка получения пользователя: %v", err)
	}
	fmt.Printf("Получен пользователь из базы: %+v\n", foundUser)
}
