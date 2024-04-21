package telegram

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var isDel = make(chan database.Filter)
var filterFile = make(chan string)

func (b *Bot) handleDeleteFilter(ctx context.Context, query *tgbotapi.CallbackQuery, filter database.Filter) {
	log.Printf("!!!!!!filter: %v", filter)
	log.Printf("!!!!!!query: %v", query)
	switch query.Data {
	case "filter_id_" + filter.Id:
		go func() {
			filterFile <- filter.Filter_file
		}()
		id, err := strconv.Atoi(filter.Id)
		if err != nil {
			log.Fatal("err")
		}
		if err := b.db.DeleteFilter(ctx, id); err != nil {
			log.Fatal(err)
		}
		go func() {
			isDel <- filter
		}()
	}
	DeleteFilter = false
}
