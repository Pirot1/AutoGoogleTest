package parser

import (
	"log"
	"strings"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func Login(page *rod.Page, b *rod.Browser, gmail string, pass string) {
	// Start loging in
	el := page.MustElementX(`//div[@class="kesdnc"]/a[1]`)
	page = b.MustPage(*el.MustAttribute("href"))
	// Gmail
	emailInput := page.MustElementX(`//input[@type="email"]`).MustWaitVisible()
	emailInput.MustInput(gmail)
	page.KeyActions().Press(input.Enter).MustDo()
	log.Println("Enter gmail successfully")
	// Password
	url := page.MustInfo().URL
	for !strings.Contains(url, "challenge") {
		url = page.MustInfo().URL
	}
	page.MustWaitStable()
	for _, char := range pass {
		page.KeyActions().Type(input.Key(char)).MustDo()
	}
	page.KeyActions().Press(input.Enter).MustDo()
	log.Println("Enter password successfully")
	page.MustWaitNavigation()
	time.Sleep(3 * time.Second)
	page.MustClose()
}
