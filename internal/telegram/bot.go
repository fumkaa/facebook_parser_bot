package telegram

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"facebook_marketplace_bot/internal/parser"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"

	"github.com/chromedp/chromedp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/looplab/fsm"
)

var (
	Admins = []int64{1295202595, 5398451124}
)

// fsm state
var (
	state_base                = "start"
	state_cancel              = "cancel"
	state_set_city            = "set_city"
	state_set_radius          = "set_radius"
	state_select_category     = "select_category"
	state_input_price         = "input_price"
	state_input_square_meters = "input_square_meters"
	state_input_year          = "input_year"
	state_input_mill          = "input_mill"
	// state_my_filter           = "my_filter"
)

type Bot struct {
	bot    *tgbotapi.BotAPI
	db     database.Database
	parser parser.Parser
	// rw     sync.RWMutex
	FSM *fsm.FSM
}

func NewBot(bot *tgbotapi.BotAPI, db database.Database, parser parser.Parser) Bot {
	return Bot{
		bot:    bot,
		db:     db,
		parser: parser,
		FSM: fsm.NewFSM(
			"run",
			fsm.Events{
				// add filter
				{Name: state_base, Src: []string{"run", state_input_square_meters, state_input_year, state_input_mill, state_set_city, state_set_radius, state_select_category, state_input_price}, Dst: state_base},
				{Name: state_cancel, Src: []string{state_input_price,
					state_input_square_meters, state_input_year, state_input_mill, state_set_city, state_set_radius, state_select_category},
					Dst: state_base},

				{Name: state_set_city, Src: []string{state_base}, Dst: state_set_city},
				{Name: state_set_radius, Src: []string{state_set_city}, Dst: state_set_radius},
				{Name: state_select_category, Src: []string{state_set_radius, state_input_price, state_input_square_meters, state_input_year, state_input_mill}, Dst: state_select_category},

				{Name: state_input_price, Src: []string{state_select_category}, Dst: state_input_price},
				{Name: state_input_square_meters, Src: []string{state_select_category}, Dst: state_input_square_meters},
				{Name: state_input_year, Src: []string{state_select_category}, Dst: state_input_year},
				{Name: state_input_mill, Src: []string{state_select_category}, Dst: state_input_mill},

				// my filter
				// {Name: "my_filter", Src: []string{"start"}, Dst: "my_filter"},
			},
			fsm.Callbacks{},
		),
	}
}

func (b *Bot) Start(ctx context.Context) error {
	log.Printf("Authorized on account %s", b.bot.Self.UserName)
	if err := b.FSM.Event(ctx, state_base); err != nil {
		return fmt.Errorf("fsm event error: %w", err)
	}
	isEmpty, err := IsEmpty(parser.Work_account)
	if err != nil {
		return fmt.Errorf("check is empty error: %w", err)
	}
	if !isEmpty {
		workFile, err := os.ReadDir(parser.Work_account)
		if err != nil {
			return fmt.Errorf("read dir work account error: %w", err)
		}
		for _, file := range workFile {
			dataFilter, err := b.db.SelectFilterToFilterFile(ctx, file.Name())
			if err != nil {
				return fmt.Errorf("SelectFilterToFilterFile error: %w", err)
			}
			data, err := b.parser.GetDataFile(file.Name())
			if err != nil {
				return fmt.Errorf("GetDataFile error: %w", err)
			}
			proxy := fmt.Sprintf("http://%s", data.Datas.IpPortPX)
			log.Printf("proxy: %s", proxy)
			opts := append(chromedp.DefaultExecAllocatorOptions[:],
				chromedp.ProxyServer(proxy),
				chromedp.WindowSize(1900, 1080), // init with a desktop view
				chromedp.Flag("enable-automation", false),
				// chromedp.Flag("headless", false),
			)
			log.Printf("ip port proxy: %v", data.Datas.IpPortPX)
			var ctxChr context.Context
			ctxChr, Cancel1 = chromedp.NewExecAllocator(context.Background(), opts...)
			Ctxt, Cancel2 = chromedp.NewContext(ctxChr) // chromedp.WithDebugf(log.Printf),
			log.Print("!!!settings!!!")
			err = b.parser.Settings(Ctxt, data)
			if err != nil {
				return fmt.Errorf("settings error: %w", err)
			}
			id, err := strconv.Atoi(dataFilter.Id)
			if err != nil {
				return fmt.Errorf("convert to int dataFilter.Id error: %w", err)
			}
			go func() { b.monitoring(ctx, int64(dataFilter.Chat_id), dataFilter.Monitoring, id) }()
		}

	}
	updates := b.initUpdatesChan()
	if err := b.handleUpdates(ctx, updates); err != nil {
		return fmt.Errorf("hadle updates error: %w", err)
	}
	return nil
}
func IsEmpty(name string) (bool, error) {
	f, err := os.Open(name)
	if err != nil {
		return false, fmt.Errorf("open file error: %w", err)
	}
	defer f.Close()

	_, err = f.Readdirnames(1) // Or f.Readdir(1)
	if err == io.EOF {
		return true, nil
	}
	return false, nil
}

