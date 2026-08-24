package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"

	// [ИЗМЕНЕНО]: Обновлены пути к переименованным пакетам core_config и core_db
	core_config "tg-echo-bot/golang_school/internal/core/config"
	core_db "tg-echo-bot/golang_school/internal/core/db"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	// 1. Инициализация конфига
	cfg := core_config.NewPostgresConfigMust()

	// 2. Инициализация пула соединений
	pool, err := core_db.New(ctx, cfg)
	if err != nil {
		log.Fatalf("Ошибка подключения к PostgreSQL пулу: %v", err)
	}
	defer pool.Close()
	log.Println("Успешное подключение к PostgreSQL через pgxpool!")

	// [ИЗМЕНЕНО]: Все тестовые создания пользователя удалены из main.go по замечанию Влада
}
