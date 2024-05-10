package telegram

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"facebook_marketplace_bot/internal/parser"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

const (
	cat_vehicles          = "/category/vehicles"
	cat_propertyrentals   = "/category/propertyrentals"
	cat_free              = "/category/free"
	cat_toys              = "/category/toys"
	cat_instruments       = "/category/instruments"
	cat_home_improvements = "/category/home-improvements"
	cat_classifieds       = "/category/classifieds"
	cat_apparel           = "/category/apparel"
	cat_propertyforsale   = "/category/propertyforsale"
	cat_entertainment     = "/category/entertainment"
	cat_family            = "/category/family"
	cat_sports            = "/category/sports"
	cat_home              = "/category/home"
	cat_pets              = "/category/pets"
	cat_office_supplies   = "/category/office-supplies"
	cat_garden            = "/category/garden"
	cat_hobbies           = "/category/hobbies"
	cat_electronics       = "/category/electronics"
	cat_groups            = "/groups"
	cat_all_listings      = "/see_all_listings"
)

var (
	Vehicles        bool
	Propertyrentals bool
	Free            bool
	Toys            bool
	Instruments     bool

	Home_improvements bool
	Classifieds       bool
	Apparel           bool
	Propertyforsale   bool
	Entertainment     bool

	Family          bool
	Sports          bool
	Home            bool
	Pets            bool
	Office_supplies bool

	Garden       bool
	Hobbies      bool
	Electronics  bool
	Groups       bool
	All_listings bool
)
var (
	All                       bool
	Cars_and_lorries          bool
	Motorcycles               bool
	Powersports               bool
	Motorhomes_and_campers    bool
	Boats                     bool
	Commercial_and_industrial bool
	Trailers                  bool
	Other                     bool
)

const (
	baseUrl           = "https://www.facebook.com/marketplace"
	creation_time     = "&sortBy=creation_time_descend"
	creation_time_one = "?sortBy=creation_time_descend"
)

func (b *Bot) handlerCityInlineKeyboard(ctx context.Context, query *tgbotapi.CallbackQuery) error {
	return b.parser.ClickSelectCity(ctx, query.Data)
}

func (b *Bot) handlerRadiusInlineKeyboard(ctx context.Context, query *tgbotapi.CallbackQuery) error {
	return b.parser.SetRadius(ctx, query.Data)
}

var (
	url = baseUrl

	ChTypeVehicles       = make(chan *tgbotapi.CallbackQuery)
	ChExteriorColour     = make(chan *tgbotapi.CallbackQuery)
	ChInteriorColour     = make(chan *tgbotapi.CallbackQuery)
	ChTranssmission      = make(chan *tgbotapi.CallbackQuery)
	ChMake               = make(chan *tgbotapi.CallbackQuery)
	ChBodyStyleAlfaRomeo = make(chan *tgbotapi.CallbackQuery)
	ChRooms              = make(chan *tgbotapi.CallbackQuery)
	ChTypeProperty       = make(chan *tgbotapi.CallbackQuery)
	ChSquareMet          = make(chan *tgbotapi.CallbackQuery)
	ChCondition          = make(chan *tgbotapi.CallbackQuery)
	ChHomeImprovement    = make(chan *tgbotapi.CallbackQuery)
	ChBrand              = make(chan *tgbotapi.CallbackQuery)
	ChMsg                = make(chan tgbotapi.Message)
	MakeBreak            bool
	MakeNext             bool
	MakePrevious         bool
)
var (
	idAd   int
	sendAd []int
)

