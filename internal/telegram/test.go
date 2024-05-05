package telegram

import (
	"context"
	"facebook_marketplace_bot/internal/parser"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/chromedp/chromedp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	Ctxt            context.Context
	CurrentFileName string
	Cancel1         context.CancelFunc
	Cancel2         context.CancelFunc
)

func (b *Bot) updateHandleMessage(ctx context.Context, message *tgbotapi.Message) error {
	log.Printf("[%s] %s", message.From.UserName, message.Text)
	if b.FSM.Current() != state_base {
		if message.Text == "Отмена" {
			if err := b.db.DeleteFilter(ctx, ID); err != nil {
				log.Printf("DeleteFilter error: %v", err)
				msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова отменить добавление фильтра")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				return nil
			}
			if err := b.FSM.Event(ctx, state_cancel); err != nil {
				msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				return nil
			}
			msg := tgbotapi.NewMessage(message.Chat.ID, replyCancel)
			msg.ReplyMarkup = StartKeyboard
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}

			return nil
		}
	}
	switch b.FSM.Current() {
	case state_base:
		if message.Text == "Добавить фильтр" {
			msg := tgbotapi.NewMessage(message.Chat.ID, replySetCity)
			msg.ReplyMarkup = CancelKeyboard
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			if err := b.FSM.Event(ctx, state_set_city); err != nil {
				msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
			}
			return nil
		} else if message.Text == "Мои фильтры" {
			Filters, err := b.db.SelectAllFilter(ctx, int(message.Chat.ID))
			if err != nil {
				log.Printf("SelectAllFilter error: %v", err)
				msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				return nil
			}
			if len(Filters) == 0 || Filters == nil {
				msg := tgbotapi.NewMessage(message.Chat.ID, "Вы еще не добавляли фильтров!")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				return nil
			}
			if len(Filters) == 1 {
				msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Город: %s\nРадиус: %s\nКатегория: %s", Filters[0].City, Filters[0].Radius, Filters[0].Category))
				msg.ReplyMarkup = SelectFilters3
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				CurFilter = Filters[0]
			} else {
				msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Город: %s\nРадиус: %s\nКатегория: %s", Filters[0].City, Filters[0].Radius, Filters[0].Category))
				msg.ReplyMarkup = SelectFilters
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				CurFilter = Filters[0]
			}
			return nil
		}
	case state_set_city:
		if message.Text == "Добавить фильтр" || message.Text == "Мои фильтры" {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Введи корректное название города!")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		log.Print("get data")
		datas, err := b.parser.GetData()
		if err == parser.ErrEmptyData {
			for _, id := range Admins {
				if message.Chat.ID != id {
					if err := b.FSM.Event(ctx, state_base); err != nil {
						msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
						msg.ReplyMarkup = StartKeyboard
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Fatalf("[handleMessage]error send message: %v", err)
						}
					}
					msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return nil
				} else if message.Chat.ID == id {
					if err := b.FSM.Event(ctx, state_base); err != nil {
						msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
						msg.ReplyMarkup = StartKeyboard
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Fatalf("[handleMessage]error send message: %v", err)
						}
					}
					msg := tgbotapi.NewMessage(id, "Не хватает аккаунтов, добавьте txt файл через команду /load")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)

					}

					return nil
				}
			}
		}
		if err != nil && err != parser.ErrEmptyData {
			log.Fatalf("get data error: %v", err)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, "Идет установка города в facebook...")
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Fatalf("[handleMessage]error send message: %v", err)
		}
		log.Printf("!!!!!data: %v", datas)

		for _, data := range datas {

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
			switch err {
			case parser.ErrProxyConnectionFailed:
				for _, id := range Admins {
					if message.Chat.ID != id {
						if err := b.FSM.Event(ctx, state_base); err != nil {
							msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
							msg.ReplyMarkup = StartKeyboard
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Fatalf("[handleMessage]error send message: %v", err)
							}
						}
						msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте позже")
						msg.ReplyMarkup = StartKeyboard
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Fatalf("[handleMessage]error send message: %v", err)
						}
						return nil
					} else if message.Chat.ID == id {
						dataFile, err := os.ReadDir(parser.Free_account)
						if err != nil {
							if err := b.FSM.Event(ctx, state_base); err != nil {
								msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
							}
							log.Printf("ReadDir error: %v", err)
							msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте позже")
							msg.ReplyMarkup = StartKeyboard
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Fatalf("[handleMessage]error send message: %v", err)
							}
							return nil
						}
						if len(dataFile) == 0 {
							for _, id := range Admins {
								if message.Chat.ID != id {
									if err := b.FSM.Event(ctx, state_base); err != nil {
										msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
										msg.ReplyMarkup = StartKeyboard
										_, err = b.bot.Send(msg)
										if err != nil {
											log.Fatalf("[handleMessage]error send message: %v", err)
										}
									}
									msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
									msg.ReplyMarkup = StartKeyboard
									_, err = b.bot.Send(msg)
									if err != nil {
										log.Fatalf("[handleMessage]error send message: %v", err)
									}
									return nil
								} else if message.Chat.ID == id {
									if err := b.FSM.Event(ctx, state_base); err != nil {
										msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
										msg.ReplyMarkup = StartKeyboard
										_, err = b.bot.Send(msg)
										if err != nil {
											log.Fatalf("[handleMessage]error send message: %v", err)
										}
									}
									msg := tgbotapi.NewMessage(id, "Не хватает аккаунтов, добавьте txt файл через команду /load")
									msg.ReplyMarkup = StartKeyboard
									_, err = b.bot.Send(msg)
									if err != nil {
										log.Fatalf("[handleMessage]error send message: %v", err)

									}
								}
							}
						}
						for _, file := range dataFile {
							if file.Name() == data.FileName {
								if err := os.Remove(parser.Free_account + data.FileName); err != nil {
									log.Printf("remove error: %v", err)
									for _, id := range Admins {
										if message.Chat.ID != id {
											if err := b.FSM.Event(ctx, state_base); err != nil {
												msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
												msg.ReplyMarkup = StartKeyboard
												_, err = b.bot.Send(msg)
												if err != nil {
													log.Fatalf("[handleMessage]error send message: %v", err)
												}
											}
											msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
											msg.ReplyMarkup = StartKeyboard
											_, err = b.bot.Send(msg)
											if err != nil {
												log.Fatalf("[handleMessage]error send message: %v", err)
											}
										} else if message.Chat.ID == id {
											if err := b.FSM.Event(ctx, state_base); err != nil {
												msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
												msg.ReplyMarkup = StartKeyboard
												_, err = b.bot.Send(msg)
												if err != nil {
													log.Fatalf("[handleMessage]error send message: %v", err)
												}
											}
											msg := tgbotapi.NewMessage(id, "Ошибка при удалении аккаунта с не валидными прокси")
											msg.ReplyMarkup = StartKeyboard
											_, err = b.bot.Send(msg)
											if err != nil {
												log.Fatalf("[handleMessage]error send message: %v", err)

											}
										}
									}
								}
								if err := b.FSM.Event(ctx, state_base); err != nil {
									msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
									msg.ReplyMarkup = StartKeyboard
									_, err = b.bot.Send(msg)
									if err != nil {
										log.Fatalf("[handleMessage]error send message: %v", err)
									}
								}
								msg := tgbotapi.NewMessage(id, "Профиль был автоматически удален, причина: не валидные прокси в файле: "+data.FileName)
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
								return nil
							}
						}

					}

				}
			case parser.ErrAccountBanned:
				for _, id := range Admins {
					if message.Chat.ID == id {
						dataFile, err := os.ReadDir(parser.Free_account)
						if err != nil {
							if err := b.FSM.Event(ctx, state_base); err != nil {
								msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
							}
							log.Printf("ReadDir error: %v", err)
							msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте позже")
							msg.ReplyMarkup = StartKeyboard
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Fatalf("[handleMessage]error send message: %v", err)
							}
							return nil
						}
						if len(dataFile) == 0 {
							for _, id := range Admins {
								if message.Chat.ID != id {
									if err := b.FSM.Event(ctx, state_base); err != nil {
										msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
										msg.ReplyMarkup = StartKeyboard
										_, err = b.bot.Send(msg)
										if err != nil {
											log.Fatalf("[handleMessage]error send message: %v", err)
										}
									}
									msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
									msg.ReplyMarkup = StartKeyboard
									_, err = b.bot.Send(msg)
									if err != nil {
										log.Fatalf("[handleMessage]error send message: %v", err)
									}
									return nil
								} else if message.Chat.ID == id {
									if err := b.FSM.Event(ctx, state_base); err != nil {
										msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
										msg.ReplyMarkup = StartKeyboard
										_, err = b.bot.Send(msg)
										if err != nil {
											log.Fatalf("[handleMessage]error send message: %v", err)
										}
									}
									msg := tgbotapi.NewMessage(id, "Не хватает аккаунтов, добавьте txt файл через команду /load")
									msg.ReplyMarkup = StartKeyboard
									_, err = b.bot.Send(msg)
									if err != nil {
										log.Fatalf("[handleMessage]error send message: %v", err)

									}
									return nil
								}
							}
						}
						for _, file := range dataFile {
							if file.Name() == data.FileName {
								if err := os.Remove(parser.Free_account + data.FileName); err != nil {
									log.Printf("remove error: %v", err)
									for _, id := range Admins {
										if message.Chat.ID != id {
											if err := b.FSM.Event(ctx, state_base); err != nil {
												msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
												msg.ReplyMarkup = StartKeyboard
												_, err = b.bot.Send(msg)
												if err != nil {
													log.Fatalf("[handleMessage]error send message: %v", err)
												}
											}
											msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
											msg.ReplyMarkup = StartKeyboard
											_, err = b.bot.Send(msg)
											if err != nil {
												log.Fatalf("[handleMessage]error send message: %v", err)
											}
											return nil
										} else if message.Chat.ID == id {
											if err := b.FSM.Event(ctx, state_base); err != nil {
												msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
												msg.ReplyMarkup = StartKeyboard
												_, err = b.bot.Send(msg)
												if err != nil {
													log.Fatalf("[handleMessage]error send message: %v", err)
												}
											}
											msg := tgbotapi.NewMessage(id, "Ошибка при удалении забаненного аккаунта")
											msg.ReplyMarkup = StartKeyboard
											_, err = b.bot.Send(msg)
											if err != nil {
												log.Fatalf("[handleMessage]error send message: %v", err)

											}
											return nil
										}
									}
								}
								for _, id := range Admins {
									if message.Chat.ID != id {
										if err := b.FSM.Event(ctx, state_base); err != nil {
											msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
											msg.ReplyMarkup = StartKeyboard
											_, err = b.bot.Send(msg)
											if err != nil {
												log.Fatalf("[handleMessage]error send message: %v", err)
											}
										}
										msg := tgbotapi.NewMessage(id, "Произошла ошибка, попробуйте снова")
										msg.ReplyMarkup = StartKeyboard
										_, err = b.bot.Send(msg)
										if err != nil {
											log.Fatalf("[handleMessage]error send message: %v", err)
										}
										return nil
									} else if message.Chat.ID == id {
										if err := b.FSM.Event(ctx, state_base); err != nil {
											msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
											msg.ReplyMarkup = StartKeyboard
											_, err = b.bot.Send(msg)
											if err != nil {
												log.Fatalf("[handleMessage]error send message: %v", err)
											}
										}
										msg := tgbotapi.NewMessage(id, "Профиль был автоматически удален, причина: аккаунт забаннен, файл: "+data.FileName)
										msg.ReplyMarkup = StartKeyboard
										_, err = b.bot.Send(msg)
										if err != nil {
											log.Fatalf("[handleMessage]error send message: %v", err)
										}
										return nil
									}
								}
							}
						}

					}
				}
				if err != nil && err != parser.ErrEmptyData && err != parser.ErrAccountBanned {
					log.Printf("settings error: %v", err)
				}
			}
			log.Print("!!!end settings!!!")
			log.Printf("city: %s", message.Text)
			cities, err := b.parser.SelectCity(Ctxt, message.Text)
			if err != nil {
				log.Printf("select city error: %v", err)
				msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				return nil
			}
			if len(cities) == 0 {
				msg := tgbotapi.NewMessage(message.Chat.ID, "Было введено не корректное название города!")
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Printf("[handleMessage]error send message: %v", err)

				}
				return nil
			}

			for _, city := range cities {
				CityInlineKeyboard.InlineKeyboard = append(CityInlineKeyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData(city, city),
				))
			}
			msg = tgbotapi.NewMessage(message.Chat.ID, replySelectCity)
			msg.ReplyMarkup = CityInlineKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Printf("[handleMessage]error send message: %v", err)
				return nil
			}
			CityInlineKeyboard = tgbotapi.InlineKeyboardMarkup{}
			CurrentFileName = data.FileName
		}
	case state_input_price:
		log.Print("InputPrice ")
		var (
			num1 string
			num2 string
		)

		price := strings.TrimSpace(message.Text)
		minMax := strings.Split(price, ",")
		if len(minMax) > 2 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Введи только 2 числа!")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		for ind, nums := range minMax {
			num := strings.Split(nums, "")
			log.Printf("num: %s", num)
			for _, n := range num {
				log.Printf("n: %s", n)
				if n == " " {
					continue
				}
				_, err := strconv.Atoi(n)
				if err != nil {
					log.Printf("Atoi err: %v", err)
					msg := tgbotapi.NewMessage(message.Chat.ID, "Введи только числа без лишних символов!")
					_, err := b.bot.Send(msg)
					if err != nil {
						return fmt.Errorf("[handleMessage]error send message: %w", err)
					}
					return nil
				}
				if ind == 0 {
					num1 += n
				} else if ind == 1 {
					num2 += n
				}
			}
		}
		n1, err := strconv.Atoi(num1)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, введите максимальную и минимальную цену еще раз")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		n2, err := strconv.Atoi(num2)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, введите максимальную и минимальную цену еще раз")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		if n1 > n2 {
			go func() {
				ChMinPrice <- num2
				ChMaxPrice <- num1
			}()
		} else if n2 > n1 {
			go func() {
				ChMinPrice <- num1
				ChMaxPrice <- num2
			}()
		} else if n1 == n2 {
			go func() {
				ChMinPrice <- num1
				ChMaxPrice <- num2
			}()
		}
		if err := b.FSM.Event(ctx, state_select_category); err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
		}
	case state_input_year:
		var (
			num1 string
			num2 string
		)
		price := strings.TrimSpace(message.Text)
		minMax := strings.Split(price, ",")
		if len(minMax) > 2 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Введи только 2 года!")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		for ind, nums := range minMax {
			num := strings.Split(nums, "")
			for _, n := range num {
				if n == " " {
					continue
				}
				_, err := strconv.Atoi(n)
				if err != nil {
					msg := tgbotapi.NewMessage(message.Chat.ID, "Введи только года без лишних символов!")
					_, err := b.bot.Send(msg)
					if err != nil {
						return fmt.Errorf("[handleMessage]error send message: %w", err)
					}
					return nil
				}
				if ind == 0 {
					num1 += n
				} else if ind == 1 {
					num2 += n
				}
			}
		}
		n1, err := strconv.Atoi(num1)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла непредвидимая ошибка, введите максимальный и минимальный год еще раз")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		n2, err := strconv.Atoi(num2)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла непредвидимая ошибка, введите максимальный и минимальный год еще раз")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		if n1 > n2 {
			go func() {
				ChMinYear <- num2
				ChMaxYear <- num1
			}()
		} else if n2 > n1 {
			go func() {
				ChMinYear <- num1
				ChMaxYear <- num2
			}()
		} else if n1 == n2 {
			go func() {
				ChMinYear <- num1
				ChMaxYear <- num2
			}()
		}
		if err := b.FSM.Event(ctx, state_select_category); err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
		}
	case state_input_square_meters:
		log.Print("input_square_meters ")
		var (
			num1 string
			num2 string
		)

		price := strings.TrimSpace(message.Text)
		minMax := strings.Split(price, ",")
		if len(minMax) > 2 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Введи только 2 числа!")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		for ind, nums := range minMax {
			num := strings.Split(nums, "")
			log.Printf("num: %s", num)
			for _, n := range num {
				log.Printf("n: %s", n)
				if n == " " {
					continue
				}
				_, err := strconv.Atoi(n)
				if err != nil {
					log.Printf("Atoi err: %v", err)
					msg := tgbotapi.NewMessage(message.Chat.ID, "Введи только числа без лишних символов!")
					_, err := b.bot.Send(msg)
					if err != nil {
						return fmt.Errorf("[handleMessage]error send message: %w", err)
					}
					return nil
				}
				if ind == 0 {
					num1 += n
				} else if ind == 1 {
					num2 += n
				}
			}
		}
		n1, err := strconv.Atoi(num1)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла непредвидимая ошибка, введите максимальные и минимальные квадратные метры еще раз")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		n2, err := strconv.Atoi(num2)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла непредвидимая ошибка, введите максимальные и минимальные квадратные метры еще раз")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		if n1 > n2 {
			go func() {
				ChMinMet <- num2
				ChMaxMet <- num1
			}()
		} else if n2 > n1 {
			go func() {
				ChMinMet <- num1
				ChMaxMet <- num2
			}()
		} else if n1 == n2 {
			go func() {
				ChMinMet <- num1
				ChMaxMet <- num2
			}()
		}
		if err := b.FSM.Event(ctx, state_select_category); err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
		}
	case state_input_mill:
		var (
			num1 string
			num2 string
		)
		price := strings.TrimSpace(message.Text)
		minMax := strings.Split(price, ",")
		if len(minMax) > 2 {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Введи только 2 пробега!")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		for ind, nums := range minMax {
			num := strings.Split(nums, "")
			for _, n := range num {
				if n == " " {
					continue
				}
				_, err := strconv.Atoi(n)
				if err != nil {
					msg := tgbotapi.NewMessage(message.Chat.ID, "Введи только пробег без лишних символов!")
					_, err := b.bot.Send(msg)
					if err != nil {
						return fmt.Errorf("[handleMessage]error send message: %w", err)
					}
					return nil
				}
				if ind == 0 {
					num1 += n
				} else if ind == 1 {
					num2 += n
				}
			}
		}
		n1, err := strconv.Atoi(num1)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла непредвидимая ошибка, введите максимальный и минимальный пробег еще раз")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		n2, err := strconv.Atoi(num2)
		if err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла непредвидимая ошибка, введите максимальный и минимальный пробег еще раз")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}
		if n1 > n2 {
			go func() {
				ChMinMill <- num2
				ChMaxMill <- num1
			}()
		} else if n2 > n1 {
			go func() {
				ChMinMill <- num1
				ChMaxMill <- num2
			}()
		} else if n1 == n2 {
			go func() {
				ChMinMill <- num1
				ChMaxMill <- num2
			}()
		}
		if err := b.FSM.Event(ctx, state_select_category); err != nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
		}
	}
	return nil
}
