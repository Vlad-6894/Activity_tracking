package core_logger

import (
	"log/slog"
	"os"
)

// [ИЗМЕНЕНО]: Функция инициализации структурированного slog-логгера
func Init() *slog.Logger {
	opts := &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}

	handler := slog.NewJSONHandler(os.Stdout, opts)
	logger := slog.New(handler)
	slog.SetDefault(logger)

	return logger
}
