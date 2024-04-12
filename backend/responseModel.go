package main

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
)

type RFC7807Error struct {
	Type     string `json:"type"`
	Title    string `json:"title"`
	Status   int    `json:"status"`
	Detail   string `json:"detail"`
	Instance string `json:"instance,omitempty"`
}

func (e *RFC7807Error) Error() string {
	return e.Title
}

// ErrorResponse writes the error response conforming to RFC 7807.
func ErrorResponse(w http.ResponseWriter, err error, status int) {
	w.Header().Set("Content-Type", "application/problem+json")
	w.WriteHeader(status)

	// Serialize error response to JSON
	json.NewEncoder(w).Encode(&RFC7807Error{
		Type:   "about:blank",
		Title:  http.StatusText(status),
		Status: status,
		Detail: err.Error(),
	})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Error marshalling JSON: %s", err)
		w.WriteHeader(500)
		return
	}
	w.WriteHeader(code)
	w.Write(dat)
}

func permissionDenied(w http.ResponseWriter) {
	ErrorResponse(w, errors.New("permission denied"), http.StatusForbidden)
}
