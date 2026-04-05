package parser

import (
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func Login(page *rod.Page, b *rod.Browser, gmail string, pass string) {
	gmail = "narbekd1@gmail.com"
	pass = "A0787040398a!"

	el := page.MustElementX(`//div[@class="kesdnc"]/a[1]`)
	page = b.MustPage(*el.MustAttribute("href"))
	previousURL := page.MustInfo().URL
	page.MustElementX(`//input[@type="email"]`).MustInput(gmail)
	page.KeyActions().Press(input.Enter).MustDo()
	log.Println("Enter gmail successfully")
	url := page.MustInfo().URL
	for url == previousURL {
		time.Sleep(1 * time.Second)
		url = page.MustInfo().URL
	}
	page = b.MustPage(url)
	page.MustElementX(`//input[@type="password"]`).MustInput(pass)
	page.KeyActions().Press(input.Enter).MustDo()
	log.Println("Enter password successfully")
	page.MustWaitNavigation()
	page.MustClose()
}
