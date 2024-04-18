package parser

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/cdproto/storage"
	"github.com/chromedp/chromedp"
)

type Cookie struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

const (
	Free_account = "data/free_account/"
	Work_account = "data/work_account/"
)

var (
	ErrEmptyData             = errors.New("err empty cookie dir")
	ErrAccountBanned         = errors.New("err account was banned")
	ErrProxyConnectionFailed = errors.New("page load error net::ERR_PROXY_CONNECTION_FAILED")
)

func setCookies(cookies []Cookie) chromedp.Tasks {
	return chromedp.Tasks{
		chromedp.ActionFunc(func(ctx context.Context) error {
			// create cookie expiration
			expr := cdp.TimeSinceEpoch(time.Now().Add(180 * 24 * time.Hour))
			// add cookies to chrome
			for _, cookie := range cookies {
				err := network.SetCookie(cookie.Name, cookie.Value).
					WithExpires(&expr).
					WithDomain("www.facebook.com/marketplace").
					WithHTTPOnly(true).
					Do(ctx)
				if err != nil {
					return fmt.Errorf("set cookie error: %w", err)
				}
			}

			return nil
		}),
		// read the returned values
		// read network values
		chromedp.ActionFunc(func(ctx context.Context) error {
			cookies, err := storage.GetCookies().Do(ctx)
			if err != nil {
				return fmt.Errorf("get cookie error: %w", err)
			}

			for i, cookie := range cookies {
				log.Printf("chrome cookie %d: %+v", i, cookie)
			}
			return nil
		}),
	}
}

func (p *StrParser) checkAndSetCookie(ctx context.Context, data Datas) error {

	log.Print(data.FileName)
	log.Print("set proxy")
	p.SetProxy(ctx, data.Datas)
	var cookies []Cookie
	if err := json.Unmarshal([]byte(data.Datas.Cookies), &cookies); err != nil {
		return fmt.Errorf("unmarshal cookie file error: %w", err)
	}
	var res interface{}
	err := chromedp.Run(ctx,
		setCookies(cookies),
		chromedp.Navigate("https://www.facebook.com"),
	)
	if err != nil {
		log.Printf("run error: %v", err)
		return ErrProxyConnectionFailed
	}

	if err := chromedp.Run(ctx,
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate(fmt.Sprintf("document.querySelector(`h4[id=%q]`).textContent;", ":R1alalqlaiktl9aqqd9emhpapd5aq:"), &res),
	); err != nil {
		log.Printf("evaluate error: %v", err)
		log.Printf("%s invalid cookies", data.FileName)
		if err := p.LoginFBaccountToLogPass(ctx, data.Datas); err != nil {
			return fmt.Errorf("login to username and password error: %w", err)
		}
		if err := chromedp.Run(ctx,
			chromedp.Sleep(5*time.Second),
			chromedp.Evaluate(fmt.Sprintf("document.querySelector(`h4[id=%q]`).textContent;", ":R1alalqlaiktl9aqqd9emhpapd5aq:"), &res),
		); err != nil {
			log.Printf("evaluate error: %v", err)
			log.Printf("%s account banned", data.FileName)
			if err := os.Remove(Free_account + data.FileName); err != nil {
				return fmt.Errorf("remove error: %w", err)
			}
			return ErrAccountBanned
		}
		log.Printf("%s valid account", data.FileName)
		return nil
	}

	log.Printf("res: %v", res)

	log.Printf("%s valid cookies", data.FileName)
	log.Print("check cookie end")
	return nil

}

func (p *StrParser) LoginFBaccountToLogPass(ctx context.Context, data Data) error {

	if err := chromedp.Run(ctx,
		network.ClearBrowserCookies(),
		chromedp.Navigate("https://www.facebook.com"),
	); err != nil {
		return fmt.Errorf("navigate error: %w", err)
	}
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible(`input[class="inputtext _55r1 inputtext _1kbt inputtext _1kbt"]`, chromedp.ByQuery),
		chromedp.Clear(`input[class="inputtext _55r1 inputtext _1kbt inputtext _1kbt"]`, chromedp.ByQuery),
		chromedp.SendKeys(`input[class="inputtext _55r1 inputtext _1kbt inputtext _1kbt"]`, data.LoginFB, chromedp.ByQuery),
	); err != nil {
		return fmt.Errorf("sendKeys login error: %w", err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Clear(`input[class="inputtext _55r1 inputtext _9npi inputtext _9npi"]`, chromedp.ByQuery),
		chromedp.SendKeys(`input[class="inputtext _55r1 inputtext _9npi inputtext _9npi"]`, data.PassFB, chromedp.ByQuery),
	); err != nil {
		return fmt.Errorf("sendKeys password error: %w", err)
	}

	if err := chromedp.Run(ctx,
		chromedp.Submit(`button[class="_42ft _4jy0 _52e0 _4jy6 _4jy1 selected _51sy"]`, chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
	); err != nil {
		return fmt.Errorf("click log in error: %w", err)
	}
	return nil
}
