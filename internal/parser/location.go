package parser

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/chromedp"
)

var isCorrect bool

func (p *StrParser) SetRadius(ctx context.Context, curRadius string) error {
	var (
		radius []*cdp.Node
		rad    string
		curNum string
	)
	var res []byte
	log.Print("run click nodes")
	if err := chromedp.Run(ctx,
		chromedp.Click(`div[class="xjyslct xjbqb8w x13fuv20 xu3j5b3 x1q0q8m5 x26u7qi x972fbf xcfux6l x1qhh985 xm0m39n x9f619 xzsf02u x78zum5 x1jchvi3 x1fcty0u x132q4wb xdj266r x11i5rnm xat24cr x1mh8g0r x1a2a7pz x9desvi x1pi30zi x1a8lsjc x1swvt13 x1n2onr6 x16tdsg8 xh8yej3 x1ja2u2z"]`,
			chromedp.ByQuery),
		chromedp.Nodes(`div[class="x1i10hfl xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x972fbf xcfux6l x1qhh985 xm0m39n xe8uvvx x1hl2dhg xggy1nq x1o1ewxj x3x9cwd x1e5q0jg x13rtm0m x87ps6o x1lku1pv x1a2a7pz x6s0dn4 xjyslct x9f619 x1ypdohk x78zum5 x1q0g3np x2lah0s xnqzcj9 x1gh759c xdj266r xat24cr x1344otq x1de53dj xz9dl7a xsag5q8 x1n2onr6 x16tdsg8 x1ja2u2z"]`,
			&radius, chromedp.ByQueryAll),
		chromedp.FullScreenshot(&res, 90),
	); err != nil {
		return fmt.Errorf("click, nodes error: %w", err)
	}
	err := os.WriteFile("test.png", res, 00644)
	if err != nil {
		log.Fatal("Error:", err)
	}
	log.Print("for text")
	for _, node := range radius {
		if err := chromedp.Run(ctx,
			chromedp.Text(`div[class="x6s0dn4 x78zum5 x1q0g3np x1iyjqo2 x1qughib xeuugli"]`,
				&rad, chromedp.ByQuery, chromedp.FromNode(node)),
		); err != nil {
			return fmt.Errorf("text error: %w", err)
		}
		log.Printf("don't fetch rad = %s", rad)
		rads := strings.Split(rad, "")
		for _, num := range rads {
			_, err := strconv.Atoi(num)
			if err != nil {
				break
			}
			curNum += num
		}
		log.Printf("curNum: %s", curNum)
		if curRadius == curNum {
			log.Print("mouse click node")
			if err := chromedp.Run(ctx,
				chromedp.MouseClickNode(node),
			); err != nil {
				return fmt.Errorf("mouse click error: %w", err)
			}
			break
		}
		curNum = ""
	}
	log.Printf("set radius %s", curNum)
	if isCorrect {
		if err := chromedp.Run(ctx,
			chromedp.Click(`div[class="x1n2onr6 x1ja2u2z x78zum5 x2lah0s xl56j7k x6s0dn4 xozqiw3 x1q0g3np xi112ho x17zwfj4 x585lrc x1403ito x972fbf xcfux6l x1qhh985 xm0m39n x9f619 xn6708d x1ye3gou xtvsq51 x1r1pt67"]`,
				chromedp.ByQuery),
		); err != nil {
			return fmt.Errorf("click error: %w", err)
		}
		log.Print("click aplay")
	}
	return nil
}

