package api

import (
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestInfoHandlerGet(t *testing.T) {
	rec := httptest.NewRecorder()
	InfoHandler()(rec, httptest.NewRequest(http.MethodGet, "/info", nil))

	res := rec.Result()
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		t.Fatalf("expected status %d, got %d", http.StatusOK, res.StatusCode)
	}
	if ct := res.Header.Get("Content-Type"); ct != "application/json" {
		t.Fatalf("expected Content-Type application/json, got %q", ct)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		t.Fatalf("read body: %v", err)
	}

	var raw map[string]json.RawMessage
	if err := json.Unmarshal(body, &raw); err != nil {
		t.Fatalf("decode body %q: %v", body, err)
	}
	if len(raw) != 3 {
		t.Fatalf("expected exactly 3 keys, got %d in body %q", len(raw), body)
	}
	for _, key := range []string{"service", "version", "status"} {
		if _, ok := raw[key]; !ok {
			t.Fatalf("missing key %q in body %q", key, body)
		}
	}

	var info InfoResponse
	if err := json.Unmarshal(body, &info); err != nil {
		t.Fatalf("decode info body %q: %v", body, err)
	}
	if info.Service != "catalog-go" {
		t.Fatalf("expected service catalog-go, got %q", info.Service)
	}
	if info.Version != "0.1.0" {
		t.Fatalf("expected version 0.1.0, got %q", info.Version)
	}
	if info.Status != "ok" {
		t.Fatalf("expected status ok, got %q", info.Status)
	}
}

func TestInfoHandlerRejectsNonGet(t *testing.T) {
	for _, method := range []string{http.MethodPost, http.MethodPut} {
		rec := httptest.NewRecorder()
		InfoHandler()(rec, httptest.NewRequest(method, "/info", nil))

		if rec.Code != http.StatusMethodNotAllowed {
			t.Fatalf("%s /info: expected status %d, got %d", method, http.StatusMethodNotAllowed, rec.Code)
		}
		if allow := rec.Header().Get("Allow"); allow != http.MethodGet {
			t.Fatalf("%s /info: expected Allow %q, got %q", method, http.MethodGet, allow)
		}

		var errBody map[string]string
		if err := json.Unmarshal(rec.Body.Bytes(), &errBody); err != nil {
			t.Fatalf("%s /info: decode body %q: %v", method, rec.Body.String(), err)
		}
		if errBody["service"] != "" || errBody["status"] != "" {
			t.Fatalf("%s /info: expected an error body, got %q", method, rec.Body.String())
		}
	}
}

func TestNewRouterRouting(t *testing.T) {
	router := NewRouter()

	info := httptest.NewRecorder()
	router.ServeHTTP(info, httptest.NewRequest(http.MethodGet, "/info", nil))
	if info.Code != http.StatusOK {
		t.Fatalf("GET /info: expected status %d, got %d", http.StatusOK, info.Code)
	}

	post := httptest.NewRecorder()
	router.ServeHTTP(post, httptest.NewRequest(http.MethodPost, "/info", nil))
	if post.Code != http.StatusMethodNotAllowed {
		t.Fatalf("POST /info: expected status %d, got %d", http.StatusMethodNotAllowed, post.Code)
	}

	missing := httptest.NewRecorder()
	router.ServeHTTP(missing, httptest.NewRequest(http.MethodGet, "/health", nil))
	if missing.Code != http.StatusNotFound {
		t.Fatalf("GET /health: expected status %d, got %d", http.StatusNotFound, missing.Code)
	}

	var errBody map[string]string
	if err := json.Unmarshal(missing.Body.Bytes(), &errBody); err != nil {
		t.Fatalf("GET /health: decode body %q: %v", missing.Body.String(), err)
	}
	if errBody["error"] != "not found" {
		t.Fatalf("GET /health: expected error not found, got %q", missing.Body.String())
	}
}
