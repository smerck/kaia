package claude

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ClaudeRequest struct {
	Model             string   `json:"model"`
	MaxTokensToSample int      `json:"maxTokensToSample"`
	Prompt            string   `json:"prompt"`
	StopSequences     []string `json:"stopSequences"`
}

type ClaudeResponse struct {
	Completion string `json:"completion"`
}

func AskClaude(prompt, cluster string) (string, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Println("[Claude] Missing ANTHROPIC_API_KEY environment variable")
		return "", errors.New("missing ANTHROPIC_API_KEY environment variable")
	}

	fullPrompt := prompt
	if cluster != "" {
		fullPrompt = fmt.Sprintf("[Cluster: %s] %s", cluster, prompt)
	}

	reqBody := ClaudeRequest{
		Model:             "claude-2",
		Prompt:            fullPrompt,
		MaxTokensToSample: 512,
	}
	b, _ := json.Marshal(reqBody)

	log.Printf("[Claude] Request body: %s\n", string(b))

	req, err := http.NewRequest("POST", "https://api.anthropic.com/v1/complete", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("[Claude] Error creating request: %v\n", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)
	req.Header.Set("anthropic-version", "2023-06-01")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[Claude] HTTP error: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Printf("[Claude] Response status: %d, body: %s\n", resp.StatusCode, string(body))

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("Claude API error: %s", string(body))
	}

	var cr ClaudeResponse
	if err := json.Unmarshal(body, &cr); err != nil {
		log.Printf("[Claude] JSON decode error: %v\n", err)
		return "", err
	}

	return cr.Completion, nil
}
