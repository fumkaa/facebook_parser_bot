package main

import (
	"context"
	"facebook_marketplace_bot/internal/configs"
	database "facebook_marketplace_bot/internal/database/migration"
	"facebook_marketplace_bot/internal/parser"
	"facebook_marketplace_bot/internal/telegram"
	"fmt"
	"log"

	"github.com/go-sql-driver/mysql"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/jmoiron/sqlx"
)

func main() {
	config := configs.NewConfiguration()
	bot_api, err := tgbotapi.NewBotAPI(config.Token_bot)
	if err != nil {
		log.Fatal(err)
	}
	dbx, err := connDB(config)
	if err != nil {
		log.Fatalf("conn db err: %v", err)
	}
	db := database.NewStorage(dbx)

	ctx := context.Background()
	bot_api.Debug = true
	tg_bot := telegram.NewBot(bot_api, db, parser.NewParser(db))
	if err = tg_bot.Start(ctx); err != nil {
		log.Fatalf("can't start bot: %v", err)
	}
}

func connDB(config *configs.Configuration) (*sqlx.DB, error) {
	log.Print("connDB start")
	cnf := mysql.Config{
		User:              config.UserDB,
		Passwd:            config.PasswordDB,
		Net:               "tcp",
		Addr:              config.PortDB,
		DBName:            config.NameDB,
		InterpolateParams: false,
	}
	log.Print()
	db, err := sqlx.Open("mysql", cnf.FormatDSN())
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("[NewStorage]open db error: %w", err)
	}
	err = db.Ping()
	if err != nil {
		log.Print(err)
		return nil, fmt.Errorf("[NewStorage]connect db error: %w", err)
	}

	return db, nil
}
