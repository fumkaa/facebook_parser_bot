package telegram

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"facebook_marketplace_bot/internal/parser"
	"fmt"
	"log"
	"os"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

// var isDel = make(chan database.Filter)
// var filterFile = make(chan string)

func (b *Bot) handleDeleteFilter(ctx context.Context, query *tgbotapi.CallbackQuery, cur_filter database.Filter, filters []database.Filter) {
	log.Printf("!!!!!!filter: %v", cur_filter)
	log.Printf("!!!!!!query: %v", query)
	switch query.Data {
	// case "filter_id_" + filter.Id:
	// 	go func() {
	// 		filterFile <- filter.Filter_file
	// 	}()
	// 	id, err := strconv.Atoi(filter.Id)
	// 	if err != nil {
	// 		log.Fatal("err")
	// 	}
	// 	if err := b.db.DeleteFilter(ctx, id); err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	go func() {
	// 		isDel <- filter
	// 	}()
	case "select_filter_delete":
		id, err := strconv.Atoi(cur_filter.Id)
		if err != nil {
			log.Fatal("err")
		}
		if err := b.db.DeleteFilter(ctx, id); err != nil {
			log.Printf("SelectAllFilter error: %v", err)
			msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			if err := b.db.DeleteWaitMessage(ctx, int(query.Message.Chat.ID)); err != nil {
				log.Fatalf("delete wait message error: %v", err)
			}
			SelectCity = false
			SelectRadius = false
			SelectInlineKB = false
			InputPrice = false
			InputYear = false
			InputMill = false
			return
		}
		if err := os.Rename(parser.Work_account+cur_filter.Filter_file, parser.Free_account+cur_filter.Filter_file); err != nil {
			log.Printf("Rename error: %v", err)
			msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова удалить фильтр")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			if err := b.db.DeleteWaitMessage(ctx, int(query.Message.Chat.ID)); err != nil {
				log.Fatalf("delete wait message error: %v", err)
			}
			SelectCity = false
			SelectRadius = false
			SelectInlineKB = false
			InputPrice = false
			InputYear = false
			InputMill = false
			return
		}
		msg := tgbotapi.NewCallback(query.ID, "Успешно удалено!")
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Fatalf("[handleMessage]error send message: %v", err)
		}
		b.previousFilter(ctx, query, cur_filter, filters)
	case "select_filter_previous":
		b.previousFilter(ctx, query, cur_filter, filters)
	case "select_filter_next":
		for ind, filter := range filters {
			if cur_filter.Id == filter.Id {
				editInlineKB := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID,
					fmt.Sprintf("Город: %s\nРадиус: %s\nКатегория: %s", filters[ind+1].City, filters[ind+1].Radius, filters[ind+1].Category),
					func() tgbotapi.InlineKeyboardMarkup {
						if ind+1 == len(filters) {
							return SelectFilters2
						} else {
							return SelectFilters1
						}
					}(),
				)
				_, err := b.bot.Send(editInlineKB)
				if err != nil {
					log.Printf("send message error: %v", err)
					msg := tgbotapi.NewMessage(query.From.ID, "Произошла ошибка попробуйте снова")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}
					if err := b.db.DeleteWaitMessage(ctx, int(query.Message.Chat.ID)); err != nil {
						log.Fatalf("delete wait message error: %v", err)
					}
					SelectCity = false
					SelectRadius = false
					SelectInlineKB = false
					InputPrice = false
					InputYear = false
					InputMill = false
					return
				}
				CurFilter = filters[ind+1]
			}
		}
	}
}

func (b *Bot) previousFilter(ctx context.Context, query *tgbotapi.CallbackQuery, cur_filter database.Filter, filters []database.Filter) {
	for ind, filter := range filters {
		if cur_filter.Id == filter.Id {
			editInlineKB := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID,
				fmt.Sprintf("Город: %s\nРадиус: %s\nКатегория: %s", filters[ind-1].City, filters[ind-1].Radius, filters[ind-1].Category),
				func() tgbotapi.InlineKeyboardMarkup {
					if ind-1 == 0 {
						return SelectFilters
					} else {
						return SelectFilters1
					}
				}(),
			)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				log.Printf("send message error: %v", err)
				msg := tgbotapi.NewMessage(query.From.ID, "Произошла ошибка попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				if err := b.db.DeleteWaitMessage(ctx, int(query.Message.Chat.ID)); err != nil {
					log.Fatalf("delete wait message error: %v", err)
				}
				SelectCity = false
				SelectRadius = false
				SelectInlineKB = false
				InputPrice = false
				InputYear = false
				InputMill = false
				return
			}
			CurFilter = filters[ind-1]
		}
	}
}
