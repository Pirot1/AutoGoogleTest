package parser

import (
	"AutoGoogleDocs/pkg/ai"
	"fmt"
	"log"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/input"
)

func Start_quiz(page *rod.Page, line string) {
	log.Println("Start solving quiz...")
	time.Sleep(500 * time.Millisecond)
	el := page.MustElementX(`//input[@type="text"]`)
	el.MustClick()
	el.MustSelectAllText()
	el.MustType(input.Backspace)
	el.MustInput(line)

	allElements := page.MustElementsX(`//div[@role="listitem"]`)
	var validQuestions []*rod.Element

	for _, el := range allElements {
		options, _ := el.ElementsX(`.//div[@role="radio"] | .//div[@role="checkbox"]`)
		if len(options) > 0 {
			validQuestions = append(validQuestions, el)
		} else {
			log.Println("Skipping block without answer options (header/intro)")
		}
	}
	step := 5
	for i := 0; i < len(validQuestions); i += step {
		end := i + step
		if end > len(validQuestions) {
			end = len(validQuestions)
		}

		chunk := validQuestions[i:end]
		var questionBlock string
		for j, el := range chunk {
			questionBlock += fmt.Sprintf("Question %d: %s\n", j+1, el.MustText())
		}
		answers, err := ai.AskAI(questionBlock)
		if err != nil {
			log.Println("AI Connection Error:", err)
			continue
		}

		log.Printf("Processing block. Answers received: %d out of %d", len(answers), len(chunk))
		for j, el := range chunk {
			if j < len(answers) {
				options := el.MustElementsX(`.//div[@role="radio"] | .//div[@role="checkbox"]`)
				targetIndex := answers[j]

				if targetIndex >= 0 && targetIndex < len(options) {
					options[targetIndex].MustClick()
					log.Printf("  - Q%d: Clicked option %d", i+j+1, targetIndex)
				}
			}
		}
		time.Sleep(800 * time.Millisecond)
	}
	log.Println("All question blocks processed successfully!")
}
