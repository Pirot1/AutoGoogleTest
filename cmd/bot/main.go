package main

import (
	"fmt"
	"io"
	"log"
	"math/rand/v2"
	"os"
	"time"

	"AutoGoogleDocs/pkg/ai"
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
	url, line, gmail, pass := browser.Autorisation() // get user's information
	//test
	url = "https://docs.google.com/forms/d/e/1FAIpQLSdnyeXX3I_pww1SUTG44Rv3pt38eOvSPKIsGQwqBy8KCDaCBQ/viewform?usp=header"
	//url = "https://docs.google.com/forms/d/e/1FAIpQLSeGGTTCvJfabnBRahbfIGCCfHX1KAbZQgPipwPGRjpszCgHew/viewform"

	page, b := browser.Init(url, false) // go to the test
	defer b.MustClose()

	parser.Login(page, b, gmail, pass)
	page.Reload()
	test := parser.Start_quiz(page, line)
	answers, err := ai.AskAI(test)
	parser.Solve_quiz(page, answers)
	log.Println("Waiting for timer")
	randomminutes := rand.IntN(5) + 5
	randomsecundes := rand.IntN(61)
	time.Sleep(time.Duration(randomminutes)*time.Minute + time.Duration(randomsecundes)*time.Second)
	log.Println("Finishing up...")
}
