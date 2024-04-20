package telegram

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"facebook_marketplace_bot/internal/parser"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"sync"

	"github.com/chromedp/chromedp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var (
	ID               int
	url1             string
	SelectCity       bool
	SelectRadius     bool
	SelectInlineKB   bool
	InputPrice       bool
	InputYear        bool
	InputMet         bool
	InputMill        bool
	DeleteFilter     bool
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
	_, err = b.bot.Send(msg)
	if err != nil {
		return fmt.Errorf("[handleFile]error send message: %w", err)
	}

	return nil
}
func (b *Bot) handleMessage(ctx context.Context, message *tgbotapi.Message) error {

	log.Printf("[%s] %s", message.From.UserName, message.Text)
	waitMsg, err := b.db.WaitMessage(ctx, int(message.Chat.ID))
	if err != nil {
		return fmt.Errorf("get wait message error: %w", err)
	}
	if message.Text == "Добавить фильтр" && !waitMsg {
		if err := b.db.AddWaitMessage(ctx, int(message.Chat.ID)); err != nil {
			return fmt.Errorf("add wait message error: %w", err)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, replySetCity)
		msg.ReplyMarkup = CancelKeyboard
		_, err := b.bot.Send(msg)
		if err != nil {
			return fmt.Errorf("[handleMessage]error send message: %w", err)
		}
	} else if message.Text == "Мои фильтры" && !waitMsg {
		filters, err := b.db.SelectAllFilter(ctx, int(message.Chat.ID))
		if err != nil {
			log.Printf("SelectAllFilter error: %v", err)
			msg := tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
				log.Fatalf("delete wait message error: %v", err)
			}
			SelectCity = false
			SelectRadius = false
			SelectInlineKB = false
			InputPrice = false
			InputYear = false
			InputMill = false
			DeleteFilter = false
			return nil
		}
		if len(filters) == 0 || filters == nil {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Вы еще не добавляли фильтров!")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			return nil
		}
		for i, filter := range filters {
			if int(message.Chat.ID) == filter.Chat_id {
				DeleteFilter = true
				FilterInlineKeyboard = tgbotapi.InlineKeyboardMarkup{}
				go func(filter database.Filter) {
					log.Print("wait get filter")
					ChFilter <- filter
					log.Print("set filter")
				}(filter)
				FilterInlineKeyboard.InlineKeyboard = append(FilterInlineKeyboard.InlineKeyboard, tgbotapi.NewInlineKeyboardRow(
					tgbotapi.NewInlineKeyboardButtonData("Удалить фильтр", "filter_id_"+filter.Id),
				))
				msg := tgbotapi.NewMessage(message.Chat.ID, fmt.Sprintf("Фильтр №%d\nГород: %s\nРадиус: %s\nКатегория: %s", i+1, filter.City, filter.Radius, filter.Category))
				msg.ReplyMarkup = FilterInlineKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
			}
		}
		go func() {
			<-isDel
			msg := tgbotapi.NewMessage(message.Chat.ID, "Успешно удалено!")
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}

		}()

	}
	if message.Text == "Отмена" && waitMsg {
		if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
			return fmt.Errorf("delete wait message error: %w", err)
		}
		msg := tgbotapi.NewMessage(message.Chat.ID, replyCancel)
		msg.ReplyMarkup = StartKeyboard
		_, err := b.bot.Send(msg)
		if err != nil {
			return fmt.Errorf("[handleMessage]error send message: %w", err)
		}
		if err := b.db.DeleteFilter(ctx, ID); err != nil {
			log.Printf("DeleteFilter error: %v", err)
			msg = tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова отменить добавление фильтра")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			SelectCity = false
			SelectRadius = false
			SelectInlineKB = false
			InputPrice = false
			InputYear = false
			InputMill = false
			DeleteFilter = false
			return nil
		}
		SelectCity = false
		SelectRadius = false
		SelectInlineKB = false
		InputPrice = false
		InputYear = false
		InputMill = false
		DeleteFilter = false
		return nil
	}

	if InputPrice {
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
		InputPrice = false
		return nil
	}
	if InputMet {
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
		InputMet = false
		return nil
	}
	if InputYear {
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
		InputYear = false
		return nil
	}
	if InputMill {
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
		InputMill = false
		return nil
	}
	if waitMsg {
		if message.Text == "Добавить фильтр" || message.Text == "Мои фильтры" || message.Text == "Отмена" {
			msg := tgbotapi.NewMessage(message.Chat.ID, "Введи корректное название города!")
			_, err := b.bot.Send(msg)
			if err != nil {
				return fmt.Errorf("[handleMessage]error send message: %w", err)
			}
			return nil
		}

		go func() {
			log.Print("!!!goruntine!!!")
			msg := tgbotapi.NewMessage(message.Chat.ID, "Идет установка города в facebook...")
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			b.rw.Lock()
			datas, err := b.parser.GetData()
			if err == parser.ErrEmptyData {
				for _, id := range Admins {
					if message.Chat.ID != id {
						msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Fatalf("[handleMessage]error send message: %v", err)
						}
					} else if message.Chat.ID == id {
						msg := tgbotapi.NewMessage(id, "Не хватает аккаунтов, добавьте txt файл через команду /load")
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Fatalf("[handleMessage]error send message: %v", err)

						}
					}
				}
				if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
					log.Fatalf("delete wait message error: %v", err)
				}
				return
			}
			if err != nil && err != parser.ErrEmptyData {
				log.Fatalf("get data error: %v", err)
			}
			b.rw.Unlock()
			log.Printf("!!!!!data: %v", datas)
			var once sync.Once
			for _, data := range datas {
				b.rw.Lock()
				opts := append(chromedp.DefaultExecAllocatorOptions[:],
					chromedp.ProxyServer("http://"+data.Datas.IpPortPX),
					chromedp.WindowSize(1900, 1080), // init with a desktop view
					chromedp.Flag("enable-automation", false),
					// chromedp.Flag("headless", false),
				)
				b.rw.Unlock()
				ctxChr, cancel := chromedp.NewExecAllocator(context.Background(), opts...)
				defer cancel()
				ctxt, cancel := chromedp.NewContext(ctxChr) // chromedp.WithDebugf(log.Printf),

				defer cancel()
				defer chromedp.Cancel(ctxt)
				log.Print("!!!settings!!!")
				b.rw.Lock()
				err = b.parser.Settings(ctxt, data)
				switch err {
				case parser.ErrProxyConnectionFailed:
					for _, id := range Admins {
						if message.Chat.ID != id {
							once.Do(func() {
								msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте позже")
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
							})
						} else if message.Chat.ID == id {
							dataFile, err := os.ReadDir(parser.Free_account)
							if err != nil {
								log.Printf("ReadDir error: %v", err)
								msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте позже")
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
							}
							if len(dataFile) == 0 {
								for _, id := range Admins {
									if message.Chat.ID != id {
										msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
										_, err = b.bot.Send(msg)
										if err != nil {
											log.Fatalf("[handleMessage]error send message: %v", err)
										}
									} else if message.Chat.ID == id {
										msg := tgbotapi.NewMessage(id, "Не хватает аккаунтов, добавьте txt файл через команду /load")
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
												msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
												_, err = b.bot.Send(msg)
												if err != nil {
													log.Fatalf("[handleMessage]error send message: %v", err)
												}
											} else if message.Chat.ID == id {
												msg := tgbotapi.NewMessage(id, "Ошибка при удалении аккаунта с не валидными прокси")
												_, err = b.bot.Send(msg)
												if err != nil {
													log.Fatalf("[handleMessage]error send message: %v", err)

												}
											}
										}
									}
									msg := tgbotapi.NewMessage(id, "Профиль был автоматически удален, причина: не валидные прокси в файле: "+data.FileName)
									_, err = b.bot.Send(msg)
									if err != nil {
										log.Fatalf("[handleMessage]error send message: %v", err)
									}
								}
							}

						}

					}
					if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
						log.Fatalf("delete wait message error: %v", err)
					}
					b.rw.Unlock()
					continue
				case parser.ErrAccountBanned:
					for _, id := range Admins {
						if message.Chat.ID == id {
							dataFile, err := os.ReadDir(parser.Free_account)
							if err != nil {
								log.Printf("ReadDir error: %v", err)
								msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте позже")
								msg.ReplyMarkup = StartKeyboard
								_, err = b.bot.Send(msg)
								if err != nil {
									log.Fatalf("[handleMessage]error send message: %v", err)
								}
							}
							if len(dataFile) == 0 {
								for _, id := range Admins {
									if message.Chat.ID != id {
										msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
										_, err = b.bot.Send(msg)
										if err != nil {
											log.Fatalf("[handleMessage]error send message: %v", err)
										}
									} else if message.Chat.ID == id {
										msg := tgbotapi.NewMessage(id, "Не хватает аккаунтов, добавьте txt файл через команду /load")
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
												msg := tgbotapi.NewMessage(message.Chat.ID, "Технические неполадки, попробуйте снова")
												_, err = b.bot.Send(msg)
												if err != nil {
													log.Fatalf("[handleMessage]error send message: %v", err)
												}
											} else if message.Chat.ID == id {
												msg := tgbotapi.NewMessage(id, "Ошибка при удалении забаненного аккаунта")
												_, err = b.bot.Send(msg)
												if err != nil {
													log.Fatalf("[handleMessage]error send message: %v", err)

												}
											}
										}
									}
									msg := tgbotapi.NewMessage(id, "Профиль был автоматически удален, причина: аккаунт забаннен, файл: "+data.FileName)
									_, err = b.bot.Send(msg)
									if err != nil {
										log.Fatalf("[handleMessage]error send message: %v", err)
									}
								}
							}

						}
					}
					b.rw.Unlock()
					continue
				}
				if err != nil && err != parser.ErrEmptyData && err != parser.ErrAccountBanned {
					log.Printf("settings error: %v", err)
					b.rw.Unlock()
					continue
				}
				b.rw.Unlock()
				log.Print("!!!end settings!!!")
				log.Printf("city: %s", message.Text)
				cities, err := b.parser.SelectCity(ctxt, message.Text)
				if err != nil {
					log.Printf("select city error: %v", err)
					msg := tgbotapi.NewMessage(message.Chat.ID, replyErr)
					b.bot.Send(msg)
					continue
				}
				if len(cities) == 0 {
					msg := tgbotapi.NewMessage(message.Chat.ID, "Было введено не корректное название города!")
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("[handleMessage]error send message: %v", err)

					}
					return
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
					continue
				}
				CityInlineKeyboard = tgbotapi.InlineKeyboardMarkup{}
				SelectCity = true
				var isCorrect bool
				callbackQuerySelectCity := <-ChSelectCity
				for _, city := range cities {
					if city == callbackQuerySelectCity.Data {

						ID, err = b.db.AddChatIDFilters(ctx, int(message.Chat.ID))
						if err != nil {
							log.Printf("AddChatIDFilters error: %v", err)
							msg = tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
							msg.ReplyMarkup = StartKeyboard
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Fatalf("[handleMessage]error send message: %v", err)
							}
							if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
								log.Fatalf("delete wait message error: %v", err)
							}
							SelectCity = false
							SelectRadius = false
							SelectInlineKB = false
							InputPrice = false
							InputYear = false
							InputMill = false
							DeleteFilter = false
							return
						}
						isCorrect = true
						log.Printf("callbackQuery: %v\ncities: %v", callbackQuerySelectCity.Data, city)
						if err := b.handlerCityInlineKeyboard(ctxt, callbackQuerySelectCity); err != nil {
							log.Printf("handlerCityInlineKeyboard error: %v", err)
							msg := tgbotapi.NewMessage(message.Chat.ID, replyErr)
							b.bot.Send(msg)
							continue
						}
						msg = tgbotapi.NewMessage(callbackQuerySelectCity.Message.Chat.ID, replySuccessfullySelectCity)

						if _, err := b.bot.Send(msg); err != nil {
							log.Printf("error send message: %v", err)
							continue
						}
						if err := b.db.AddCityFilter(ctx, ID, city); err != nil {
							log.Printf("AddCityFilter error: %v", err)
							msg = tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
							msg.ReplyMarkup = StartKeyboard
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Fatalf("[handleMessage]error send message: %v", err)
							}
							if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
								log.Fatalf("delete wait message error: %v", err)
							}
							SelectCity = false
							SelectRadius = false
							SelectInlineKB = false
							InputPrice = false
							InputYear = false
							DeleteFilter = false
							InputMill = false
							return
						}
						SelectCity = false
						SelectRadius = true
						msg = tgbotapi.NewMessage(message.Chat.ID, replySelectRadius)
						msg.ReplyMarkup = RadiusInlineKeyboard
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Printf("error send message: %v", err)
							continue
						}
						log.Print("wait click inline keyboard")

						callbackQuerySelectRadius := <-ChSelectRadius
						if err := b.handlerRadiusInlineKeyboard(ctxt, callbackQuerySelectRadius); err != nil {
							log.Printf("handlerRadiusInlineKeyboard error: %v", err)
							msg := tgbotapi.NewMessage(message.Chat.ID, replyErr)
							b.bot.Send(msg)
							continue
						}
						msg = tgbotapi.NewMessage(callbackQuerySelectCity.Message.Chat.ID, replySuccessfullySelectRadius)

						if _, err := b.bot.Send(msg); err != nil {
							log.Printf("error send message: %v", err)
							continue
						}
						if err := b.db.AddRadiusFilter(ctx, int(message.Chat.ID), callbackQuerySelectRadius.Data); err != nil {
							log.Printf("AddRadiusFilter error: %v", err)
							msg = tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
							msg.ReplyMarkup = StartKeyboard
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Fatalf("[handleMessage]error send message: %v", err)
							}
							if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
								log.Fatalf("delete wait message error: %v", err)
							}
							SelectCity = false
							SelectRadius = false
							SelectInlineKB = false
							InputPrice = false
							InputYear = false
							InputMill = false
							DeleteFilter = false
							return
						}
						SelectRadius = false
						SelectInlineKB = true
						msg = tgbotapi.NewMessage(callbackQuerySelectCity.Message.Chat.ID, replySelectCategoty)
						msg.ReplyMarkup = CategoryInlineKeyboard1
						sendmsg, err := b.bot.Send(msg)
						if err != nil {
							log.Printf("error send message: %v", err)
							continue
						}
						SelectCategory = sendmsg
						go func() {
							inPrice := <-ChInputPrice
							log.Print(inPrice)
							msg := tgbotapi.NewMessage(message.Chat.ID, "Введите минимальную и максимальную цену (в валюте, которая в том городе, который вы выбрали ранее, например если этот город в США, то валюта будет доллары) через запятую:\n111, 99999\nНеважно в каком порядке, большее число будет считаться как максимальное")
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Printf("error send message: %v", err)
							}
							InputPrice = true
						}()

						go func() {
							inYear := <-ChInputYear
							log.Print(inYear)
							msg = tgbotapi.NewMessage(message.Chat.ID, "Введите минимальный и максимальный год через запятую:\n2020, 2024\nНеважно в каком порядке, больший год будет считаться как максимальный")
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Printf("error send message: %v", err)
							}
							InputYear = true
						}()
						go func() {
							inMet := <-ChSquareMet
							log.Print(inMet)
							msg = tgbotapi.NewMessage(message.Chat.ID, "Введите минимальные и максимальные квадратные метры через запятую:\n100, 3000\nНеважно в каком порядке, большее число будет считаться как максимальный")
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Printf("error send message: %v", err)
							}
							InputMet = true
						}()
						go func() {
							<-ChInputMill
							msg = tgbotapi.NewMessage(message.Chat.ID, "Введите минимальный и максимальный пробег (в той метрической системе, которая в выбранном вами городе, например если город в США, то тогда пробег измеряется в милях) через запятую:\n100, 5000\nНеважно в каком порядке, больший пробег будет считаться как максимальный")
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Printf("error send message: %v", err)
							}
							InputMill = true
						}()

						url1 = <-ChUrl
						slcCat := <-ChSelectCategory
						log.Print(slcCat)
						if err := b.db.AddMonitoringFilter(ctx, ID, url1); err != nil {
							log.Printf("AddMonitoringFilter error: %v", err)
							msg = tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, добавьте фильтр еще раз")
							msg.ReplyMarkup = StartKeyboard
							_, err = b.bot.Send(msg)
							if err != nil {
								log.Fatalf("[handleMessage]error send message: %v", err)
							}
							if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
								log.Fatalf("delete wait message error: %v", err)
							}
							SelectCity = false
							SelectRadius = false
							SelectInlineKB = false
							InputPrice = false
							InputYear = false
							InputMill = false
							DeleteFilter = false
							return
						}
						SelectInlineKB = false
						msg = tgbotapi.NewMessage(message.Chat.ID, "Успешно создан фильтр!")
						msg.ReplyMarkup = StartKeyboard
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Printf("error send message: %v", err)
						}
					}
				}
				if !isCorrect {
					msg = tgbotapi.NewMessage(callbackQuerySelectCity.Message.Chat.ID, replyError)
					msg.ReplyMarkup = StartKeyboard
					if _, err := b.bot.Send(msg); err != nil {
						log.Printf("error send message: %v", err)
						continue
					}
					return
				}
				if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
					log.Printf("DeleteWaitMessage error: %v", err)
					msg = tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}
					if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
						log.Fatalf("delete wait message error: %v", err)
					}
					SelectCity = false
					SelectRadius = false
					SelectInlineKB = false
					InputPrice = false
					InputYear = false
					InputMill = false
					DeleteFilter = false
					return
				}
				if err := os.Rename(parser.Free_account+data.FileName, parser.Work_account+data.FileName); err != nil {
					log.Printf("Rename error: %v", err)
					msg = tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}
					if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
						log.Fatalf("delete wait message error: %v", err)
					}
					SelectCity = false
					SelectRadius = false
					SelectInlineKB = false
					InputPrice = false
					InputYear = false
					InputMill = false
					DeleteFilter = false
					return
				}
				go func() {
					for id := range parser.ChId {
						log.Print(id)
						msg = tgbotapi.NewMessage(message.Chat.ID, "Найдено новое объявление: https://www.facebook.com/marketplace/item/"+strconv.Itoa(id))
						_, err = b.bot.Send(msg)
						if err != nil {
							log.Printf("error send message: %v", err)
						}
					}
				}()
				err = b.parser.Monitoring(ctxt, url1, ID)
				if err != nil {
					log.Printf("Monitoring error: %v", err)
					msg = tgbotapi.NewMessage(message.Chat.ID, "Произошла ошибка мониторинга, добавьте фильтр еще раз")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}
					if err := b.db.DeleteWaitMessage(ctx, int(message.Chat.ID)); err != nil {
						log.Fatalf("delete wait message error: %v", err)
					}
					SelectCity = false
					SelectRadius = false
					SelectInlineKB = false
					InputPrice = false
					InputYear = false
					InputMill = false
					DeleteFilter = false
					return
				}
				break
			}
		}()
	}
	return nil
}
