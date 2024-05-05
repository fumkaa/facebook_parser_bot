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
	log.Print("setCookies Navigate")
	if err := chromedp.Run(ctx,
		chromedp.Sleep(5*time.Second),
		chromedp.Evaluate("document.querySelector(`h4`).textContent;", &res),
	); err != nil {
		log.Printf("evaluate error: %v", err)
		log.Printf("%s invalid cookies", data.FileName)
		if err := p.LoginFBaccountToLogPass(ctx, data.Datas); err != nil {
			return fmt.Errorf("login to username and password error: %w", err)
		}
		if err := chromedp.Run(ctx,
			chromedp.Sleep(5*time.Second),
			chromedp.Evaluate(fmt.Sprintf("document.querySelector(`h4[id=%q]`).textContent;", ":R5alanalaajml5bb9l5qq9papd5aq:"), &res),
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
	log.Print("start LoginFBaccountToLogPass")
	if err := chromedp.Run(ctx,
		network.ClearBrowserCookies(),
	); err != nil {
		return fmt.Errorf("navigate error: %w", err)
	}
	var res2 interface{}
	var res1 interface{}
	log.Print(" ClearBrowserCookies")
	if err := chromedp.Run(ctx,
		chromedp.Navigate("https://www.facebook.com"),
	); err != nil {
		return fmt.Errorf("navigate error: %v", err)
	}

	log.Print(" Navigate")
	if err := chromedp.Run(ctx,
		chromedp.Evaluate(fmt.Sprintf("document.querySelector(`div[class=%q]`).textContent;", "x9f619 x1n2onr6 x1ja2u2z x78zum5 xdt5ytf x193iq5w xeuugli x1iyjqo2 xs83m0k x150jy0e x1e558r4 xjkvuk6 x1iorvi4 xdl72j9"), &res1),
		// chromedp.Evaluate(fmt.Sprintf("document.querySelector(`button[class=%q]`).textContent;", "_42ft _4jy0 _al65 _4jy3 _4jy1 selected _51sy"), &res1),
	); err != nil {
		log.Print("Evaluate err")
		if err := chromedp.Run(ctx,
			chromedp.Evaluate(fmt.Sprintf("document.querySelector(`button[class=%q]`).textContent;", "_42ft _4jy0 _al65 _4jy3 _4jy1 selected _51sy"), &res2),
		); err != nil {
			if err := chromedp.Run(ctx,
				chromedp.WaitVisible(`input[class="inputtext _55r1 inputtext _1kbt inputtext _1kbt"]`, chromedp.ByQuery),
			); err != nil {
				return fmt.Errorf("sendKeys WaitVisible login error: %v", err)
			}
			log.Print(" WaitVisible")
			if err := chromedp.Run(ctx,
				chromedp.Clear(`input[class="inputtext _55r1 inputtext _1kbt inputtext _1kbt"]`, chromedp.ByQuery),
			); err != nil {
				return fmt.Errorf("sendKeys Clear login error: %v", err)
			}
			log.Print(" Clear")
			if err := chromedp.Run(ctx,
				chromedp.SendKeys(`input[class="inputtext _55r1 inputtext _1kbt inputtext _1kbt"]`, data.LoginFB, chromedp.ByQuery),
			); err != nil {
				return fmt.Errorf("sendKeys login error: %v", err)
			}
			log.Print(" SendKeys")
			if err := chromedp.Run(ctx,
				chromedp.Clear(`input[class="inputtext _55r1 inputtext _9npi inputtext _9npi"]`, chromedp.ByQuery),
				chromedp.SendKeys(`input[class="inputtext _55r1 inputtext _9npi inputtext _9npi"]`, data.PassFB, chromedp.ByQuery),
			); err != nil {
				return fmt.Errorf("sendKeys password error: %v", err)
			}

			if err := chromedp.Run(ctx,
				chromedp.Submit(`button[class="_42ft _4jy0 _52e0 _4jy6 _4jy1 selected _51sy"]`, chromedp.ByQuery),
				chromedp.Sleep(5*time.Second),
			); err != nil {
				return fmt.Errorf("click log in error: %v", err)
			}
			return nil
		}
		log.Printf("res: %v", res2)
		if err := chromedp.Run(ctx,
			chromedp.Click(`button[class="_42ft _4jy0 _al65 _4jy3 _4jy1 selected _51sy"]`, chromedp.ByQuery),
		); err != nil {
			return fmt.Errorf("click error: %w", err)
		}
		if err := chromedp.Run(ctx,
			chromedp.WaitVisible(`input[class="inputtext _55r1 _6luy"]`, chromedp.ByQuery),
		); err != nil {
			return fmt.Errorf("sendKeys WaitVisible login error: %v", err)
		}
		log.Print(" WaitVisible")
		if err := chromedp.Run(ctx,
			chromedp.Clear(`input[class="inputtext _55r1 _6luy"]`, chromedp.ByQuery),
		); err != nil {
			return fmt.Errorf("sendKeys Clear login error: %v", err)
		}
		log.Print(" Clear")
		if err := chromedp.Run(ctx,
			chromedp.SendKeys(`input[class="inputtext _55r1 _6luy"]`, data.LoginFB, chromedp.ByQuery),
		); err != nil {
			return fmt.Errorf("sendKeys login error: %v", err)
		}
		log.Print(" SendKeys")
		if err := chromedp.Run(ctx,
			chromedp.Clear(`input[class="inputtext _55r1 _6luy"]`, chromedp.ByQuery),
			chromedp.SendKeys(`input[class="inputtext _55r1 _6luy _9npi"]`, data.PassFB, chromedp.ByQuery),
		); err != nil {
			return fmt.Errorf("sendKeys password error: %v", err)
		}

		if err := chromedp.Run(ctx,
			chromedp.Submit(`button[class="_42ft _4jy0 _6lth _4jy6 _4jy1 selected _51sy"]`, chromedp.ByQuery),
			chromedp.Sleep(5*time.Second),
		); err != nil {
			return fmt.Errorf("click log in error: %v", err)
		}
		log.Printf("login: %v", data.LoginFB)
		log.Printf("pass: %v", data.PassFB)
		return nil
	}
	log.Printf("res1: %v", res1)
	if err := chromedp.Run(ctx,
		chromedp.Click(`div[class="x9f619 x1n2onr6 x1ja2u2z x78zum5 xdt5ytf x193iq5w xeuugli x1iyjqo2 xs83m0k x150jy0e x1e558r4 xjkvuk6 x1iorvi4 xdl72j9"]`, chromedp.ByQuery),
	); err != nil {
		return fmt.Errorf("sendKeys login error: %v", err)
	}
	// input[class="inputtext _55r1 inputtext _1kbt inputtext _1kbt"]
	if err := chromedp.Run(ctx,
		chromedp.WaitVisible(`input[class="inputtext _55r1 _6luy"]`, chromedp.ByQuery),
	); err != nil {
		return fmt.Errorf("sendKeys WaitVisible login error: %v", err)
	}
	log.Print(" WaitVisible")
	if err := chromedp.Run(ctx,
		chromedp.Clear(`input[class="inputtext _55r1 _6luy"]`, chromedp.ByQuery),
	); err != nil {
		return fmt.Errorf("sendKeys Clear login error: %v", err)
	}
	log.Print(" Clear")
	if err := chromedp.Run(ctx,
		chromedp.SendKeys(`input[class="inputtext _55r1 _6luy"]`, data.LoginFB, chromedp.ByQuery),
	); err != nil {
		return fmt.Errorf("sendKeys login error: %v", err)
	}
	log.Print(" SendKeys")
	if err := chromedp.Run(ctx,
		chromedp.Clear(`input[class="inputtext _55r1 _6luy"]`, chromedp.ByQuery),
		chromedp.SendKeys(`input[class="inputtext _55r1 _6luy _9npi"]`, data.PassFB, chromedp.ByQuery),
	); err != nil {
		return fmt.Errorf("sendKeys password error: %v", err)
	}
	if err := chromedp.Run(ctx,
		chromedp.Submit(`button[class="_42ft _4jy0 _6lth _4jy6 _4jy1 selected _51sy"]`, chromedp.ByQuery),
		chromedp.Sleep(5*time.Second),
	); err != nil {
		return fmt.Errorf("click log in error: %v", err)
	}
	log.Printf("login: %v", data.LoginFB)
	log.Printf("pass: %v", data.PassFB)
	return nil
}
