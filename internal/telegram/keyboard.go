package telegram

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

var RadiusInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1", "1"),
		tgbotapi.NewInlineKeyboardButtonData("2", "2"),
		tgbotapi.NewInlineKeyboardButtonData("5", "5"),
		tgbotapi.NewInlineKeyboardButtonData("10", "10"),
		tgbotapi.NewInlineKeyboardButtonData("20", "20"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("40", "40"),
		tgbotapi.NewInlineKeyboardButtonData("60", "60"),
		tgbotapi.NewInlineKeyboardButtonData("80", "80"),
		tgbotapi.NewInlineKeyboardButtonData("100", "100"),
		tgbotapi.NewInlineKeyboardButtonData("250", "250"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("500", "500"),
	),
)

var VehiclesTypeInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("all", "vehicles_type_all"),
		tgbotapi.NewInlineKeyboardButtonData("cars and lorries", "vehicles_type_cars_and_lorries"),
		tgbotapi.NewInlineKeyboardButtonData("motorcycles", "vehicles_type_motorcycles"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("powersports", "vehicles_type_powersports"),
		tgbotapi.NewInlineKeyboardButtonData("motorhomes and campers", "vehicles_type_motorhomes_and_campers"),
		tgbotapi.NewInlineKeyboardButtonData("boats", "vehicles_type_boats"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("commercial and industrial", "vehicles_type_commercial_and_industrial"),
		tgbotapi.NewInlineKeyboardButtonData("trailers", "vehicles_type_trailers"),
		tgbotapi.NewInlineKeyboardButtonData("other", "vehicles_type_other"),
	),
)
var ExteriorColourInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("black", "exterior_black"),
		tgbotapi.NewInlineKeyboardButtonData("charcoal", "exterior_charcoal"),
		tgbotapi.NewInlineKeyboardButtonData("grey", "exterior_grey"),
		tgbotapi.NewInlineKeyboardButtonData("silver", "exterior_silver"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("white", "exterior_white"),
		tgbotapi.NewInlineKeyboardButtonData("off white", "exterior_off_white"),
		tgbotapi.NewInlineKeyboardButtonData("tan", "exterior_tan"),
		tgbotapi.NewInlineKeyboardButtonData("beige", "exterior_beige"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("yellow", "exterior_yellow"),
		tgbotapi.NewInlineKeyboardButtonData("gold", "exterior_gold"),
		tgbotapi.NewInlineKeyboardButtonData("brown", "exterior_brown"),
		tgbotapi.NewInlineKeyboardButtonData("orange", "exterior_orange"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("red", "exterior_red"),
		tgbotapi.NewInlineKeyboardButtonData("burgundy", "exterior_burgundy"),
		tgbotapi.NewInlineKeyboardButtonData("pink", "exterior_pink"),
		tgbotapi.NewInlineKeyboardButtonData("purple", "exterior_purple"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("blue", "exterior_blue"),
		tgbotapi.NewInlineKeyboardButtonData("turquoise", "exterior_turquoise"),
		tgbotapi.NewInlineKeyboardButtonData("green", "exterior_green"),
		tgbotapi.NewInlineKeyboardButtonData("other", "exterior_other"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("никакой цвет", "exterior_nothing"),
	),
)
var InteriorColourInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("black", "interior_black"),
		tgbotapi.NewInlineKeyboardButtonData("charcoal", "interior_charcoal"),
		tgbotapi.NewInlineKeyboardButtonData("grey", "interior_grey"),
		tgbotapi.NewInlineKeyboardButtonData("silver", "interior_silver"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("white", "interior_white"),
		tgbotapi.NewInlineKeyboardButtonData("off white", "interior_off_white"),
		tgbotapi.NewInlineKeyboardButtonData("tan", "interior_tan"),
		tgbotapi.NewInlineKeyboardButtonData("beige", "interior_beige"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("yellow", "interior_yellow"),
		tgbotapi.NewInlineKeyboardButtonData("gold", "interior_gold"),
		tgbotapi.NewInlineKeyboardButtonData("brown", "interior_brown"),
		tgbotapi.NewInlineKeyboardButtonData("orange", "interior_orange"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("red", "interior_red"),
		tgbotapi.NewInlineKeyboardButtonData("burgundy", "interior_burgundy"),
		tgbotapi.NewInlineKeyboardButtonData("pink", "interior_pink"),
		tgbotapi.NewInlineKeyboardButtonData("purple", "interior_purple"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("blue", "interior_blue"),
		tgbotapi.NewInlineKeyboardButtonData("turquoise", "interior_turquoise"),
		tgbotapi.NewInlineKeyboardButtonData("green", "interior_green"),
		tgbotapi.NewInlineKeyboardButtonData("other", "interior_other"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("никакой цвет", "interior_nothing"),
	),
)
var MakeInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("all", "make_all"),
		tgbotapi.NewInlineKeyboardButtonData("alfa romeo", "make_alfa_romeo"),
		tgbotapi.NewInlineKeyboardButtonData("aston martin", "make_aston_martin"),
		tgbotapi.NewInlineKeyboardButtonData("audi", "make_audi"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("BMW", "make_BMW"),
		tgbotapi.NewInlineKeyboardButtonData("bentley", "make_bentley"),
		tgbotapi.NewInlineKeyboardButtonData("buick", "make_buick"),
		tgbotapi.NewInlineKeyboardButtonData("cadillac", "make_cadillac"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("chevrolet", "make_chevrolet"),
		tgbotapi.NewInlineKeyboardButtonData("chrysler", "make_chrysler"),
		tgbotapi.NewInlineKeyboardButtonData("daewoo", "make_daewoo"),
		tgbotapi.NewInlineKeyboardButtonData("dodge", "make_dodge"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("ferrari", "make_ferrari"),
		tgbotapi.NewInlineKeyboardButtonData("FIAT", "make_FIAT"),
		tgbotapi.NewInlineKeyboardButtonData("ford", "make_ford"),
		tgbotapi.NewInlineKeyboardButtonData("honda", "make_honda"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("hummer", "make_hummer"),
		tgbotapi.NewInlineKeyboardButtonData("hyundai", "make_hyundai"),
		tgbotapi.NewInlineKeyboardButtonData("INFINITI", "make_INFINITI"),
		tgbotapi.NewInlineKeyboardButtonData("isuzu", "make_isuzu"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("jaguar", "make_jaguar"),
		tgbotapi.NewInlineKeyboardButtonData("->", "make_next"),
		tgbotapi.NewInlineKeyboardButtonData("jeep", "make_jeep"),
	),
)
var MakeInlineKeyboard1 = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("kia", "make_kia"),
		tgbotapi.NewInlineKeyboardButtonData("lamborghini", "make_lamborghini"),
		tgbotapi.NewInlineKeyboardButtonData("land rover", "make_land_rover"),
		tgbotapi.NewInlineKeyboardButtonData("lotus", "make_lotus"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("MINI", "make_MINI"),
		tgbotapi.NewInlineKeyboardButtonData("maserati", "make_maserati"),
		tgbotapi.NewInlineKeyboardButtonData("mazda", "make_mazda"),
		tgbotapi.NewInlineKeyboardButtonData("McLaren", "make_McLaren"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("mercedes-menz", "make_mercedes_menz"),
		tgbotapi.NewInlineKeyboardButtonData("mitsubishi", "make_mitsubishi"),
		tgbotapi.NewInlineKeyboardButtonData("nissan", "make_nissan"),
		tgbotapi.NewInlineKeyboardButtonData("pontiac", "make_pontiac"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("porsche", "make_porsche"),
		tgbotapi.NewInlineKeyboardButtonData("rolls-royce", "make_rolls_royce"),
		tgbotapi.NewInlineKeyboardButtonData("saab", "make_saab"),
		tgbotapi.NewInlineKeyboardButtonData("smart", "make_smart"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("subaru", "make_subaru"),
		tgbotapi.NewInlineKeyboardButtonData("suzuki", "make_suzuki"),
		tgbotapi.NewInlineKeyboardButtonData("tesla", "make_tesla"),
		tgbotapi.NewInlineKeyboardButtonData("toyota", "make_toyota"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("volkswagen", "make_volkswagen"),
		tgbotapi.NewInlineKeyboardButtonData("<-", "make_previous"),
		tgbotapi.NewInlineKeyboardButtonData("volvo", "make_volvo"),
	),
)

var TransmissonInlineKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("automatic", "automatic"),
		tgbotapi.NewInlineKeyboardButtonData("manual", "manual"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("all", "transmisson_all"),
	),
)

var BodyStyleAlfaRomeoInlineKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("convertible", "body_style_alfa_romeo_convertible"),
		tgbotapi.NewInlineKeyboardButtonData("coupe", "body_style_alfa_romeo_Coupe"),
		tgbotapi.NewInlineKeyboardButtonData("hatchback", "body_style_alfa_romeo_hatchback"),
		tgbotapi.NewInlineKeyboardButtonData("MPV/People carrier", "body_style_alfa_romeo_carrier"),
		tgbotapi.NewInlineKeyboardButtonData("saloon", "body_style_alfa_romeo_saloon"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("estate", "body_style_alfa_romeo_estate"),
		tgbotapi.NewInlineKeyboardButtonData("4x4", "body_style_alfa_romeo_4x4"),
		tgbotapi.NewInlineKeyboardButtonData("van", "body_style_alfa_romeo_van"),
		tgbotapi.NewInlineKeyboardButtonData("small car", "body_style_alfa_romeo_small_car"),
		tgbotapi.NewInlineKeyboardButtonData("other", "body_style_alfa_romeo_other"),
	),
)

var RoomsInlineKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("1+", "rooms_1"),
		tgbotapi.NewInlineKeyboardButtonData("2+", "rooms_2"),
		tgbotapi.NewInlineKeyboardButtonData("3+", "rooms_3"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("4+", "rooms_4"),
		tgbotapi.NewInlineKeyboardButtonData("5+", "rooms_5"),
		tgbotapi.NewInlineKeyboardButtonData("6+", "rooms_6"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("all", "rooms_all"),
	),
)
var TypePropertyInlineKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("flat/apartment", "type_property_flat_apartment"),
		tgbotapi.NewInlineKeyboardButtonData("house", "type_property_house"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("room only", "type_property_room_only"),
		tgbotapi.NewInlineKeyboardButtonData("townhouse", "type_property_townhouse"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("all", "type_property_all"),
	),
)
var ConditionInlineKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("new", "condition_new"),
		tgbotapi.NewInlineKeyboardButtonData("used - like new", "condition_used"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("used - good", "condition_good"),
		tgbotapi.NewInlineKeyboardButtonData("used - fair", "condition_fair"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("all", "condition_all"),
	),
)

var BrandInlineKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("AMD", "electronics_AMD"),
		tgbotapi.NewInlineKeyboardButtonData("ASUS", "electronics_ASUS"),
		tgbotapi.NewInlineKeyboardButtonData("Apple", "electronics_Apple"),
		tgbotapi.NewInlineKeyboardButtonData("beats by Dr.Dre", "electronics_Dre"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("bose", "electronics_bose"),
		tgbotapi.NewInlineKeyboardButtonData("canon", "electronics_canon"),
		tgbotapi.NewInlineKeyboardButtonData("dell", "electronics_dell"),
		tgbotapi.NewInlineKeyboardButtonData("google", "electronics_google"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("panasonic", "electronics_panasonic"),
		tgbotapi.NewInlineKeyboardButtonData("intel", "electronics_intel"),
		tgbotapi.NewInlineKeyboardButtonData("nintedo", "electronics_nintedo"),
		tgbotapi.NewInlineKeyboardButtonData("OtterBox", "electronics_otterBox"),
	),
	tgbotapi.NewInlineKeyboardRow(

		tgbotapi.NewInlineKeyboardButtonData("all", "electronics_all"),
		tgbotapi.NewInlineKeyboardButtonData("VIZIO", "electronics_VIZIO"),
	),
)
var HomeImprovementInlineKB = tgbotapi.NewInlineKeyboardMarkup(
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("craftsman", "home_Improvement_craftsman"),
		tgbotapi.NewInlineKeyboardButtonData("DEWALT", "home_Improvement_DEWALT"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("milwaukee", "home_Improvement_milwaukee"),
		tgbotapi.NewInlineKeyboardButtonData("snap-on", "home_Improvement_snap_on"),
	),
	tgbotapi.NewInlineKeyboardRow(
		tgbotapi.NewInlineKeyboardButtonData("all", "home_Improvement_all"),
	),
)

var CityInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup()
var FilterInlineKeyboard = tgbotapi.NewInlineKeyboardMarkup()

var StartKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Мои фильтры"),
		tgbotapi.NewKeyboardButton("Добавить фильтр"),
	),
)
var CancelKeyboard = tgbotapi.NewReplyKeyboard(
	tgbotapi.NewKeyboardButtonRow(
		tgbotapi.NewKeyboardButton("Отмена"),
	),
)
