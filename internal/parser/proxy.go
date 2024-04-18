package parser

import (
	"context"
	"log"

	"github.com/chromedp/cdproto/cdp"
	"github.com/chromedp/cdproto/fetch"
	"github.com/chromedp/chromedp"
)

func (p *StrParser) SetProxy(ctx context.Context, data Data) {
	chromedp.ListenTarget(ctx, func(ev interface{}) {
		go func() {
			switch ev := ev.(type) {
			case *fetch.EventAuthRequired:
				c := chromedp.FromContext(ctx)
				execCtx := cdp.WithExecutor(ctx, c.Target)

				resp := &fetch.AuthChallengeResponse{
					Response: fetch.AuthChallengeResponseResponseProvideCredentials,
					Username: data.LoginPX,
					Password: data.PassPX,
				}

				err := fetch.ContinueWithAuth(ev.RequestID, resp).Do(execCtx)
				if err != nil {
					log.Fatalf("ContinueWithAuth error: %v", err)
				}

			case *fetch.EventRequestPaused:
				c := chromedp.FromContext(ctx)
				execCtx := cdp.WithExecutor(ctx, c.Target)

				err := fetch.ContinueRequest(ev.RequestID).Do(execCtx)
				if err != nil {
					log.Fatalf("ContinueReques error: %v", err)
				}
			}
		}()
	})

	if err := chromedp.Run(ctx,
		fetch.Enable().WithHandleAuthRequests(true),
	); err != nil {
		log.Fatalf("WithHandleAuthRequests error: %v", err)
	}

	log.Print("end set proxy")
}
