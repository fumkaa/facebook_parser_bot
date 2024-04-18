package parser

import (
	"context"
	database "facebook_marketplace_bot/internal/database/migration"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var (
	idAd     int
	isCancel bool
	ChId     = make(chan int)
)

func (p *StrParser) Monitoring(ctx context.Context, url string, id int) error {
	isCancel = false
	for {
		log.Print("MONITORING......")

		_, err := p.db.MonitoringByIDFilter(ctx, id)
		if err == database.ErrNoRows {
			isCancel = true
			log.Printf("END MONITORING  BY ID %d", id)
		}
		if err != nil && err != database.ErrNoRows {
			log.Printf("MonitoringByIDFilter eror: %v", err)
			return nil
		}

		if isCancel {
			log.Printf("END MONITORING BY ID %d", id)
			break
		}
		var (
			nodes   []*cdp.Node
			elUrlAd []string
		)
		err = chromedp.Run(ctx,
			chromedp.Navigate(url),
		)
		if err != nil {
			return fmt.Errorf("[monitoring]run error: %w", err)
		}
		log.Print("Navigate")
		err = chromedp.Run(ctx,
			chromedp.Nodes(`a[class="x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xt0psk2 xe8uvvx xdj266r x11i5rnm xat24cr x1mh8g0r xexx8yu x4uap5 x18d9i69 xkhd6sd x16tdsg8 x1hl2dhg xggy1nq x1a2a7pz x1heor9g x1lku1pv"]`,
				&nodes, chromedp.ByQuery),
		)
		if err != nil {
			return fmt.Errorf("[monitoring]Nodes error: %w", err)
		}
		log.Print("Nodes")
		for _, node := range nodes {
			urlAd := node.AttributeValue("href")
			elUrlAd = strings.Split(urlAd, "/")
		}
		log.Print("for")
		curId, err := strconv.Atoi(elUrlAd[3])
		if err != nil {
			return fmt.Errorf("convert elUrlAd error: %w", err)
		}
		if idAd == 0 {
			idAd = curId
			time.Sleep(30 * time.Second)
			continue
		}
		if idAd == curId {
			log.Print("id ad not change")
			log.Printf("cur id ad: %d\nold id ad: %d", curId, idAd)
			time.Sleep(30 * time.Second)
			continue
		} else if idAd != curId {
			ChId <- curId
			idAd = curId
			continue
		}
	}
	return nil
}
