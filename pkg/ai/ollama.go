package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

type OllamaRequest struct {
	Model  string `json:"model"`
	Prompt string `json:"prompt"`
	Stream bool   `json:"stream"`
}

type OllamaResponse struct {
	Response string `json:"response"`
}

func AskAI(testContent string) ([]int, error) {
	fullPrompt := fmt.Sprintf(`
Analyze these questions and provide ONLY the correct option letter (A, B, C, D, or E) for each.
One letter per line. No numbers, no dots, no explanations.

Questions:
%s`, testContent)
	reqBody, _ := json.Marshal(OllamaRequest{
		Model:  "llama3.1",
		Prompt: fullPrompt,
		Stream: false,
	})

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		return nil, fmt.Errorf("Ollama connection error: %v", err)
	}
	defer resp.Body.Close()

	var result OllamaResponse
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, fmt.Errorf("decode error: %v", err)
	}
	log.Println("--- AI RAW RESPONSE ---")
	fmt.Println(result.Response)
	log.Println("-----------------------")

	var results []int
	lines := strings.Split(result.Response, "\n")

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		found := false
		for _, char := range strings.ToUpper(line) {
			if char >= 'A' && char <= 'E' {
				num := int(char - 'A')
				results = append(results, num)
				found = true
				break
			}
		}
		if !found {
			log.Printf("Warning: No valid letter found in line: %s", line)
		}
	}

	return results, nil
}
