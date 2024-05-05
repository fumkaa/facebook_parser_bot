package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCategory(query *tgbotapi.CallbackQuery) error {
	switch query.Data {
	case "next1":
		switch {
		case Home_improvements:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2home_improvements)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
		case Classifieds:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2classifieds)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
		case Apparel:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2apparel)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
		case Propertyforsale:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2propertyforsale)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Entertainment:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2entertainment)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		}
	case "next2":
		switch {
		case Family:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3family)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Sports:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3sports)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Home:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3home)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Pets:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3pets)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Office_supplies:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3office_supplies)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		}
	case "next3":
		switch {
		case Garden:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4garden)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Hobbies:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4hobbies)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Electronics:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4electronics)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Groups:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4groups)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case All_listings:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4all_listings)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
		}
	case "previous1":
		switch {
		case Vehicles:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1vehicles)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Propertyrentals:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1propertyrentals)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Free:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1free)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Toys:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1toys)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Instruments:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1instruments)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		}
	case "previous2":
		switch {
		case Home_improvements:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2home_improvements)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Classifieds:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2classifieds)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Apparel:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2apparel)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Propertyforsale:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2propertyforsale)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Entertainment:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2entertainment)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		}
	case "previous3":
		switch {
		case Family:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3family)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Sports:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3sports)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Home:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3home)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Pets:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3pets)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		case Office_supplies:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3office_supplies)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3)
			_, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}

		}
	case "vehicles":
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false

		Vehicles = true
		log.Printf("vehicles: %t", Vehicles)
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1vehicles)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "vehicles1":
		Vehicles = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "propertyrentals":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Propertyrentals = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1propertyrentals)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "propertyrentals1":
		Propertyforsale = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "free":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Free = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1free)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "free1":
		Free = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "toys":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Toys = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1toys)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "toys1":
		Toys = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "instruments":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Instruments = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1instruments)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "instruments1":
		Instruments = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard1)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "home_improvements":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Home_improvements = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2home_improvements)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "home_improvements1":
		Home_improvements = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "classifieds":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Classifieds = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2classifieds)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "classifieds1":
		Classifieds = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "apparel":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Apparel = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2apparel)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "apparel1":
		Apparel = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "propertyforsale":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Propertyforsale = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2propertyforsale)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "propertyforsale1":
		Propertyforsale = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "entertainment":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Entertainment = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2entertainment)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "entertainment1":
		Entertainment = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard2)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "family":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Family = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3family)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "family1":
		Family = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "sports":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Sports = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3sports)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "sports1":
		Sports = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "home":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Home = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3home)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "home1":
		Home = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "pets":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Pets = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3pets)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "pets1":
		Pets = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "office_supplies":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Office_supplies = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3office_supplies)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "office_supplies1":
		Office_supplies = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard3)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "garden":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Garden = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4garden)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "garden1":
		Garden = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "hobbies":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Hobbies = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4hobbies)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "hobbies1":
		Hobbies = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "electronics":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Electronics = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4electronics)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "electronics1":
		Electronics = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "groups":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		Groups = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4groups)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "groups1":
		Groups = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "all_listings":
		Vehicles = false
		Propertyrentals = false
		Free = false
		Toys = false
		Instruments = false

		Home_improvements = false
		Classifieds = false
		Apparel = false
		Propertyforsale = false
		Entertainment = false

		Family = false
		Sports = false
		Home = false
		Pets = false
		Office_supplies = false

		Garden = false
		Hobbies = false
		Electronics = false
		Groups = false
		All_listings = false
		All_listings = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4all_listings)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	case "all_listings1":
		All_listings = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(query.Message.Chat.ID, query.Message.MessageID, CategoryInlineKeyboard4)
		_, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}

	}
	return nil
}
