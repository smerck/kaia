package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/smerck/kaia/api"
)

type Config struct {
	AskerBackend string `json:"asker_backend"`
}

func loadConfig() Config {
	f, err := os.Open("config.json")
	if err != nil {
		log.Println("[Config] config.json not found, defaulting to openai backend")
		return Config{AskerBackend: "openai"}
	}
	defer f.Close()
	var cfg Config
	if err := json.NewDecoder(f).Decode(&cfg); err != nil {
		log.Println("[Config] Error decoding config.json, defaulting to openai backend")
		return Config{AskerBackend: "openai"}
	}
	if cfg.AskerBackend == "" {
		cfg.AskerBackend = "openai"
	}
	return cfg
}

func main() {
	cfg := loadConfig()

	http.HandleFunc("/api/v1/ask", api.MakeAskHandler(cfg.AskerBackend))

	log.Println("kaia API server running on :8080 (backend:", cfg.AskerBackend, ")")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("server failed: %v", err)
	}
}
