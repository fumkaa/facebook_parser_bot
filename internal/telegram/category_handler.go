package telegram

import (
	"fmt"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (b *Bot) handleCategory(query *tgbotapi.CallbackQuery, message tgbotapi.Message) error {
	switch query.Data {
	case "next1":
		switch {
		case Home_improvements:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2home_improvements)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Classifieds:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2classifieds)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Apparel:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2apparel)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Propertyforsale:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2propertyforsale)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Entertainment:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2entertainment)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		}
	case "next2":
		switch {
		case Family:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3family)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Sports:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3sports)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Home:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3home)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Pets:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3pets)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Office_supplies:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3office_supplies)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		}
	case "next3":
		switch {
		case Garden:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4garden)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Hobbies:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4hobbies)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Electronics:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4electronics)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Groups:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4groups)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case All_listings:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4all_listings)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		}
	case "previous1":
		switch {
		case Vehicles:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1vehicles)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Propertyrentals:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1propertyrentals)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Free:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1free)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Toys:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1toys)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Instruments:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1instruments)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		}
	case "previous2":
		switch {
		case Home_improvements:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2home_improvements)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Classifieds:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2classifieds)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Apparel:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2apparel)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Propertyforsale:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2propertyforsale)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Entertainment:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2entertainment)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		}
	case "previous3":
		switch {
		case Family:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3family)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Sports:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3sports)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Home:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3home)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Pets:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3pets)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case Office_supplies:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3office_supplies)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		default:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1vehicles)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "vehicles1":
		Vehicles = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1propertyrentals)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "propertyrentals1":
		Propertyforsale = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1free)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "free1":
		Free = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1toys)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "toys1":
		Toys = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1instruments)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "instruments1":
		Instruments = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg

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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2home_improvements)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "home_improvements1":
		Home_improvements = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2classifieds)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "classifieds1":
		Classifieds = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2apparel)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "apparel1":
		Apparel = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2propertyforsale)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "propertyforsale1":
		Propertyforsale = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2entertainment)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "entertainment1":
		Entertainment = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg

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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3family)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "family1":
		Family = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3sports)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "sports1":
		Sports = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3home)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "home1":
		Home = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3pets)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "pets1":
		Pets = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3office_supplies)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "office_supplies1":
		Office_supplies = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg

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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4garden)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "garden1":
		Garden = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4hobbies)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "hobbies1":
		Hobbies = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4electronics)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "electronics1":
		Electronics = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4groups)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "groups1":
		Groups = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
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
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4all_listings)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "all_listings1":
		All_listings = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg

	}
	return nil
}