func (b *Bot) successfullCreateFilter(ctx context.Context, ChatID int64, url1 string) {
	log.Printf("!!!!!ready url: %s", url1)
	if err := b.db.AddMonitoringFilter(ctx, ID, url1); err != nil {
		log.Printf("AddMonitoringFilter error: %v", err)
		msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, добавьте фильтр еще раз")
		msg.ReplyMarkup = StartKeyboard
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Fatalf("[handleMessage]error send message: %v", err)
		}
		return
	}

	if err := b.FSM.Event(ctx, state_base); err != nil {
		msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
		msg.ReplyMarkup = StartKeyboard
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Fatalf("[handleMessage]error send message: %v", err)
		}
		return
	}

	if err := b.db.AddFilterFile(ctx, ID, CurrentFileName); err != nil {
		log.Printf("AddFilterFile error: %v", err)
		msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова добавить фильтр")
		msg.ReplyMarkup = StartKeyboard
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Fatalf("[handleMessage]error send message: %v", err)
		}
		if err := b.FSM.Event(ctx, state_base); err != nil {
			msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			return
		}
		return
	}
	if err := os.Rename(parser.Free_account+CurrentFileName, parser.Work_account+CurrentFileName); err != nil {
		log.Printf("Rename error: %v", err)
		msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова добавить фильтр")
		msg.ReplyMarkup = StartKeyboard
		_, err = b.bot.Send(msg)
		if err != nil {
			log.Fatalf("[handleMessage]error send message: %v", err)
		}
		if err := b.FSM.Event(ctx, state_base); err != nil {
			msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			return
		}
		return
	}
	msg := tgbotapi.NewMessage(ChatID, "Успешно создан фильтр!")
	msg.ReplyMarkup = StartKeyboard
	_, err := b.bot.Send(msg)
	if err != nil {
		log.Printf("error send message: %v", err)
		return
	}
	CurrentFileName = ""
	defer Cancel1()
	defer Cancel2()
	defer chromedp.Cancel(Ctxt)
	for {
		log.Print("MONITORING......")

		_, err := b.db.MonitoringByIDFilter(ctx, ID)
		if err == database.ErrNoRows {
			log.Printf("END MONITORING  BY ID %d", ID)
			break
		}
		if err != nil && err != database.ErrNoRows {
			log.Printf("MonitoringByIDFilter error: %v", err)
			return
		}
		var (
			nodes   []*cdp.Node
			elUrlAd []string
		)
		err = chromedp.Run(Ctxt,
			chromedp.Navigate(url1),
		)
		if err != nil {
			log.Printf("[monitoring]run error: %v", err)
			msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			if err := b.FSM.Event(ctx, state_base); err != nil {
				msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				return
			}
			return
		}
		log.Print("Navigate")
		err = chromedp.Run(Ctxt,
			chromedp.Nodes(`a[class="x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 xggy1nq x1a2a7pz x1heor9g xt0b8zv x1hl2dhg x1lku1pv"]`,
				&nodes, chromedp.ByQuery),
		)
		if err != nil {
			log.Printf("[monitoring]Nodes error: %v", err)
			msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			if err := b.FSM.Event(ctx, state_base); err != nil {
				msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				return
			}
			return
		}
		log.Print("Nodes")
		for _, node := range nodes {
			urlAd := node.AttributeValue("href")
			elUrlAd = strings.Split(urlAd, "/")
		}
		log.Print("for")
		curId, err := strconv.Atoi(elUrlAd[3])
		if err != nil {
			log.Printf("convert elUrlAd error: %v", err)
			msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
			msg.ReplyMarkup = StartKeyboard
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Fatalf("[handleMessage]error send message: %v", err)
			}
			if err := b.FSM.Event(ctx, state_base); err != nil {
				msg := tgbotapi.NewMessage(ChatID, "Произошла ошибка, попробуйте снова")
				msg.ReplyMarkup = StartKeyboard
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Fatalf("[handleMessage]error send message: %v", err)
				}
				return
			}
			return
		}
		if idAd == 0 {
			idAd = curId

			time.Sleep(5 * time.Second)
			continue
		}
		if idAd == curId {
			log.Print("id ad not change")
			log.Printf("cur id ad: %d\nold id ad: %d", curId, idAd)
			time.Sleep(10 * time.Second)
			continue
		} else if idAd != curId {
			if sendAd == nil {
				msg = tgbotapi.NewMessage(ChatID, "Найдено новое объявление: https://www.facebook.com/marketplace/item/"+strconv.Itoa(curId))
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				sendAd = append(sendAd, curId)
				idAd = curId
				continue
			}
			for _, sndAd := range sendAd {
				if curId == sndAd {
					idAd = curId
					continue
				}
			}
			msg = tgbotapi.NewMessage(ChatID, "Найдено новое объявление: https://www.facebook.com/marketplace/item/"+strconv.Itoa(curId))
			_, err = b.bot.Send(msg)
			if err != nil {
				log.Printf("error send message: %v", err)
			}
			sendAd = append(sendAd, curId)
			idAd = curId
			continue
		}
	}
}

