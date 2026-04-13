package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
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
	fullPrompt := fmt.Sprintf("Question: %s\nProvide only the answer index or text. No explanations.", testContent)

	reqBody, _ := json.Marshal(OllamaRequest{
		Model:  "llama3.1",
		Prompt: fullPrompt,
		Stream: false,
	})

	resp, err := http.Post("http://localhost:11434/api/generate", "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Fatalf("ERROR %v", err)
	}
	defer resp.Body.Close()

	var result OllamaResponse
	json.NewDecoder(resp.Body).Decode(&result)
	var results []int
	var intres int
	for _, res := range strings.Split(result.Response, ",") {
		intres, _ = strconv.Atoi(res)
		results = append(results, intres)
	}
	return results, nil
}
