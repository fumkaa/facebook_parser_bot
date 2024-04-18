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
		case home_improvements:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2home_improvements)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case classifieds:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2classifieds)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case apparel:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2apparel)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case propertyforsale:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2propertyforsale)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case entertainment:
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
		case family:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3family)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case sports:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3sports)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case home:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3home)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case pets:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3pets)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case office_supplies:
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
		case garden:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4garden)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case hobbies:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4hobbies)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case electronics:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4electronics)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case groups:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4groups)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case all_listings:
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
		case vehicles:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1vehicles)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case propertyrentals:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1propertyrentals)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case free:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1free)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case toys:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1toys)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case instruments:
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
		case home_improvements:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2home_improvements)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case classifieds:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2classifieds)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case apparel:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2apparel)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case propertyforsale:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2propertyforsale)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case entertainment:
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
		case family:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3family)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case sports:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3sports)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case home:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3home)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case pets:
			editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3pets)
			msg, err := b.bot.Send(editInlineKB)
			if err != nil {
				return fmt.Errorf("send message error: %w", err)
			}
			message = msg
		case office_supplies:
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
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false

		vehicles = true
		log.Printf("vehicles: %t", vehicles)
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1vehicles)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "vehicles1":
		vehicles = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "propertyrentals":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		propertyrentals = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1propertyrentals)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "propertyrentals1":
		propertyforsale = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "free":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		free = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1free)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "free1":
		free = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "toys":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		toys = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1toys)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "toys1":
		toys = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "instruments":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		instruments = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1instruments)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "instruments1":
		instruments = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard1)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg

	case "home_improvements":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		home_improvements = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2home_improvements)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "home_improvements1":
		home_improvements = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "classifieds":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		classifieds = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2classifieds)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "classifieds1":
		classifieds = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "apparel":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		apparel = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2apparel)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "apparel1":
		apparel = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "propertyforsale":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		propertyforsale = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2propertyforsale)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "propertyforsale1":
		propertyforsale = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "entertainment":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		entertainment = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2entertainment)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "entertainment1":
		entertainment = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard2)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg

	case "family":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		family = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3family)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "family1":
		family = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "sports":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		sports = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3sports)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "sports1":
		sports = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "home":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		home = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3home)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "home1":
		home = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "pets":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		pets = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3pets)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "pets1":
		pets = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "office_supplies":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		office_supplies = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3office_supplies)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "office_supplies1":
		office_supplies = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard3)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg

	case "garden":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		garden = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4garden)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "garden1":
		garden = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "hobbies":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		hobbies = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4hobbies)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "hobbies1":
		hobbies = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "electronics":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		electronics = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4electronics)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "electronics1":
		electronics = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "groups":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		groups = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4groups)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "groups1":
		groups = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "all_listings":
		vehicles = false
		propertyrentals = false
		free = false
		toys = false
		instruments = false

		home_improvements = false
		classifieds = false
		apparel = false
		propertyforsale = false
		entertainment = false

		family = false
		sports = false
		home = false
		pets = false
		office_supplies = false

		garden = false
		hobbies = false
		electronics = false
		groups = false
		all_listings = false
		all_listings = true
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4all_listings)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg
	case "all_listings1":
		all_listings = false
		editInlineKB := tgbotapi.NewEditMessageReplyMarkup(message.Chat.ID, message.MessageID, CategoryInlineKeyboard4)
		msg, err := b.bot.Send(editInlineKB)
		if err != nil {
			return fmt.Errorf("send message error: %w", err)
		}
		message = msg

	}
	return nil
}
