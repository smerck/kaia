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

type OpenAIAsker struct{}

func NewOpenAIAsker() *OpenAIAsker {
	return &OpenAIAsker{}
}

type openaiRequest struct {
	Model    string      `json:"model"`
	Messages []oaMessage `json:"messages"`
}

type oaMessage struct {
	Role    string `json:"role"`
	Content string `json:"content"`
}

type openaiResponse struct {
	Choices []struct {
		Message oaMessage `json:"message"`
	} `json:"choices"`
}

func (a *OpenAIAsker) Ask(ctx context.Context, prompt, cluster string) (string, error) {
	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Println("[OpenAI] Missing OPENAI_API_KEY environment variable")
		return "", errors.New("missing OPENAI_API_KEY environment variable")
	}

	fullPrompt := prompt
	if cluster != "" {
		fullPrompt = fmt.Sprintf("[Cluster: %s] %s", cluster, prompt)
	}

	reqBody := openaiRequest{
		Model:    "gpt-3.5-turbo",
		Messages: []oaMessage{{Role: "user", Content: fullPrompt}},
	}
	b, _ := json.Marshal(reqBody)

	log.Printf("[OpenAI] Request body: %s\n", string(b))

	req, err := http.NewRequestWithContext(ctx, "POST", "https://api.openai.com/v1/chat/completions", bytes.NewBuffer(b))
	if err != nil {
		log.Printf("[OpenAI] Error creating request: %v\n", err)
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Printf("[OpenAI] HTTP error: %v\n", err)
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	log.Printf("[OpenAI] Response status: %d, body: %s\n", resp.StatusCode, string(body))

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("OpenAI API error: %s", string(body))
	}

	var oaiResp openaiResponse
	if err := json.Unmarshal(body, &oaiResp); err != nil {
		log.Printf("[OpenAI] JSON decode error: %v\n", err)
		return "", err
	}
	if len(oaiResp.Choices) > 0 {
		return oaiResp.Choices[0].Message.Content, nil
	}
	return "", nil
}
