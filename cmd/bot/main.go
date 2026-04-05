package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"AutoGoogleDocs/pkg/browser"
	"AutoGoogleDocs/pkg/parser"
)

func main() {
	file, err := os.OpenFile("bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // log initialization
	if err != nil {
		fmt.Println("Error, couldn't make bot.log file:", err)
		return
	}
	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)
	log.SetFlags(log.Ltime)
	// Main
	url, _, gmail, pass := browser.Autorisation() // get user's information
	//test
	url = "https://docs.google.com/forms/d/e/1FAIpQLSdnyeXX3I_pww1SUTG44Rv3pt38eOvSPKIsGQwqBy8KCDaCBQ/viewform?usp=header"
	//url = "https://docs.google.com/forms/d/e/1FAIpQLSeGGTTCvJfabnBRahbfIGCCfHX1KAbZQgPipwPGRjpszCgHew/viewform"

	page, b := browser.Init(url, false) // go to the test
	defer b.MustClose()

	parser.Login(page, b, gmail, pass)
	time.Sleep(300 * time.Second)
}
