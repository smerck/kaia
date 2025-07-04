package api

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/smerck/kaia/internal/agent"
)

type AskRequest struct {
	Prompt  string `json:"prompt"`
	Cluster string `json:"cluster,omitempty"`
	Backend string `json:"backend,omitempty"`
}

type AskResponse struct {
	Response string `json:"response"`
}

func MakeAskHandler(defaultBackend string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			w.WriteHeader(http.StatusMethodNotAllowed)
			return
		}

		var req AskRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		// Use the requested backend, or default if not specified
		backend := req.Backend
		if backend == "" {
			backend = defaultBackend
		}

		ag := agent.New(backend)
		respText, err := ag.Ask(context.Background(), req.Prompt, req.Cluster)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := AskResponse{Response: respText}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			log.Printf("Error encoding response: %v", err)
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
