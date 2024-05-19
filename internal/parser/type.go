package parser

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
)

type Parser interface {
	SelectCity(ctx context.Context, city string) ([]string, error)
	SetRadius(ctx context.Context, curRadius string) error
	Settings(ctx context.Context, datas Datas) error
	ClickSelectCity(ctx context.Context, city string) error
	GetData() ([]Datas, error)
	GetDataFile(name string) (Datas, error)
}

type StrParser struct {
	Data Data
	db   database.Database
	// customSettings  customSettings
}
type Datas struct {
	FileName string
	Datas    Data
}
type Data struct {
	LoginFB  string
	PassFB   string
	IpPortPX string
	LoginPX  string
	PassPX   string
	Cookies  string
}

// type customSettings struct {
// 	radius string
// 	url    string
// }

func NewParser(db database.Database) *StrParser {
	return &StrParser{
		db: db,
	}
}

// type Category struct {
// 	vehicles          bool
// 	propertyrentals   bool
// 	free              bool
// 	toys              bool
// 	instruments       bool
// 	home_improvements bool
// 	classifieds       bool
// 	apparel           bool
// 	propertyforsale   bool
// 	entertainment     bool
// 	family            bool
// 	sports            bool
// 	home              bool
// 	pets              bool
// 	office_supplies   bool
// 	garden            bool
// 	hobbies           bool
// 	electronics       bool
// 	groups            bool
// 	all_listings      bool
// }