func (p *StrParser) SelectCity(ctx context.Context, city string) ([]string, error) {
	log.Print("start set city")
	var (
		cityNodes  []*cdp.Node
		selectCity string
		cities     []string
	)
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.facebook.com/marketplace/"),
	); err != nil {
		return nil, fmt.Errorf("navigate error: %w", err)
	}
	var res interface{}
	if err := chromedp.Run(ctx,
		chromedp.Click(`div[class="x1i10hfl x1qjc9v5 xjbqb8w xjqpnuy xa49m3k xqeqjp1 x2hbi6w x13fuv20 xu3j5b3 x1q0q8m5 x26u7qi x972fbf xcfux6l x1qhh985 xm0m39n x9f619 x1ypdohk xdl72j9 x2lah0s xe8uvvx x11i5rnm xat24cr x1mh8g0r x2lwn1j xeuugli xexx8yu x4uap5 x18d9i69 xkhd6sd x1n2onr6 x16tdsg8 x1hl2dhg xggy1nq x1ja2u2z x1t137rt x1o1ewxj x3x9cwd x1e5q0jg x13rtm0m x1q0g3np x87ps6o x1lku1pv x78zum5 x1a2a7pz x1xmf6yo"]`,
			chromedp.ByQuery),
		chromedp.SendKeys(`input[class="x1i10hfl xggy1nq x1s07b3s x1kdt53j x1a2a7pz xjbqb8w x1ejq31n xd10rxx x1sy0etr x17r0tee x9f619 xzsf02u x1uxerd5 x1fcty0u x132q4wb x1a8lsjc x1pi30zi x1swvt13 x9desvi xh8yej3 x15h3p50 x10emqs4"]`,
			city, chromedp.ByQuery),
	); err != nil {
		return nil, fmt.Errorf("sendkeys, nodes error: %w", err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(`document.querySelector("li.xh8yej3").textContent`, &res),
	); err != nil {
		log.Printf("evaluate error: %v", err)
		return cities, nil
	}
	log.Printf("res: %v", res)
	if res == nil {
		return cities, nil
	}

	if err := chromedp.Run(ctx,
		chromedp.Nodes(`li[class="xh8yej3"]`, &cityNodes, chromedp.ByQueryAll),
	); err != nil {
		return nil, fmt.Errorf("nodes error: %w", err)
	}
	for _, node := range cityNodes {
		if err := chromedp.Run(ctx,
			chromedp.Text(`div[class="x6s0dn4 x1ypdohk x78zum5 x6ikm8r x10wlt62 x1n2onr6 x8du52y x1lq5wgf xgqcy7u x30kzoy x9jhf4c xdj266r xat24cr x1y1aw1k x1sxyh0 xwib8y2 xurb0ha"]`,
				&selectCity, chromedp.ByQuery, chromedp.FromNode(node),
			),
		); err != nil {
			return nil, fmt.Errorf("text error: %w", err)
		}
		cit := strings.Split(selectCity, "\n")
		cities = append(cities, cit[0])
	}
	log.Printf("cities: %v", cities)
	log.Print("select city end")
	return cities, nil

}

func (p *StrParser) ClickSelectCity(ctx context.Context, city string) error {
	var (
		cityNodes  []*cdp.Node
		selectCity string
	)

	if err := chromedp.Run(ctx,
		chromedp.Nodes(`li[class="xh8yej3"]`, &cityNodes, chromedp.ByQueryAll),
	); err != nil {
		return fmt.Errorf("nodes error: %w", err)
	}

	for _, node := range cityNodes {
		if err := chromedp.Run(ctx,
			chromedp.Text(`div[class="x6s0dn4 x1ypdohk x78zum5 x6ikm8r x10wlt62 x1n2onr6 x8du52y x1lq5wgf xgqcy7u x30kzoy x9jhf4c xdj266r xat24cr x1y1aw1k x1sxyh0 xwib8y2 xurb0ha"]`,
				&selectCity, chromedp.ByQuery, chromedp.FromNode(node),
			),
		); err != nil {
			return fmt.Errorf("text error: %w", err)
		}
		cit := strings.Split(selectCity, "\n")
		log.Printf("selectCity: %s", cit[0])
		log.Printf("city: %s", city)
		if city == cit[0] {
			isCorrect = true
			log.Print("selectCity == cit1")
			if err := chromedp.Run(ctx,
				chromedp.MouseClickNode(node),
			); err != nil {
				return fmt.Errorf("mouse click child error: %w", err)
			}

			break
		}
	}
	return nil
}
