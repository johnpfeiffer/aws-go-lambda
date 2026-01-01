package main

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestHandler(t *testing.T) {
	var ctx context.Context
	response, err := HandleRequest(ctx, MyRequest{Value: "john"})
	if err != nil {
		t.Fatalf("error: %s", err)
	}
	expected := "hi john"
	if response.Message != expected {
		t.Errorf("Incorrect response, expected: '%s' , but received: %s", expected, response.Message)
	}
}

func TestGenericHandler_HTTP(t *testing.T) {
	body := `{"value":"john"}`
	req := httptest.NewRequest(http.MethodPost, "http://example.com/", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	rr := httptest.NewRecorder()

	GenericHandler(rr, req)

	if rr.Code != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, rr.Code)
	}
	if ct := rr.Header().Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", ct)
	}

	var resp MyResponse
	if err := json.NewDecoder(rr.Body).Decode(&resp); err != nil {
		t.Fatalf("decode response: %v", err)
	}
	if resp.Message != "hi john" {
		t.Fatalf("unexpected message: %q", resp.Message)
	}
	if resp.Created == "" {
		t.Fatal("expected Created to be set")
	}
}
