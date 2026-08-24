package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	core_config "github.com/Vlad-6894/Activity_tracking/internal/core/config"
	core_db "github.com/Vlad-6894/Activity_tracking/internal/core/db"
	core_logger "github.com/Vlad-6894/Activity_tracking/internal/core/logger"
	"go.uber.org/zap"
)

func main() {
	logger, err := core_logger.NewLogger(core_logger.NewLoggerConfigMust())
	if err != nil {
		fmt.Println("init logger error:", err)
		os.Exit(1)
	}
	defer logger.Close()

	logger.Info("Логер успешно инициализирован")

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	cfg := core_config.NewPostgresConfigMust()

	pool, err := core_db.New(ctx, cfg)
	if err != nil {
		logger.Error("Ошибка подключения к PostgreSQL пулу", zap.Error(err))
		os.Exit(1)
	}
	defer pool.Close()

	logger.Info("Успешное подключение к PostgreSQL через pgxpool!")
}
