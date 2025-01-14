package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/chromedp/cdproto/network"
	"github.com/chromedp/chromedp"
)

func main() {
	// Create a new context
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()

	// Run chromedp tasks
	var cookies []*network.Cookie
	err := chromedp.Run(ctx,
		network.Enable(), // Enable network events
		chromedp.Navigate(`https://www.ozon.ru/category/smartfony-15502/?category_was_predicted=true&deny_category_prediction=true&from_global=true&text=Xiaomi+%D0%A1%D0%BC%D0%B0%D1%80%D1%82%D1%84%D0%BE%D0%BD+Redmi+14C`), // Navigate to the site
		chromedp.ActionFunc(func(ctx context.Context) error {
			var err error
			// Get all cookies
			cookies, err = network.GetCookies().Do(ctx)
			return err
		}),
	)
	if err != nil {
		log.Fatalf("Failed to get cookies: %v", err)
	}

	// Concatenate cookies into a single line
	var cookieStrings []string
	for _, cookie := range cookies {
		cookieStrings = append(cookieStrings, fmt.Sprintf("%s=%s", cookie.Name, cookie.Value))
	}

	// Join the cookie strings with "; "
	cookieLine := strings.Join(cookieStrings, "; ")
	fmt.Println("Cookies:", cookieLine)
}
