package telegram

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"facebook_marketplace_bot/internal/parser"
	"fmt"
	"log"
	"sync"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Admins = []int64{1295202595, 5398451124}
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	db     database.Database
	parser parser.Parser
	rw     sync.RWMutex
}

func NewBot(bot *tgbotapi.BotAPI, db database.Database, parser parser.Parser) Bot {
	return Bot{
		bot:    bot,
		db:     db,
		parser: parser,
	}
}

func (b *Bot) Start(ctx context.Context) error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	updates := b.initUpdatesChan()
	if err := b.handleUpdates(ctx, updates); err != nil {
		return fmt.Errorf("hadle updates error: %w", err)
	}
	return nil
}

func (b *Bot) initUpdatesChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) error {
	for update := range updates {
		if update.Message != nil {
			if update.Message.IsCommand() {
				if err := b.handleCommand(ctx, update.Message); err != nil {
					return fmt.Errorf("handle command error: %w", err)
				}
				continue
			}
			if SendFile {
				if update.Message.Document != nil {
					if err := b.handleFile(update.Message); err != nil {
						return fmt.Errorf("handle file error: %w", err)
					}
					continue
				}
			}
			if err := b.handleMessage(ctx, update.Message); err != nil {
				return fmt.Errorf("handle message error: %w", err)
			}
		} else if update.CallbackQuery != nil {
			log.Print("update callback query")
			b.handleDeleteFilter(ctx, update.CallbackQuery, CurFilter, Filters)
			if SelectCity {
				ChSelectCity <- update.CallbackQuery
				log.Print("send update.CallbackQuery to ChSelectCity")
			}
			if SelectRadius {
				ChSelectRadius <- update.CallbackQuery
				log.Print("send update.CallbackQuery to ChSelectRadius")
			}
			if SelectInlineKB {
				log.Print("update.CallbackQuery SelectInlineKB")
				if err := b.handlerCategoryInlineKeyboard(ctx, ID, update.CallbackQuery, SelectCategory); err != nil {
					return fmt.Errorf("handlerCategoryInlineKeyboard error: %w", err)
				}
			}
		}

	}
	return nil
}
