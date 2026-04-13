package parser

import (
	"log"
	"time"

	"AutoGoogleDocs/pkg/ai"

	"github.com/go-rod/rod"
)

func Start_quiz(page *rod.Page, line string) {
	log.Println("Start solving quiz...")
	time.Sleep(500 * time.Millisecond)
	page.MustElementX(`//input[@type="text"]`).MustClick().MustInput(line)
	step := 5
	elements := page.MustElementsX(`//div[@role="listitem"]`)
	var totalAns []int
	for i := 0; i < len(elements); i += step {
		end := i + step
		if end > len(elements) {
			end = len(elements)
		}
		chunk := elements[i:end]
		var questionBlock string
		for _, el := range chunk {
			questionBlock += el.MustText() + "\n"
		}
		answer, err := ai.AskAI(questionBlock)
		if err != nil {
			log.Println("Error with loading AI:", err)
			continue
		}
		for j := 0; j < len(answer); j++ {
			totalAns = append(totalAns, answer[j])
		}
	}
	Solve_quiz(page, totalAns)
}

func Solve_quiz(page *rod.Page, answer []int) {
	questions := page.MustElementsX(`//div[@role="listitem"]`)
	for i, questionEl := range questions {
		if i >= len(answer) {
			log.Printf("Caution: Answers is less than questions. Stop on question %d", i+1)
			break
		}
		options := questionEl.MustElementsX(`.//div[@class="nWQGrd zwllIb"]`)
		targetIndex := answer[i]
		if targetIndex >= 0 && targetIndex < len(options) {
			btn := options[targetIndex].MustElementX(`.//div[@class="d7L4fc bJNwt FXLARc aomaEc ECvBRb"]`)
			btn.MustClick()
			log.Printf("Question %d: the answer is %d", i+1, targetIndex)
		} else {
			log.Printf("ERROR: index %d is out of range (with length %d)", targetIndex, len(options))
		}
	}
	log.Println("Test is done!")
}
