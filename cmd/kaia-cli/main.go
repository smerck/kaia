package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"io"
	"log"
	"net/http"
	"os"
)

type AskRequest struct {
	Prompt  string `json:"prompt"`
	Cluster string `json:"cluster,omitempty"`
}

type AskResponse struct {
	Response string `json:"response"`
}

func main() {
	prompt := flag.String("prompt", "", "Prompt to send to kaia API")
	cluster := flag.String("cluster", "", "EKS cluster name (optional)")
	apiURL := flag.String("api", "http://localhost:8080/api/v1/ask", "kaia API URL")
	flag.Parse()

	if *prompt == "" {
		log.Println("--prompt is required")
		os.Exit(1)
	}

	reqBody, _ := json.Marshal(AskRequest{Prompt: *prompt, Cluster: *cluster})
	resp, err := http.Post(*apiURL, "application/json", bytes.NewBuffer(reqBody))
	if err != nil {
		log.Printf("Request failed: %v", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != 200 {
		log.Printf("API error: %s", string(body))
		os.Exit(1)
	}

	var askResp AskResponse
	if err := json.Unmarshal(body, &askResp); err != nil {
		log.Printf("Invalid response: %v", err)
		os.Exit(1)
	}
	log.Println(askResp.Response)
}
