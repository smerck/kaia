package asker

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

type ClaudeAsker struct{}

func NewClaudeAsker() *ClaudeAsker {
	return &ClaudeAsker{}
}

type claudeMessageRequest struct {
	Model     string          `json:"model"`
	Messages  []claudeMessage `json:"messages"`
	MaxTokens int             `json:"maxTokens"`
}

type claudeMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type claudeMessageResponse struct {
	Content []struct {
		Text string `json:"text"`
	} `json:"content"`
}

func (a *ClaudeAsker) Ask(ctx context.Context, prompt, cluster string) (string, error) {
	apiKey := os.Getenv("ANTHROPIC_API_KEY")
	if apiKey == "" {
		log.Println("[Claude] Missing ANTHROPIC_API_KEY environment variable")
		return "", errors.New("missing ANTHROPIC_API_KEY environment variable")
	}

	fullPrompt := prompt
	if cluster != "" {
		fullPrompt = fmt.Sprintf("[Cluster: %s] %s", cluster, prompt)
	}

	reqBody := claudeMessageRequest{
		Model: "claude-sonnet-4-20250514",
		Messages: []claudeMessage{
			{Role: "user", Content: fullPrompt},
		},
		MaxTokens: 512,
	}
	b, _ := json.Marshal(reqBody)

	log.Printf("[Claude] Request body: %s\n", string(b))

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.anthropic.com/v1/messages", bytes.NewBuffer(b))
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

	var cr claudeMessageResponse
	if err := json.Unmarshal(body, &cr); err != nil {
		log.Printf("[Claude] JSON decode error: %v\n", err)
		return "", err
	}

	if len(cr.Content) > 0 {
		return cr.Content[0].Text, nil
	}
	return "", nil
}
