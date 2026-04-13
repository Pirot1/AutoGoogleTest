package parser

import (
	"fmt"
	"log"

	"github.com/go-rod/rod"
)

func Start_quiz(page *rod.Page, line string) (data string) {
	log.Println("Start solving quiz...")
	page.MustElementX(`//input[@type="text"][1]`).MustInput(line)
	els := page.MustElementsX(`//div[@role="listitem"]`)
	data = ""
	for _, el := range els {
		text := el.MustText()
		data = fmt.Sprintf("%s\n%s", data, text)
	}
	return data
}

func Solve_quiz(page *rod.Page, answer []int) {
	elements := page.MustElementsX(`//div[@class="Qr7Oae"]`)

	for i := 1; i <= len(elements); i++ {
		answers := elements[i].MustElementsX(`//div[@class="nWQGrd zwllIb"]`)
		for id, ans := range answers {
			if id == answer[i-1] {
				ans.MustElementX(`//div[@class="d7L4fc bJNwt  FXLARc aomaEc ECvBRb"]`).MustClick()
				break
			}
		}
	}
	log.Println("Test is done!")
}
