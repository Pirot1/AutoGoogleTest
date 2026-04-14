package main

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
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
	url, line, gmail, pass := browser.Autorisation()
	page, b := browser.Init(url, true)
	defer b.MustClose()

	parser.Login(page, b, gmail, pass)
	time.Sleep(3 * time.Second)
	page.Reload()
	parser.Start_quiz(page, line)
	log.Println("Waiting for timer")
	randomminutes := rand.IntN(5) + 5
	randomsecundes := rand.IntN(61)
	time.Sleep(time.Duration(randomminutes)*time.Minute + time.Duration(randomsecundes)*time.Second)
	log.Println("Finishing up...")
	page.MustElementsX(`//div[@role="button"]`)[0].MustClick()
	log.Println("Finish!")
}
