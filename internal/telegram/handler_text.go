package telegram

import (
	database "facebook_marketplace_bot/internal/database/migration"
	"fmt"
	"io"
	"net/http"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	ID int
	// url1             string
	CurFilter database.Filter
	Filters   []database.Filter
	// SelectCity       bool
	// SelectRadius     bool

	// SelectInlineKB   bool
	// InputPrice       bool
	// InputYear        bool

	// InputMet         bool
	// InputMill        bool
	SelectCategory   tgbotapi.Message
	ChInputPrice     = make(chan *tgbotapi.CallbackQuery)
	ChInputMMPrice   = make(chan int)
	ChInputYear      = make(chan *tgbotapi.CallbackQuery)
	ChInputMill      = make(chan *tgbotapi.CallbackQuery)
	ChSelectCategory = make(chan *tgbotapi.CallbackQuery)
	ChSelectCity     = make(chan *tgbotapi.CallbackQuery)
	ChSelectRadius   = make(chan *tgbotapi.CallbackQuery)
	ChMaxPrice       = make(chan string)
	ChMinPrice       = make(chan string)
	ChMaxYear        = make(chan string)
	ChMinYear        = make(chan string)
	ChMaxMill        = make(chan string)
	ChMinMill        = make(chan string)
	ChUrl            = make(chan string)
	ChMaxMet         = make(chan string)
	ChMinMet         = make(chan string)
	ChFilter         = make(chan database.Filter)
)

func (b *Bot) handleFile(message *tgbotapi.Message) error {

	url, err := b.bot.GetFileDirectURL(message.Document.FileID)
	if err != nil {
		return fmt.Errorf("[handleFile]error GetFile : %w", err)
	}
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("[handleFile]error get : %w", err)
	}
	res, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("[handleFile]error read all : %w", err)
	}
	file, err := os.Create("data/free_account/" + message.Document.FileName)
	if err != nil {
		return fmt.Errorf("[handleFile]error create : %w", err)
	}
	defer file.Close()
	if _, err = file.Write(res); err != nil {
		return fmt.Errorf("[handleFile]error write : %w", err)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, "Файл успешно сохранен!")
	msg.ReplyMarkup = StartKeyboard
	_, err = b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("[handleFile]error send message: %w", err)
	}

	return nil
}
