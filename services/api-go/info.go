package api

import (
	"encoding/json"
	"net/http"
)

// ServiceName is the logical service name reported by the /info endpoint.
const ServiceName = "catalog-go"

// ServiceVersion is the service version reported by the /info endpoint.
const ServiceVersion = "0.1.0"

// InfoResponse is the JSON payload returned by GET /info.
type InfoResponse struct {
	Service string `json:"service"`
	Version string `json:"version"`
	Status  string `json:"status"`
}

// Info returns the current service information payload.
func Info() InfoResponse {
	return InfoResponse{
		Service: ServiceName,
		Version: ServiceVersion,
		Status:  "ok",
	}
}

// InfoHandler answers GET /info and rejects every other method with 405.
//
// The method is checked explicitly because go.mod declares go 1.21, where the
// method-aware "GET /info" ServeMux patterns are not available.
func InfoHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodGet {
			w.Header().Set("Allow", http.MethodGet)
			writeJSON(w, http.StatusMethodNotAllowed, map[string]string{"error": "method not allowed"})
			return
		}
		writeJSON(w, http.StatusOK, Info())
	}
}

// NewRouter returns the HTTP router for the service.
func NewRouter() http.Handler {
	mux := http.NewServeMux()
	mux.Handle("/info", InfoHandler())
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		writeJSON(w, http.StatusNotFound, map[string]string{"error": "not found"})
	})
	return mux
}

func writeJSON(w http.ResponseWriter, status int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
