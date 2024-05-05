package telegram

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"facebook_marketplace_bot/internal/parser"
	"fmt"
	"log"

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
				{Name: state_base, Src: []string{"run", state_set_city}, Dst: state_base},
				{Name: state_cancel, Src: []string{state_input_price,
					state_input_square_meters, state_input_year, state_input_mill, state_set_city, state_set_radius, state_select_category},
					Dst: state_base},

				{Name: state_set_city, Src: []string{state_base}, Dst: state_set_city},
				{Name: state_set_radius, Src: []string{state_set_city}, Dst: state_set_radius},
				{Name: state_select_category, Src: []string{state_set_radius}, Dst: state_select_category},

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
		log.Printf("current state: %s", b.FSM.Current())
		if update.Message != nil {
			if update.Message.IsCommand() {
				if err := b.handleCommand(update.Message); err != nil {
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
			if err := b.updateHandleMessage(ctx, update.Message); err != nil {
				return fmt.Errorf("handle message error: %w", err)
			}
		} else if update.CallbackQuery != nil {
			log.Print("update callback query")
			var err error
			b.handleDeleteFilter(ctx, update.CallbackQuery, CurFilter, Filters)
			switch b.FSM.Current() {
			case state_set_city:
				ID, err = b.db.AddChatIDFilters(ctx, int(update.CallbackQuery.Message.Chat.ID))
				if err != nil {
					log.Printf("AddChatIDFilters error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}
					return nil
				}
				if err := b.handlerCityInlineKeyboard(Ctxt, update.CallbackQuery); err != nil {
					log.Printf("handlerCityInlineKeyboard error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, replyErr)
					b.bot.Send(msg)
					continue
				}
				if err := b.db.AddCityFilter(ctx, ID, update.CallbackQuery.Data); err != nil {
					log.Printf("AddCityFilter error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}
					return nil
				}
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, replySuccessfullySelectCity)
				if _, err := b.bot.Send(msg); err != nil {
					log.Printf("error send message: %v", err)
				}
				if err := b.FSM.Event(ctx, state_set_radius); err != nil {
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
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
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
					b.bot.Send(msg)
				}
				if err := b.db.AddRadiusFilter(ctx, ID, update.CallbackQuery.Data); err != nil {
					log.Printf("AddRadiusFilter error: %v", err)
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}
					return nil
				}
				msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, replySuccessfullySelectRadius)
				if _, err := b.bot.Send(msg); err != nil {
					log.Printf("error send message: %v", err)
					continue
				}

				if err := b.FSM.Event(ctx, state_select_category); err != nil {
					msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
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
				go func() {
					for {
						if Vehicles || Propertyrentals || Toys || Instruments || Home_improvements || Classifieds || Apparel || Propertyforsale || Entertainment || Family || Sports || Home || Pets || Office_supplies || Garden || Hobbies || Electronics {
							<-ChInputPrice
							msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите минимальную и максимальную цену (в валюте, которая в том городе, который вы выбрали ранее, например если этот город в США, то валюта будет доллары) через запятую:\n111, 99999\nНеважно в каком порядке, большее число будет считаться как максимальное")
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Printf("error send message: %v", err)
							}
							if err := b.FSM.Event(ctx, state_input_price); err != nil {
								msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
								return
							}
							return
						} else if Free || Groups || All_listings {
							return
						}
					}

				}()

				go func() {
					for {
						if Vehicles {
							log.Print("wait ChInputYear")
							<-ChInputYear
							log.Print("get ChInputYear")
							msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите минимальный и максимальный год через запятую:\n2020, 2024\nНеважно в каком порядке, больший год будет считаться как максимальный")
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Printf("error send message: %v", err)
							}
							if err := b.FSM.Event(ctx, state_input_year); err != nil {
								msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
								return
							}
							return
						} else if Propertyrentals || Free || Toys || Instruments || Home_improvements || Classifieds || Apparel || Propertyforsale || Entertainment || Family || Sports || Home || Pets || Office_supplies || Garden || Hobbies || Electronics || Groups || All_listings {
							return
						}
					}
				}()
				go func() {
					for {
						if Propertyrentals {
							<-ChSquareMet
							msg = tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Введите минимальные и максимальные квадратные метры через запятую:\n100, 3000\nНеважно в каком порядке, большее число будет считаться как максимальный")
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Printf("error send message: %v", err)
							}
							if err := b.FSM.Event(ctx, state_input_square_meters); err != nil {
								msg := tgbotapi.NewMessage(update.CallbackQuery.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
								return
							}
							return
						} else if Vehicles || Free || Toys || Instruments || Home_improvements || Classifieds || Apparel || Propertyforsale || Entertainment || Family || Sports || Home || Pets || Office_supplies || Garden || Hobbies || Electronics || Groups || All_listings {
							return
						}
					}

				}()
				go func(ChatID int64) {
					for {
						if Vehicles {
							if All || Cars_and_lorries {
								log.Print("wait ChInputMill")
								<-ChInputMill
								log.Print("get ChInputMill")
								msg = tgbotapi.NewMessage(ChatID, "Введите минимальный и максимальный пробег (в той метрической системе, которая в выбранном вами городе, например если город в США, то тогда пробег измеряется в милях) через запятую:\n100, 5000\nНеважно в каком порядке, больший пробег будет считаться как максимальный")
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Printf("error send message: %v", err)
								}
								if err := b.FSM.Event(ctx, state_input_mill); err != nil {
									msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
									msg.ReplyMarkup = StartKeyboard
									_, err = b.bot.Send(msg)
									if err != nil {
										log.Fatalf("[handleMessage]error send message: %v", err)
									}
									return
								}
								return
							} else if Motorcycles || Powersports || Motorhomes_and_campers || Boats || Commercial_and_industrial || Trailers || Other {
								return
							}
						} else if Propertyrentals || Free || Toys || Instruments || Home_improvements || Classifieds || Apparel || Propertyforsale || Entertainment || Family || Sports || Home || Pets || Office_supplies || Garden || Hobbies || Electronics || Groups || All_listings {
							return
						}
					}
				}(update.CallbackQuery.Message.Chat.ID)

			case state_select_category:
				log.Print("update.CallbackQuery SelectInlineKB")
				if err := b.handlerCategoryInlineKeyboard(ctx, update.CallbackQuery); err != nil {
					return fmt.Errorf("handlerCategoryInlineKeyboard error: %w", err)
				}
			}
		}
	}
	return nil
}
