package main

import (
	"fmt"

	"github.com/VolcharaWIN/web-scraper/internal/refresh-cookie"
	"github.com/gocolly/colly"
)

func main() {
	// Create a new collector
	c := colly.NewCollector()
	c.UserAgent = "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/126.0.0.0 Safari/537.36"
	/*err := c.SetProxy("http://200.174.198.86:8888")
	if err != nil {
		fmt.Println("Ошибка при подключении к прокси: ", err)
	}*/
	// On every HTML element which matches the selector
	cookie := refresh.Refresh()
	fmt.Println("cookie: ", cookie)
	fmt.Printf("\n\n\n\n\n\n\n")
	targetCookie := "__Secure-ETC=40180dd7d5c12a19d63576f4098de0ec; abt_data=7.1yxcWSoDlyGIM2vw-zsrLmxNRisLCJNKHuKC88mw5N0MNvQSNnoKiEsk3RoywwsAaGRDUSOXPEZR9BuByFeJtLrUuVXi-cngoXsYLbG1mZpQ86PnGRoa4OKljRvVD9uPQNtrmvk5KWIgZKTEuOol01iUjCXvI7T6ip-Yh6hYXPBARpFkPdGJkiIzneAN2BSJvZl7y1TeuWObrmazvrUz69D6uMSllZ5qLuNOgSMnqf8TcyXljLD7dW893GAATOmphREGPCpf85ASfJ9QUB61i1igA0bC--1f9pK-wFFfGqU7o7kTjjAspLztJXeFH65HPxT2RkPPZuySdRcrNTou3-lDDvGbTAADaUpwcTdw9tJAsWZV1aqx5xQgKnXJh8p7kA5-QJPIRo4aP3FwLa-CMacOyxzzzE-SuWkR8RMthYcZF7cwsOjoOmc-vnqDvh76wSEhWXlfDlpXFxngHg; __Secure-access-token=6.0.fIP1nkLKRRWyJzaKcmPbFg.75.ASSCi_SFpMXf7Tj1nuZkPDkpEAMEanhkah_45eSwpMGLoV4hrnQ1zP5U175TfFG4OQ..20250114182310.pbY3r44SgryyW8x3GEZCL3qbxeKegmlDDJ7RbJOxukk.18a7968413532296; __Secure-user-id=0; __Secure-refresh-token=6.0.fIP1nkLKRRWyJzaKcmPbFg.75.ASSCi_SFpMXf7Tj1nuZkPDkpEAMEanhkah_45eSwpMGLoV4hrnQ1zP5U175TfFG4OQ..20250114182310.RSnXHagLpZtgDnufGYy_XODlv96e_E02eCcXGx3-qMw.196545d7d6c60d259; __Secure-ab-group=75; TS0149423d=0187c00a18aa27cda8fd193581028c4cedbb0e0af5b3a8322efcc8122ec43399aa03ffcc62cdc5621e0c96f2c27b43fccdb49063b0; TS0121feed=0187c00a18894928e388360410f5d26a67b165db1ce50ba3f43d56601d5ec76c2bbdf7b1c3ea10ca24f1c80c57f54a9f6568998629; xcid=e5b713df499d30b7bf2b39cfd3c041c9; __Secure-ext_xcid=e5b713df499d30b7bf2b39cfd3c041c9; rfuid=LTE5NTAyNjU0NzAsMzUuNzQ5OTcyMDkzODUwMzc0LC0xMzUwMDY5OTUyLFdpbmRvd3MgTlQgMTAuMDsgV2luNjQ7IHg2NCwxMDcxMDMwMTQxLFczc2libUZ0WlNJNklsQkVSaUJXYVdWM1pYSWlMQ0prWlhOamNtbHdkR2x2YmlJNklsQnZjblJoWW14bElFUnZZM1Z0Wlc1MElFWnZjbTFoZENJc0ltMXBiV1ZVZVhCbGN5STZXM3NpZEhsd1pTSTZJbUZ3Y0d4cFkyRjBhVzl1TDNCa1ppSXNJbk4xWm1acGVHVnpJam9pY0dSbUluMHNleUowZVhCbElqb2lkR1Y0ZEM5d1pHWWlMQ0p6ZFdabWFYaGxjeUk2SW5Ca1ppSjlYWDBzZXlKdVlXMWxJam9pUTJoeWIyMWxJRkJFUmlCV2FXVjNaWElpTENKa1pYTmpjbWx3ZEdsdmJpSTZJbEJ2Y25SaFlteGxJRVJ2WTNWdFpXNTBJRVp2Y20xaGRDSXNJbTFwYldWVWVYQmxjeUk2VzNzaWRIbHdaU0k2SW1Gd2NHeHBZMkYwYVc5dUwzQmtaaUlzSW5OMVptWnBlR1Z6SWpvaWNHUm1JbjBzZXlKMGVYQmxJam9pZEdWNGRDOXdaR1lpTENKemRXWm1hWGhsY3lJNkluQmtaaUo5WFgwc2V5SnVZVzFsSWpvaVEyaHliMjFwZFcwZ1VFUkdJRlpwWlhkbGNpSXNJbVJsYzJOeWFYQjBhVzl1SWpvaVVHOXlkR0ZpYkdVZ1JHOWpkVzFsYm5RZ1JtOXliV0YwSWl3aWJXbHRaVlI1Y0dWeklqcGJleUowZVhCbElqb2lZWEJ3YkdsallYUnBiMjR2Y0dSbUlpd2ljM1ZtWm1sNFpYTWlPaUp3WkdZaWZTeDdJblI1Y0dVaU9pSjBaWGgwTDNCa1ppSXNJbk4xWm1acGVHVnpJam9pY0dSbUluMWRmU3g3SW01aGJXVWlPaUpOYVdOeWIzTnZablFnUldSblpTQlFSRVlnVm1sbGQyVnlJaXdpWkdWelkzSnBjSFJwYjI0aU9pSlFiM0owWVdKc1pTQkViMk4xYldWdWRDQkdiM0p0WVhRaUxDSnRhVzFsVkhsd1pYTWlPbHQ3SW5SNWNHVWlPaUpoY0hCc2FXTmhkR2x2Ymk5d1pHWWlMQ0p6ZFdabWFYaGxjeUk2SW5Ca1ppSjlMSHNpZEhsd1pTSTZJblJsZUhRdmNHUm1JaXdpYzNWbVptbDRaWE1pT2lKd1pHWWlmVjE5TEhzaWJtRnRaU0k2SWxkbFlrdHBkQ0JpZFdsc2RDMXBiaUJRUkVZaUxDSmtaWE5qY21sd2RHbHZiaUk2SWxCdmNuUmhZbXhsSUVSdlkzVnRaVzUwSUVadmNtMWhkQ0lzSW0xcGJXVlVlWEJsY3lJNlczc2lkSGx3WlNJNkltRndjR3hwWTJGMGFXOXVMM0JrWmlJc0luTjFabVpwZUdWeklqb2ljR1JtSW4wc2V5SjBlWEJsSWpvaWRHVjRkQzl3WkdZaUxDSnpkV1ptYVhobGN5STZJbkJrWmlKOVhYMWQsV3lKeWRTMVNWU0lzSW5KMUxWSlZJaXdpY25VaUxDSmxiaTFWVXlJc0ltVnVJbDA9LDAsMSwwLDI0LDIzNzQxNTkzMCwtMSwyMjcxMjY1MjAsMCwxLDAsLTQ5MTI3NTUyMyxJRTVsZEhOallYQmxJRWRsWTJ0dklGZHBiak15SURVdU1DQW9WMmx1Wkc5M2N5a2dNakF4TURBeE1ERWdUVzk2YVd4c1lRPT0sZTMwPSw2NSwtMTI4NTU1MTMsMSwxLC0xLDE2OTk5NTQ4ODcsMTY5OTk1NDg4NywzMzYwMDc5MzMsMTY="
	fmt.Println("target cookie: ", targetCookie)

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		// Print the link found
		fmt.Println("Link found:", e.Attr("href"))
	})

	// On request being made
	c.OnRequest(func(r *colly.Request) {
		//r.Headers.Set("user-agent", )
		r.Headers.Set("cookie", cookie)
		fmt.Println("Visiting", r.URL.String())
	})

	// Start the scraping process on the target URL
	err := c.Visit("https://ozon.ru/")
	if err != nil {
		fmt.Println("Ошибка при попытки скрейпинга: ", err)
	}
}
