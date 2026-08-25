package main

import (
	"github.com/Vlad-6894/Activity_tracking/internal/core/db"
	"github.com/Vlad-6894/Activity_tracking/internal/core/logger"
	"github.com/Vlad-6894/Activity_tracking/internal/core/transport/http/server"
	"github.com/Vlad-6894/Activity_tracking/internal/feature/telegram"
)

type config struct {
	logger   logger.Config
	db       db.Config
	telegram telegram.Config
	http     server.Config
}
