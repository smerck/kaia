package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAskHandler_ValidRequest(t *testing.T) {
	reqBody := AskRequest{Prompt: "Why are my pods crashing?", Cluster: "test-cluster"}
	b, _ := json.Marshal(reqBody)
	req := httptest.NewRequest(http.MethodPost, "/api/v1/ask", bytes.NewBuffer(b))
	w := httptest.NewRecorder()

	AskHandler(w, req)

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
	AskHandler(w, req)
	if w.Result().StatusCode != http.StatusMethodNotAllowed {
		t.Error("expected 405 for GET method")
	}
}

func TestAskHandler_BadJSON(t *testing.T) {
	req := httptest.NewRequest(http.MethodPost, "/api/v1/ask", bytes.NewBuffer([]byte("not-json")))
	w := httptest.NewRecorder()
	AskHandler(w, req)
	if w.Result().StatusCode != http.StatusBadRequest {
		t.Error("expected 400 for bad json")
	}
}
