package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleMake(message tgbotapi.Message) error {
	for {
		if MakeNext {

			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, MakeInlineKeyboard1)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			MakeNext = false
		} else if MakePrevious {
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, MakeInlineKeyboard)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			MakePrevious = false
		} else if MakeBreak {
			log.Print("exit make monitoring")
			break
		}
	}

	return nil
}