func (b *Bot) handlerCategoryInlineKeyboard(ctx context.Context, query *tgbotapi.CallbackQuery) error {

	switch query.Data {
	case "make_next":
		MakeNext = true
	case "make_previous":
		MakePrevious = true
	case "make_all":
		go func() {
			ChMake <- query
		}()
	case "make_alfa_romeo":
		url += "?make=2070821889637770"
		go func() {
			ChMake <- query
		}()
	case "make_aston_martin":
		url += "?make=341553276467626"
		go func() {
			ChMake <- query
		}()
	case "make_audi":
		url += "?make=1274042129420222"
		go func() {
			ChMake <- query
		}()
	case "make_BMW":
		url += "?make=313316249374727"
		go func() {
			ChMake <- query
		}()
	case "make_bentley":
		url += "?make=382125202569476"
		go func() {
			ChMake <- query
		}()
	case "make_buick":
		url += "?make=436791410393181"
		go func() {
			ChMake <- query
		}()
	case "make_cadillac":
		url += "?make=489276414939966"
		go func() {
			ChMake <- query
		}()
	case "make_chevrolet":
		url += "?make=1914016008726893"
		go func() {
			ChMake <- query
		}()
	case "make_chrysler":
		url += "?make=398368117562414"
		go func() {
			ChMake <- query
		}()
	case "make_daewoo":
		url += "?make=636736676747295"
		go func() {
			ChMake <- query
		}()
	case "make_dodge":
		url += "?make=402915273826151"
		go func() {
			ChMake <- query
		}()
	case "make_ferrari":
		url += "?make=2233936113511813"
		go func() {
			ChMake <- query
		}()
	case "make_FIAT":
		url += "?make=973850736337551"
		go func() {
			ChMake <- query
		}()
	case "make_ford":
		url += "?make=297354680962030"
		go func() {
			ChMake <- query
		}()
	case "make_honda":
		url += "?make=308436969822020"
		go func() {
			ChMake <- query
		}()
	case "make_hummer":
		url += "?make=351240532159178"
		go func() {
			ChMake <- query
		}()
	case "make_hyundai":
		url += "?make=590755841400441"
		go func() {
			ChMake <- query
		}()
	case "make_INFINITI":
		url += "?make=1361484827327051"
		go func() {
			ChMake <- query
		}()
	case "make_isuzu":
		url += "?make=1865279260243853"
		go func() {
			ChMake <- query
		}()
	case "make_jaguar":
		url += "?make=2127086467326917"
		go func() {
			ChMake <- query
		}()
	case "make_jeep":
		url += "?make=408221723080125"
		go func() {
			ChMake <- query
		}()
	case "make_kia":
		url += "?make=417670842327686"
		go func() {
			ChMake <- query
		}()
	case "make_lamborghini":
		url += "?make=622700624835375"
		go func() {
			ChMake <- query
		}()
	case "make_land_rover":
		url += "?make=264916861051142"
		go func() {
			ChMake <- query
		}()
	case "make_lotus":
		url += "?make=764630980585596"
		go func() {
			ChMake <- query
		}()
	case "ake_MINI":
		url += "?make=357388888199059"
		go func() {
			ChMake <- query
		}()
	case "ake_maserati":
		url += "?make=382869215628928"
		go func() {
			ChMake <- query
		}()
	case "make_mazda":
		url += "?make=410067716491465"
		go func() {
			ChMake <- query
		}()
	case "make_McLaren":
		url += "?make=258881574989122"
		go func() {
			ChMake <- query
		}()
	case "make_mercedes_menz":
		url += "?make=391196981458827"
		go func() {
			ChMake <- query
		}()
	case "make_mitsubishi":
		url += "?make=343841119808433"
		go func() {
			ChMake <- query
		}()
	case "make_nissan":
		url += "?make=2621742507840619"
		go func() {
			ChMake <- query
		}()
	case "make_pontiac":
		url += "?make=2034940493472692"
		go func() {
			ChMake <- query
		}()
	case "make_porsche":
		url += "?make=2233878893526787"
		go func() {
			ChMake <- query
		}()
	case "make_rolls_royce":
		url += "?make=1199922706845728"
		go func() {
			ChMake <- query
		}()
	case "make_saab":
		url += "?make=327726657883762"
		go func() {
			ChMake <- query
		}()
	case "make_smart":
		url += "?make=573131263170521"
		go func() {
			ChMake <- query
		}()
	case "make_subaru":
		url += "?make=2571870739551112"
		go func() {
			ChMake <- query
		}()
	case "make_suzuki":
		url += "?make=2449096741789800"
		go func() {
			ChMake <- query
		}()
	case "make_tesla":
		url += "?make=621439278281940"
		go func() {
			ChMake <- query
		}()
	case "make_toyota":
		url += "?make=2318041991806363"
		go func() {
			ChMake <- query
		}()
	case "make_volkswagen":
		url += "?make=523665818157652"
		go func() {
			ChMake <- query
		}()
	case "make_volvo":
		url += "?make=367577657422824"
		go func() {
			ChMake <- query
		}()
	case "automatic":
		url += "&transmissionType=automatic"
		go func() {
			ChTranssmission <- query
		}()
	case "manual":
		url += "&transmissionType=manual"
		go func() {
			ChTranssmission <- query
		}()
	case "transmisson_all":
		go func() {
			ChTranssmission <- query
		}()
	case "vehicles_type_all":
		All = true
		Cars_and_lorries = false
		Motorcycles = false
		Powersports = false
		Motorhomes_and_campers = false
		Boats = false
		Commercial_and_industrial = false
		Trailers = false
		Other = false
		go func() {
			ChTypeVehicles <- query
		}()
	case "vehicles_type_cars_and_lorries":
		All = false
		Cars_and_lorries = true
		Motorcycles = false
		Powersports = false
		Motorhomes_and_campers = false
		Boats = false
		Commercial_and_industrial = false
		Trailers = false
		Other = false
		url += "&topLevelVehicleType=car_truck"
		go func() {
			ChTypeVehicles <- query
		}()
	case "vehicles_type_motorcycles":
		All = false
		Cars_and_lorries = false
		Motorcycles = true
		Powersports = false
		Motorhomes_and_campers = false
		Boats = false
		Commercial_and_industrial = false
		Trailers = false
		Other = false
		url += "&topLevelVehicleType=motorcycle"
		go func() {
			ChTypeVehicles <- query
		}()
	case "vehicles_type_powersports":
		All = false
		Cars_and_lorries = false
		Motorcycles = false
		Powersports = true
		Motorhomes_and_campers = false
		Boats = false
		Commercial_and_industrial = false
		Trailers = false
		Other = false
		url += "&topLevelVehicleType=powersport"
		go func() {
			ChTypeVehicles <- query
		}()
	case "vehicles_type_motorhomes_and_campers":
		All = false
		Cars_and_lorries = false
		Motorcycles = false
		Powersports = false
		Motorhomes_and_campers = true
		Boats = false
		Commercial_and_industrial = false
		Trailers = false
		Other = false
		url += "&topLevelVehicleType=rv_camper"
		go func() {
			ChTypeVehicles <- query
		}()
	case "vehicles_type_boats":
		All = false
		Cars_and_lorries = false
		Motorcycles = false
		Powersports = false
		Motorhomes_and_campers = false
		Boats = true
		Commercial_and_industrial = false
		Trailers = false
		Other = false
		url += "&topLevelVehicleType=boat"
		go func() {
			ChTypeVehicles <- query
		}()
	case "vehicles_type_commercial_and_industrial":
		All = false
		Cars_and_lorries = false
		Motorcycles = false
		Powersports = false
		Motorhomes_and_campers = false
		Boats = false
		Commercial_and_industrial = true
		Trailers = false
		Other = false
		url += "&topLevelVehicleType=commercial"
		go func() {
			ChTypeVehicles <- query
		}()
	case "vehicles_type_trailers":
		All = false
		Cars_and_lorries = false
		Motorcycles = false
		Powersports = false
		Motorhomes_and_campers = false
		Boats = false
		Commercial_and_industrial = false
		Trailers = true
		Other = false
		url += "&topLevelVehicleType=trailer"
		go func() {
			ChTypeVehicles <- query
		}()
	case "vehicles_type_other":
		All = false
		Cars_and_lorries = false
		Motorcycles = false
		Powersports = false
		Motorhomes_and_campers = false
		Boats = false
		Commercial_and_industrial = false
		Trailers = false
		Other = true
		url += "&topLevelVehicleType=other"
		go func() {
			ChTypeVehicles <- query
		}()
	case "exterior_black":
		url += "&vehicleExteriorColors=black"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_charcoal":
		url += "&vehicleExteriorColors=charcoal"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_grey":
		url += "&vehicleExteriorColors=grey"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_silver":
		url += "&vehicleExteriorColors=silver"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_white":
		url += "&vehicleExteriorColors=white"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_off_white":
		url += "&vehicleExteriorColors=off_white"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_tan":
		url += "&vehicleExteriorColors=tan"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_beige":
		url += "&vehicleExteriorColors=beige"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_yellow":
		url += "&vehicleExteriorColors=yellow"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_gold":
		url += "&vehicleExteriorColors=gold"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_brown":
		url += "&vehicleExteriorColors=brown"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_orange":
		url += "&vehicleExteriorColors=orange"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_red":
		url += "&vehicleExteriorColors=red"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_burgundy":
		url += "&vehicleExteriorColors=burgundy"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_pink":
		url += "&vehicleExteriorColors=pink"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_purple":
		url += "&vehicleExteriorColors=purple"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_blue":
		url += "&vehicleExteriorColors=blue"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_turquoise":
		url += "&vehicleExteriorColors=turquoise"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_green":
		url += "&vehicleExteriorColors=green"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_other":
		url += "&vehicleExteriorColors=other"
		go func() {
			ChExteriorColour <- query
		}()
	case "exterior_nothing":
		go func() {
			ChExteriorColour <- query
		}()
	case "interior_black":
		url += "&vehicleInteriorColors=black"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_charcoal":
		url += "&vehicleInteriorColors=charcoal"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_grey":
		url += "&vehicleInteriorColors=grey"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_silver":
		url += "&vehicleInteriorColors=silver"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_white":
		url += "&vehicleInteriorColors=white"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_off_white":
		url += "&vehicleInteriorColors=off_white"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_tan":
		url += "&vehicleInteriorColors=tan"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_beige":
		url += "&vehicleInteriorColors=beige"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_yellow":
		url += "&vehicleInteriorColors=yellow"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_gold":
		url += "&vehicleInteriorColors=gold"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_brown":
		url += "&vehicleInteriorColors=brown"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_orange":
		url += "&vehicleInteriorColors=orange"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_red":
		url += "&vehicleInteriorColors=red"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_burgundy":
		url += "&vehicleInteriorColors=burgundy"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_pink":
		url += "&vehicleInteriorColors=pink"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_purple":
		url += "&vehicleInteriorColors=purple"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_blue":
		url += "&vehicleInteriorColors=blue"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_turquoise":
		url += "&vehicleInteriorColors=turquoise"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_green":
		url += "&vehicleInteriorColors=green"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_other":
		url += "&vehicleInteriorColors=other"
		go func() {
			ChInteriorColour <- query
		}()
	case "interior_nothing":
		go func() {
			ChInteriorColour <- query
		}()
	case "confirm":
		log.Print("confirm")
		log.Printf("vehicles: %t", Vehicles)
		log.Printf("Propertyrentals: %t", Propertyrentals)
		log.Printf("Free: %t", Free)
		log.Printf("Toys: %t", Toys)
		log.Printf("Instruments: %t", Instruments)
		log.Printf("Home_improvements: %t", Home_improvements)
		log.Printf("Classifieds: %t", Classifieds)
		log.Printf("Apparel: %t", Apparel)
		log.Printf("Propertyforsale: %t", Propertyforsale)
		log.Printf("Entertainment: %t", Entertainment)
		log.Printf("Family: %t", Family)
		log.Printf("Sports: %t", Sports)
		log.Printf("Home: %t", Home)
		log.Printf("Pets: %t", Pets)
		log.Printf("Office_supplies: %t", Office_supplies)
		log.Printf("Garden: %t", Garden)
		log.Printf("Hobbies: %t", Hobbies)
		log.Printf("Electronics: %t", Electronics)
		log.Printf("Groups: %t", Groups)
		log.Printf("All_listings: %t", All_listings)
		go func() {
			for {
				if Vehicles || Propertyrentals || Toys || Instruments || Home_improvements || Classifieds || Apparel || Propertyforsale || Entertainment || Family || Sports || Home || Pets || Office_supplies || Garden || Hobbies || Electronics {
					<-ChInputPrice
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Введите минимальную и максимальную цену (в валюте, которая в том городе, который вы выбрали ранее, например если этот город в США, то валюта будет доллары) через запятую:\n111, 99999\nНеважно в каком порядке, большее число будет считаться как максимальное")
					_, err := b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					if err := b.FSM.Event(ctx, state_input_price); err != nil {
						msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
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
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Введите минимальный и максимальный год через запятую:\n2020, 2024\nНеважно в каком порядке, больший год будет считаться как максимальный")
					_, err := b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					if err := b.FSM.Event(ctx, state_input_year); err != nil {
						msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
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
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Введите минимальные и максимальные квадратные метры через запятую:\n100, 3000\nНеважно в каком порядке, большее число будет считаться как максимальный")
					_, err := b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					if err := b.FSM.Event(ctx, state_input_square_meters); err != nil {
						msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова")
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
						msg := tgbotapi.NewMessage(ChatID, "Введите минимальный и максимальный пробег (в той метрической системе, которая в выбранном вами городе, например если город в США, то тогда пробег измеряется в милях) через запятую:\n100, 5000\nНеважно в каком порядке, больший пробег будет считаться как максимальный")
						_, err := b.bot.Send(msg)
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
		}(query.Message.Chat.ID)
		switch {
		case Vehicles:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "vehicles"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}
					return
				}
				url += cat_vehicles
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice)
				log.Printf("vehicles url: %s", url)
				ChInputYear <- query
				log.Print("set ChInputYear")
				minYear := <-ChMinYear
				maxYear := <-ChMaxYear
				url += fmt.Sprintf("&maxYear=%s&minYear=%s", maxYear, minYear)
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите тип транспортного средства:")
				msg.ReplyMarkup = VehiclesTypeInlineKeyboard
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChTypeVehicles
				log.Print("set ChTypeVehicles")
				switch {
				case All:

					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите фирму транспортного средства:")
					msg.ReplyMarkup = MakeInlineKeyboard
					msg1, err := b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					MakeBreak = false
					go func() {
						if err := b.handleMake(msg1); err != nil {
							log.Printf("handleMake error: %v", err)
							return
						}
					}()

					<-ChMake
					MakeBreak = true
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите стиль кузова:")
					msg.ReplyMarkup = BodyStyleAlfaRomeoInlineKB
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}

					<-ChBodyStyleAlfaRomeo

					ChInputMill <- query
					minMill := <-ChMinMill
					maxMill := <-ChMaxMill
					url += fmt.Sprintf("&maxMileage=%s&minMileage=%s", maxMill, minMill) + creation_time
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите цвет салона транспортного средства:")
					msg.ReplyMarkup = InteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChInteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите тип коропки передач:")
					msg.ReplyMarkup = TransmissonInlineKB
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChTranssmission
					url += "&exact=false"
					log.Printf("url: %s", url)
					go func(url1 string) { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url1) }(url)
					time.Sleep(1 * time.Second)
					url = baseUrl
				case Cars_and_lorries:

					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите фирму транспортного средства:")
					msg.ReplyMarkup = MakeInlineKeyboard
					msg1, err := b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					MakeBreak = false
					go func() {
						if err := b.handleMake(msg1); err != nil {
							log.Printf("handleMake error: %v", err)
							return
						}
					}()
					<-ChMake
					MakeBreak = true
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите стиль кузова:")
					msg.ReplyMarkup = BodyStyleAlfaRomeoInlineKB
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}

					<-ChBodyStyleAlfaRomeo
					ChInputMill <- query
					minMill := <-ChMinMill
					maxMill := <-ChMaxMill
					url += fmt.Sprintf("&maxMileage=%s&minMileage=%s", maxMill, minMill) + creation_time
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите цвет салона транспортного средства:")
					msg.ReplyMarkup = InteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChInteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите тип коропки передач:")
					msg.ReplyMarkup = TransmissonInlineKB
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChTranssmission
					url += "&exact=false"
					log.Printf("url: %s", url)
					go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
					time.Sleep(1 * time.Second)
					url = baseUrl
				case Motorcycles:
					ChInputMill <- query
					minMill := <-ChMinMill
					maxMill := <-ChMaxMill
					url += fmt.Sprintf("&maxMileage=%s&minMileage=%s", maxMill, minMill) + creation_time
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите тип коропки передач:")
					msg.ReplyMarkup = TransmissonInlineKB
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChTranssmission
					url += "&exact=false"
					log.Printf("url: %s", url)

					go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
					time.Sleep(1 * time.Second)
					url = baseUrl
				case Powersports:
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите цвет салона транспортного средства:")
					msg.ReplyMarkup = InteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChInteriorColour
					url += creation_time + "&exact=false"
					log.Printf("url: %s", url)

					go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
					time.Sleep(1 * time.Second)
					url = baseUrl
				case Motorhomes_and_campers:
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите цвет салона транспортного средства:")
					msg.ReplyMarkup = InteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChInteriorColour
					url += creation_time + "&exact=false"
					log.Printf("url: %s", url)

					go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
					time.Sleep(1 * time.Second)
					url = baseUrl
				case Boats:
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите цвет салона транспортного средства:")
					msg.ReplyMarkup = InteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChInteriorColour
					log.Print("!!!!!!!!!click InteriorColour")
					url += creation_time + "&exact=false"
					log.Printf("url: %s", url)

					go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
					time.Sleep(1 * time.Second)
					url = baseUrl
				case Commercial_and_industrial:
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите цвет салона транспортного средства:")
					msg.ReplyMarkup = InteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChInteriorColour
					url += creation_time + "&exact=false"
					log.Printf("url: %s", url)

					go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
					time.Sleep(1 * time.Second)
					url = baseUrl
				case Trailers:
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите цвет салона транспортного средства:")
					msg.ReplyMarkup = InteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChInteriorColour
					url += creation_time + "&exact=false"
					log.Printf("url: %s", url)

					go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
					time.Sleep(1 * time.Second)
					url = baseUrl
				case Other:
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите внешний цвет транспортного средства:")
					msg.ReplyMarkup = ExteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChExteriorColour
					msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите цвет салона транспортного средства:")
					msg.ReplyMarkup = InteriorColourInlineKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Printf("error send message: %v", err)
					}
					<-ChInteriorColour
					url += creation_time + "&exact=false"
					log.Printf("url: %s", url)

					go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
					time.Sleep(1 * time.Second)
					url = baseUrl
				}
			}()
		case Propertyrentals:
			go func() {

				if err := b.db.AddCategoryFilter(ctx, ID, "propertyrentals"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_propertyrentals
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice)
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите количество комнат:")
				msg.ReplyMarkup = RoomsInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChRooms
				msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите тип недвижимости, сдаваемой в аренду:")
				msg.ReplyMarkup = TypePropertyInlineKB
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}

				<-ChTypeProperty
				ChSquareMet <- query
				minMet := <-ChMinMet
				maxMet := <-ChMaxMet
				log.Print("get minMet maxMet!!!")
				url += fmt.Sprintf("&maxAreaSize=%s&minAreaSize=%s", maxMet, minMet) + creation_time
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Free:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "free"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_free
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Вы выбрали категорию free")
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				log.Printf("url: %s", url)
				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Toys:
			go func() {
				log.Print("Toys")
				if err := b.db.AddCategoryFilter(ctx, ID, "toys"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_toys + creation_time
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice)
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Instruments:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "instruments"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_instruments
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Home_improvements:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "home improvements"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_home_improvements
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите бренд:")
				msg.ReplyMarkup = HomeImprovementInlineKB
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChHomeImprovement
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Classifieds:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "classifieds"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_classifieds
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Apparel:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "apparel"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_apparel
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Propertyforsale:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "property for sale"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_propertyforsale
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Entertainment:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "entertainment"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_entertainment
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Family:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "family"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_family
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Sports:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "sports"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_sports
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Home:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "home"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_home
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Pets:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "pets"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_pets
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Office_supplies:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "office supplies"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_office_supplies
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Garden:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "garden"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_garden
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Hobbies:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "hobbies"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_hobbies
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Electronics:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "electronics"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_electronics
				ChInputPrice <- query
				log.Print("wait min max price")
				minPrice := <-ChMinPrice
				maxPrice := <-ChMaxPrice
				url += fmt.Sprintf("?minPrice=%s&maxPrice=%s", minPrice, maxPrice) + creation_time
				log.Printf("vehicles url: %s", url)
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите состояние предмета:")
				msg.ReplyMarkup = ConditionInlineKB
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChCondition
				msg = tgbotapi.NewMessage(query.Message.Chat.ID, "Выберите бренд:")
				msg.ReplyMarkup = BrandInlineKB
				_, err = b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}
				<-ChBrand
				url += "&exact=false"
				log.Printf("url: %s", url)

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case Groups:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "groups"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_groups + creation_time_one
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Вы выбрали категорию groups")
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		case All_listings:
			go func() {
				if err := b.db.AddCategoryFilter(ctx, ID, "all listings"); err != nil {
					log.Printf("AddCategoryFilter error: %v", err)
					msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Произошла ошибка, попробуйте снова добавить фильтр")
					msg.ReplyMarkup = StartKeyboard
					_, err = b.bot.Send(msg)
					if err != nil {
						log.Fatalf("[handleMessage]error send message: %v", err)
					}

					return
				}
				url += cat_all_listings + creation_time_one
				msg := tgbotapi.NewMessage(query.Message.Chat.ID, "Вы выбрали категорию all listings")
				_, err := b.bot.Send(msg)
				if err != nil {
					log.Printf("error send message: %v", err)
				}

				go func() { b.successfullCreateFilter(ctx, query.Message.Chat.ID, url) }()
				time.Sleep(1 * time.Second)
				url = baseUrl
			}()
		default:
			msg := tgbotapi.NewCallback(query.ID, replyNothingSelect)
			_, err := b.bot.Request(msg)
			if err != nil {
				log.Fatalf("error send Request: %v", err)
			}
		}
	}
	if err := b.handleCategory(query); err != nil {
		return err
	}
	b.handleBodyStyle(query, &url)
	b.handleRooms(query, &url)
	b.handleTypeProperty(query, &url)
	b.handleCondition(query, &url)
	b.handleBrand(query, &url)
	b.handleBrandHome(query, &url)
	return nil
}
