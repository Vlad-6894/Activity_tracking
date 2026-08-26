package main

import (
	"context"
	"errors"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/Vlad-6894/Activity_tracking/internal/core/db"
	"github.com/Vlad-6894/Activity_tracking/internal/core/logger"
	"github.com/Vlad-6894/Activity_tracking/internal/core/transport/http/server"
	"github.com/Vlad-6894/Activity_tracking/internal/feature/miniapp"
	"github.com/Vlad-6894/Activity_tracking/internal/feature/telegram"
	"golang.org/x/sync/errgroup"
)

func run() error {
	cfg, err := loadConfig()
	if err != nil {
		return err
	}

	log, err := logger.New(cfg.logger)
	if err != nil {
		return fmt.Errorf("init logger: %w", err)
	}
	defer log.Close()

	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	pool, err := db.New(ctx, cfg.db)
	if err != nil {
		return fmt.Errorf("connect postgres: %w", err)
	}
	defer pool.Close()

	log.Info("postgres connected")

	bot, err := telegram.New(cfg.telegram, log)
	if err != nil {
		return fmt.Errorf("init telegram bot: %w", err)
	}

	miniappHandler := miniapp.NewHandler(cfg.telegram.BotToken, log)

	srv := server.New(cfg.http, log)

	apiV1 := server.NewAPIVersionRouter(server.APIVersion1)
	apiV1.AddRoutes(miniappHandler.Routes()...)

	srv.RegisterAPIRouters(apiV1)
	srv.RegisterRoutes(miniappHandler.RootRoutes()...)

	return runServices(ctx, bot, srv)
}

func loadConfig() (config, error) {
	var (
		cfg  config
		errs []error
		err  error
	)

	if cfg.logger, err = logger.NewConfig(); err != nil {
		errs = append(errs, err)
	}

	if cfg.db, err = db.NewConfig(); err != nil {
		errs = append(errs, err)
	}

	if cfg.telegram, err = telegram.NewConfig(); err != nil {
		errs = append(errs, err)
	}

	if cfg.http, err = server.NewConfig(); err != nil {
		errs = append(errs, err)
	}

	if err := errors.Join(errs...); err != nil {
		return config{}, fmt.Errorf("load config: %w", err)
	}

	return cfg, nil
}

func runServices(ctx context.Context, bot *telegram.Bot, srv *server.Server) error {
	grp, grpCtx := errgroup.WithContext(ctx)

	grp.Go(func() error {
		return bot.Run(grpCtx)
	})

	grp.Go(func() error {
		return srv.Run(grpCtx)
	})

	if err := grp.Wait(); err != nil {
		return fmt.Errorf("services stopped %w", err)
	}

	return nil
}
