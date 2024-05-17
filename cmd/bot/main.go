package main

import (
	"context"
	"facebook_marketplace_bot/internal/configs"
	database "facebook_marketplace_bot/internal/database/migration"
	"facebook_marketplace_bot/internal/parser"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	"github.com/go-sql-driver/mysql"

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

// var (
// 	idAd   int
// 	SendAd []int64
// )

// func test() {
// 	config := configs.NewConfiguration()
// 	dbx, _ := connDB(config)
// 	db := database.NewStorage(dbx)
// 	parser := parser.NewParser(db)
// 	datas, _ := parser.GetData()
// 	proxy := fmt.Sprintf("http://%s", datas[0].Datas.IpPortPX)
// 	opts := append(chromedp.DefaultExecAllocatorOptions[:],
// 		chromedp.ProxyServer(proxy),
// 		chromedp.WindowSize(1900, 1080), // init with a desktop view
// 		chromedp.Flag("enable-automation", false),
// 		chromedp.Flag("headless", false),
// 	)
// 	log.Printf("ip port proxy: %v", datas[0].Datas.IpPortPX)
// 	var ctxChr context.Context
// 	ctxChr, Cancel1 := chromedp.NewExecAllocator(context.Background(), opts...)
// 	defer Cancel1()
// 	Ctxt, Cancel2 := chromedp.NewContext(ctxChr) // chromedp.WithDebugf(log.Printf),
// 	defer Cancel2()
// 	log.Print("!!!settings!!!")
// 	parser.Settings(Ctxt, datas[0])
// 	defer chromedp.Cancel(Ctxt)
// 	for {
// 		log.Print("MONITORING......")
// 		var (
// 			nodes []*cdp.Node
// 			// res     *runtime.RemoteObject
// 			// res     string
// 			elUrlAd []string
// 		)
// 		err := chromedp.Run(Ctxt,
// 			chromedp.Navigate("https://facebook.com/marketplace/"),
// 		)
// 		if err != nil {
// 			log.Printf("[monitoring]run error: %v", err)
// 			return
// 		}
// 		log.Print("Navigate")
// 		err = chromedp.Run(Ctxt,
// 			// chromedp.EvaluateAsDevTools(`document.querySelector("img").closest("a");`, &res /*, chromedp.EvalAsValue*/),
// 			chromedp.Nodes(`img`, &nodes, chromedp.ByQuery),
// 		)
// 		log.Printf("!!!!!!!!nodes1: %#v", nodes)
// 		if err != nil {
// 			log.Printf("[monitoring]Nodes error: %v", err)
// 			return
// 		}
// 		log.Print("Nodes")
// 		for _, node := range nodes {
// 			parent := node.Parent
// 			for {
// 				if parent.LocalName == "div" {
// 					log.Printf("parent.LocalName == div")
// 					parent = parent.Parent
// 				} else if parent.LocalName == "a" {
// 					log.Printf("parent.LocalName == a")
// 					break
// 				}
// 			}
// 			log.Printf("parent: %v", parent)
// 			urlAd := parent.AttributeValue("href")
// 			elUrlAd = strings.Split(urlAd, "/")
// 		}

// 		log.Print("for")
// 		curId, err := strconv.Atoi(elUrlAd[3])
// 		if err != nil {
// 			log.Printf("convert elUrlAd error: %v", err)
// 			return
// 		}
// 		if idAd == 0 {
// 			idAd = curId

// 			time.Sleep(5 * time.Second)
// 			continue
// 		}
// 		if idAd == curId {
// 			log.Print("id ad not change")
// 			log.Printf("cur id ad: %d\nold id ad: %d", curId, idAd)
// 			time.Sleep(10 * time.Second)
// 			continue
// 		} else if idAd != curId {
// 			if SendAd == nil {
// 				log.Println("SendAd == nil")
// 				SendAd = append(SendAd, int64(curId))
// 				idAd = curId
// 				continue
// 			}
// 			for _, sndAd := range SendAd {
// 				if int64(curId) == sndAd {
// 					idAd = curId
// 					continue
// 				}
// 			}
// 			log.Println("find new ad")
// 			SendAd = append(SendAd, int64(curId))
// 			idAd = curId
// 			continue
// 		}
// 	}
// }
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
