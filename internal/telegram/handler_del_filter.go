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

func (b *Bot) handleDeleteFilter(ctx context.Context, query *tgbotapi.CallbackQuery, cur_filter database.Filter) {
	log.Printf("!!!!!!filter: %v", cur_filter)
	log.Printf("!!!!!!filters: %v", Filters)
	log.Printf("!!!!!!query: %v", query)
	switch query.Data {
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
			return
		}
		var filters1 []database.Filter
		for ind, filter := range Filters {
			if filter.Id == strconv.Itoa(id) {
				filters1 = append(Filters[:ind], Filters[ind+1:]...)
			}
		}
		log.Printf("!!!!!!!!!!!!!!!!!filters: %v\nfilters1: %v", Filters, filters1)
		msg := tgbotapi.NewCallback(query.ID, "Успешно удалено!")
		_, err = b.bot.Request(msg)
		if err != nil {
			log.Fatalf("[handleMessage]error send message: %v", err)
		}
		if len(filters1) == 0 || filters1 == nil {
			msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Вы удалили все фильтры!")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			CurFilter = database.Filter{}
			return
		}
		editInlineKB := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID,
			fmt.Sprintf("Город: %s\nРадиус: %s\nКатегория: %s", filters1[0].City, filters1[0].Radius, filters1[0].Category),
			func() tgbotapi.InlineKeyboardMarkup {
				if len(filters1) == 1 {
					return SelectFilters3
				} else {
					return SelectFilters
				}
			}(),
		)
		_, err = b.bot.Send(editInlineKB)
		if err != nil {
			log.Printf("send message error: %v", err)
			msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}

			return
		}
		CurFilter = filters1[0]
	case "select_filter_previous":
		b.previousFilter(query, cur_filter, Filters)
	case "select_filter_next":
		b.nextFilter(query, cur_filter, Filters)
	}
}

func (b *Bot) previousFilter(query *tgbotapi.CallbackQuery, cur_filter database.Filter, filters []database.Filter) {
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
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}

				return
			}
			CurFilter = filters[ind-1]
		}
	}
}

func (b *Bot) nextFilter(query *tgbotapi.CallbackQuery, cur_filter database.Filter, filters []database.Filter) {
	for ind, filter := range filters {
		if cur_filter.Id == filter.Id {
			log.Printf("!!!!!ind: %d\nlen filter: %d", ind, len(filters))
			editInlineKB := tgbotapi.NewEditMessageTextAndMarkup(query.Message.Chat.ID, query.Message.MessageID,
				fmt.Sprintf("Город: %s\nРадиус: %s\nКатегория: %s", filters[ind+1].City, filters[ind+1].Radius, filters[ind+1].Category),
				func() tgbotapi.InlineKeyboardMarkup {
					if ind+1 == len(filters)-1 {
						return SelectFilters2
					} else {
						return SelectFilters1
					}
				}(),
			)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				log.Printf("send message error: %v", err)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}

				return
			}
			CurFilter = filters[ind+1]
		}
	}
}
