package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

func (b *Bot) handleBodyStyle(query *tgbotapi.CallbackQuery, url *string) {
	switch query.Data {
	case "body_style_alfa_romeo_convertible":
		*url += "&carType=convertible"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_Coupe":
		*url += "&carType=coupe"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_hatchback":
		*url += "&carType=hatchback"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_carrier":
		*url += "&carType=minivan"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_saloon":
		*url += "&carType=sedan"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_estate":
		*url += "&carType=wagon"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_4x4":
		*url += "&carType=suv"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_van":
		*url += "&carType=truck"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_small_car":
		*url += "&carType=small_car"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_other":
		*url += "&carType=other_body_style"
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	case "body_style_alfa_romeo_all":
		go func() {
			ChBodyStyleAlfaRomeo <- query
		}()
	}
}
func (b *Bot) handleRooms(query *tgbotapi.CallbackQuery, url *string) {
	switch query.Data {
	case "rooms_1":
		*url += "&minRooms=1"
		go func() {
			ChRooms <- query
		}()
	case "rooms_2":
		*url += "&minRooms=2"
		go func() {
			ChRooms <- query
		}()
	case "rooms_3":
		*url += "&minRooms=3"
		go func() {
			ChRooms <- query
		}()
	case "rooms_4":
		*url += "&minRooms=4"
		go func() {
			ChRooms <- query
		}()
	case "rooms_5":
		*url += "&minRooms=5"
		go func() {
			ChRooms <- query
		}()
	case "rooms_6":
		*url += "&minRooms=6"
		go func() {
			ChRooms <- query
		}()
	case "rooms_all":
		go func() {
			ChRooms <- query
		}()
	}
}

func (b *Bot) handleTypeProperty(query *tgbotapi.CallbackQuery, url *string) {
	switch query.Data {
	case "type_property_flat_apartment":
		*url += "&propertyType=apartment-condo"
		go func() {
			ChTypeProperty <- query
		}()
	case "type_property_house":
		*url += "&propertyType=house"
		go func() {
			ChTypeProperty <- query
		}()
	case "type_property_room_only":
		*url += "&propertyType=private_room-shared_room"
		go func() {

			ChTypeProperty <- query
		}()
	case "type_property_townhouse":
		*url += "&propertyType=townhouse"
		go func() {
			ChTypeProperty <- query
		}()
	case "type_property_all":
		go func() {
			ChTypeProperty <- query
		}()
	}
}

func (b *Bot) handleCondition(query *tgbotapi.CallbackQuery, url *string) {
	switch query.Data {
	case "condition_new":
		*url += "?itemCondition=new"
		go func() {
			ChCondition <- query
		}()
	case "condition_used":
		*url += "?itemCondition=used_like_new"
		go func() {
			ChCondition <- query
		}()
	case "condition_good":
		*url += "?itemCondition=used_good"
		go func() {
			ChCondition <- query
		}()
	case "condition_fair":
		*url += "?itemCondition=used_fair"
		go func() {
			ChCondition <- query
		}()
	case "condition_all":
		go func() {
			ChCondition <- query
		}()
	}

}

func (b *Bot) handleBrand(query *tgbotapi.CallbackQuery, url *string) {
	switch query.Data {
	case "electronics_AMD":
		*url += "?contextual[brand]=874490466442333"
		go func() {
			ChBrand <- query
		}()
	case "electronics_ASUS":
		*url += "?contextual[brand]=2806224329413131"
		go func() {
			ChBrand <- query
		}()
	case "electronics_Apple":
		*url += "?contextual[brand]=177366900264876"
		go func() {
			ChBrand <- query
		}()
	case "electronics_Dre":
		*url += "?contextual[brand]=198430931580104"
		go func() {
			ChBrand <- query
		}()
	case "electronics_bose":
		*url += "?contextual[brand]=3294094674148107"
		go func() {
			ChBrand <- query
		}()
	case "electronics_canon":
		*url += "?contextual[brand]=2565877940322146"
		go func() {
			ChBrand <- query
		}()
	case "electronics_dell":
		*url += "?contextual[brand]=508407939787911"
		go func() {
			ChBrand <- query
		}()
	case "electronics_google":
		*url += "?contextual[brand]=2444794618966021"
		go func() {
			ChBrand <- query
		}()
	case "electronics_intel":
		*url += "?contextual[brand]=193214212244799"
		go func() {
			ChBrand <- query
		}()
	case "electronics_nintedo":
		*url += "?contextual[brand]=365440017874455"
		go func() {
			ChBrand <- query
		}()
	case "electronics_otterBox":
		*url += "?contextual[brand]=175149841088021"
		go func() {
			ChBrand <- query
		}()
	case "electronics_panasonic":
		*url += "?contextual[brand]=200020264439100"
		go func() {
			ChBrand <- query
		}()
	case "electronics_all":
		go func() {
			ChBrand <- query
		}()
	case "electronics_VIZIO":
		*url += "?contextual[brand]=312948339662051"
		go func() {
			ChBrand <- query
		}()
	}
}

func (b *Bot) handleBrandHome(query *tgbotapi.CallbackQuery, url *string) {
	switch query.Data {
	case "home_Improvement_craftsman":
		*url += "?contextual[brand]=287826089094325"
		go func() {
			ChHomeImprovement <- query
		}()
	case "home_Improvement_DEWALT":
		*url += "?contextual[brand]=624772404808190"
		go func() {
			ChHomeImprovement <- query
		}()
	case "home_Improvement_milwaukee":
		*url += "?contextual[brand]=2730829260488440"
		go func() {
			ChHomeImprovement <- query
		}()
	case "home_Improvement_snap_on":
		*url += "?contextual[brand]=814748759128299"
		go func() {
			ChHomeImprovement <- query
		}()
	case "home_Improvement_all":
		go func() {
			ChHomeImprovement <- query
		}()
	}
}
