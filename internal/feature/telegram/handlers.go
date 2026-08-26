package telegram

import (
	"context"
	"fmt"
	"time"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"go.uber.org/zap"
)

const startText = `Привет!`

func (b *Bot) handleStart(ctx context.Context, api *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	b.send(ctx, api, update.Message.Chat.ID, startText, nil)
}

func (b *Bot) handlePing(ctx context.Context, api *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	uptime := time.Since(b.startedAt).Truncate(time.Second)
	b.send(ctx, api, update.Message.Chat.ID, fmt.Sprintf("pong -> uptime %s", uptime), nil)
}

func (b *Bot) handleFallback(ctx context.Context, api *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return
	}

	b.send(ctx, api, update.Message.Chat.ID, "Не знаю такой команды. Попробуй /start или /ping", nil)
}

func (b *Bot) send(ctx context.Context, api *bot.Bot, chatID int64, text string, markup models.ReplyMarkup) {
	_, err := api.SendMessage(ctx, &bot.SendMessageParams{
		ChatID:      chatID,
		Text:        text,
		ReplyMarkup: markup,
	})
	if err != nil {
		b.log.Warn("send message", zap.Int64("chat_id", chatID), zap.Error(err))
	}
}
