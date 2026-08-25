package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"

	"github.com/Vlad-6894/Activity_tracking/internal/core/logger"
)

type Bot struct {
	api       *bot.Bot
	cfg       Config
	log       *logger.Logger
	startedAt time.Time
}

func New(cfg Config, log *logger.Logger) (*Bot, error) {
	b := &Bot{
		cfg:       cfg,
		log:       log,
		startedAt: time.Now(),
	}

	api, err := bot.New(
		cfg.BotToken,
		bot.WithDefaultHandler(b.handleFallback),
		bot.WithErrorsHandler(func(err error) {
			log.Warn("telegram update failed", zap.Error(err))
		}),
	)
	if err != nil {
		return nil, fmt.Errorf("create telegram bot: %w", err)
	}

	b.api = api

	api.RegisterHandler(bot.HandlerTypeMessageText, "start", bot.MatchTypeCommandStartOnly, b.handleStart)
	api.RegisterHandler(bot.HandlerTypeMessageText, "ping", bot.MatchTypeCommandStartOnly, b.handlePing)

	return b, nil
}

func (b *Bot) Run(ctx context.Context) error {
	me, err := b.api.GetMe(ctx)
	if err != nil {
		return fmt.Errorf("telegram getMe: %w", err)
	}

	b.log.Info("bot authorized",
		zap.String("bot_username", "@"+me.Username),
		zap.Int64("bot_id", me.ID),
	)
	b.log.Info("webapp url", zap.String("url", b.cfg.WebAppURL))

	if err := b.registerCommands(ctx); err != nil {
		return err
	}

	if err := b.registerMenuButton(ctx); err != nil {
		return err
	}

	b.log.Info("bot polling started")
	b.api.Start(ctx)
	b.log.Info("bot stopped")

	return nil
}

func (b *Bot) registerCommands(ctx context.Context) error {
	_, err := b.api.SetMyCommands(ctx, &bot.SetMyCommandsParams{
		Commands: []models.BotCommand{
			{Command: "start", Description: "Просто 'Привет'"},
			{Command: "ping", Description: "Проверить, что бот жив"},
		},
	})
	if err != nil {
		return fmt.Errorf("set bot commands: %w", err)
	}

	return nil
}

func (b *Bot) registerMenuButton(ctx context.Context) error {
	_, err := b.api.SetChatMenuButton(ctx, &bot.SetChatMenuButtonParams{
		MenuButton: models.MenuButtonWebApp{
			Type:   models.MenuButtonTypeWebApp,
			Text:   "Open",
			WebApp: models.WebAppInfo{URL: b.cfg.WebAppURL},
		},
	})
	if err != nil {
		return fmt.Errorf("set chat menu button: %w", err)
	}

	return nil
}
