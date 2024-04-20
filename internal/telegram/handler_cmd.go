package telegram

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cmdStart = "start"
	cmdHelp  = "help"
	cmdLoad  = "load"
)
const (
	replyStart                    = "Привет! Это Facebook marketplace парсер бот"
	replyHelp                     = "Я оповещаю тебе о новых объявлениях на Facebook marketplace"
	replyUnknown                  = "Неизвестная команда("
	replyCancel                   = "Успешно отменено"
	replySetCity                  = "Для начала введите свой город"
	replySelectCity               = "Выберите свой город:"
	replySuccessfullySelectCity   = "Успешно выбран город!"
	replyWaitMsg                  = "Ты должен ввести название города!"
	replyError                    = "Произошла непредвидимая ошибка"
	replySelectRadius             = "Выберите радиус:"
	replySuccessfullySelectRadius = "Успешно выбран радиус"
	replySelectCategoty           = "Выберите категорию(можно выбрать одну категорию, если нужно мониторить несколько категорий, то создайте новый фильтр):"
	replyErr                      = "Произошла ошибка, подождите пару секунд..."
	replyNothingSelect            = `Вы не выбрали категорию, выберите категорию, а потом нажмите "подтвердить"`
)

var (
	SendFile  bool
	replyLoad = fmt.Sprintf("Пришлите фалй с расширением .txt. Данные в файлике должны быть в таком формате:\nloginfb:passwordfb\nipProxy:portProxy@loginProxy:passwordProxy\n[{%q: %q}, {%q: %q}]", "nameCookie1", "valueCookie1", "nameCookie2", "valueCookie2")
)

func (b *Bot) handleCommand(ctx context.Context, message *tgbotapi.Message) error {
	switch message.Command() {
	case cmdStart:
		return b.handleStartCommand(ctx, message)
	case cmdHelp:
		return b.handleHelpCommand(message)
	case cmdLoad:
		return b.handleLoadCommand(message)
	default:
		return b.handleUnknownCommand(message)
	}

}
func (b *Bot) handleStartCommand(ctx context.Context, message *tgbotapi.Message) error {
	err := b.db.AddChatID(ctx, int(message.Chat.ID))
	if err == database.ErrDublicateKey {
		log.Print("ErrDublicateKey")
	}
	if err != nil && err != database.ErrDublicateKey {
		return fmt.Errorf("add chat id error: %w", err)
	}
	msg := tgbotapi.NewMessage(message.Chat.ID, replyStart)
	msg.ReplyMarkup = StartKeyboard
	_, err = b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("[handleStartCommand]error send message: %w", err)
	}
	return nil
}

func (b *Bot) handleHelpCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, replyHelp)
	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("[handleHelpCommand]error send message: %w", err)
	}
	return nil
}

func (b *Bot) handleLoadCommand(message *tgbotapi.Message) error {
	for _, id := range Admins {
		if message.Chat.ID == id {
			msg := tgbotapi.NewMessage(message.Chat.ID, replyLoad)
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleHelpCommand]error send message: %w", err)
			}
			SendFile = true
		}
	}
	return nil
}
func (b *Bot) handleUnknownCommand(message *tgbotapi.Message) error {
	msg := tgbotapi.NewMessage(message.Chat.ID, replyUnknown)
	_, err := b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("[handleStartCommand]error send message: %w", err)
	}
	return nil
}