func (b *Bot) initUpdatesChan() tgbotapi.UpdatesChannel {
	u := tgbotapi.NewUpdate(0)
	u.Timeout = 30

	return b.bot.GetUpdatesChan(u)
}

func (b *Bot) handleUpdates(ctx context.Context, updates tgbotapi.UpdatesChannel) error {
	for update := range updates {

		log.Printf("current state: %s", b.FSM.Current())
		if update.Message != nil {
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
					msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					return fmt.Errorf("handle command error: %w", err)
				}
				continue
			}
			if SendFile {
				if update.Message.Document != nil {
					if err := b.handleFile(update.Message); err != nil {
						msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Printf("error send message: %v", err)
						}
						return fmt.Errorf("handle file error: %w", err)
					}
					continue
				}
			}
			if err := b.updateHandleMessage(ctx, update.Message); err != nil {
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				return fmt.Errorf("handle message error: %w", err)
			}
		} else if update.CallbackQuery != nil {
			log.Print("update callback query")
			var err error
			b.handleDeleteFilter(ctx, update.CallbackQuery, CurFilter)
			switch b.FSM.Current() {
			case state_set_city:
				ID, err = b.db.AddChatIDFilters(ctx, int(update.CallbackQuery.Message.Chat.ID))
				if err != nil {
					log.Printf("AddChatIDFilters error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					return nil
				}
				if err := b.handlerCityInlineKeyboard(Ctxt, update.CallbackQuery); err != nil {
					log.Printf("handlerCityInlineKeyboard error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, replyErr)
					b.bot.Send(msg)
					continue
				}
				if err := b.db.AddCityFilter(ctx, ID, update.CallbackQuery.Data); err != nil {
					log.Printf("AddCityFilter error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					return nil
				}
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, replySuccessfullySelectCity)
				if _, err := b.bot.Send(msg); err != nil {
					log.Printf("error send message: %v", err)
				}
				if err := b.FSM.Event(ctx, state_set_radius); err != nil {
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					return nil
				}
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, replySelectRadius)
				msg.ReplyMarkup = RadiusInlineKeyboard
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
					return nil
				}
				log.Print("wait click inline keyboard")
			case state_set_radius:
				if err := b.handlerRadiusInlineKeyboard(Ctxt, update.CallbackQuery); err != nil {
					log.Printf("handlerRadiusInlineKeyboard error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
				}
				if err := b.db.AddRadiusFilter(ctx, ID, update.CallbackQuery.Data); err != nil {
					log.Printf("AddRadiusFilter error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					return nil
				}
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, replySuccessfullySelectRadius)
				if _, err := b.bot.Send(msg); err != nil {
					log.Printf("error send message: %v", err)
					continue
				}

				if err := b.FSM.Event(ctx, state_select_category); err != nil {
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					return nil
				}
				msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, replySelectCategoty)
				msg.ReplyMarkup = CategoryInlineKeyboard1
				sendmsg, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
					continue
				}
				SelectCategory = sendmsg
			case state_select_category:
				log.Print("update.CallbackQuery SelectInlineKB")
				if err := b.handlerCategoryInlineKeyboard(ctx, update.CallbackQuery); err != nil {
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, fmt.Sprintf("произошла ошибка: %v", err))
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					return fmt.Errorf("handlerCategoryInlineKeyboard error: %w", err)
				}
			}
		}
	}
	return nil
}
