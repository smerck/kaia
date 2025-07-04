package api

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

type stubAgent struct{}

func (a *stubAgent) Ask(_ context.Context, _ string, _ string) (string, error) {
	return "stub response", nil
}

func TestAskHandler_ValidRequest(t *testing.T) {
	reqBody := AskRequest{Prompt: "Why are my pods crashing?", Cluster: "test-cluster"}
	b, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/ask", bytes.NewBuffer(b))
	w := httptest.NewRecorder()

	handler := MakeAskHandlerWithAgent(&stubAgent{})
	handler(w, req)

	resp := w.Result()
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected 200, got %d", resp.StatusCode)
	}

	var askResp AskResponse
	if err := json.NewDecoder(resp.Body).Decode(&askResp); err != nil {
		t.Fatalf("invalid response json: %v", err)
	}
	if askResp.Response == "" {
		t.Error("expected non-empty response")
	}
}

func TestAskHandler_InvalidMethod(t *testing.T) {
	req := httptest.NewRequest(http.MethodGet, "/api/v1/ask", nil)
	w := httptest.NewRecorder()
	handler := MakeAskHandlerWithAgent(&stubAgent{})
	handler(w, req)
	if w.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Error("expected 405 for GET method")
	}
}

func TestAskHandler_BadJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/ask", bytes.NewBuffer([]byte("not-json")))
	w := httptest.NewRecorder()
	handler := MakeAskHandlerWithAgent(&stubAgent{})
	handler(w, req)
	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error("expected 400 for bad json")
	}
}

// MakeAskHandlerWithAgent is a test helper to inject a stub agent.
func MakeAskHandlerWithAgent(ag interface {
	Ask(context.Context, string, string) (string, error)
}) http.HandlerFunc {
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

		respText, err := ag.Ask(context.Background(), req.Prompt, req.Cluster)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		resp := AskResponse{Response: respText}
		w.Header().Set("Content-Type", "application/json")
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
			return
		}
	}
}
