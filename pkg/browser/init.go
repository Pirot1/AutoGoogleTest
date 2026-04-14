package browser

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
)

func Autorisation() (url string, name string, gmail string, pass string) {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input url:")
	url, _ = reader.ReadString('\n')
	fmt.Println("Input your (NAME SURNAME, GROUP):")
	name, _ = reader.ReadString('\n')
	fmt.Println("Input your Gmail:")
	gmail, _ = reader.ReadString('\n')
	fmt.Println("Input your password:")
	pass, _ = reader.ReadString('\n')
	return url, name, gmail, pass
}

func Init(site string, head bool) (page *rod.Page, browser *rod.Browser) {
	chromePath := `C:\Program Files\Google\Chrome\Application\chrome.exe`
	l := launcher.New().Bin(chromePath).Headless(head).Devtools(false).Leakless(false)
	l.Set("autoplay-policy", "no-user-gesture-required")
	l.Set("disable-gpu")
	l.Set("no-sandbox")
	l.Set("disable-dev-shm-usage")
	l.Set("disable-extensions")
	url, err := l.Launch()
	if err != nil {
		log.Panicf("Couldn't init browser: %v", err)
	}
	browser = rod.New().ControlURL(url).MustConnect().NoDefaultDevice()
	page = stealth.MustPage(browser)
	page.MustSetViewport(1920, 1080, 1, false)
	page.MustSetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	})
	page = page.MustNavigate(site)
	log.Println("Booting up...")
	return page, browser
}
