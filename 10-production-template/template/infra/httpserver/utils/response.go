package utils

import (
	"encoding/json"
	"net/http"
)

type response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

func writeJSON(w http.ResponseWriter, status int, body response) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}

func WriteSuccess(w http.ResponseWriter, status int, message string, data any) {
	writeJSON(w, status, response{
		Status:  true,
		Message: message,
		Data:    data,
	})
}

func WriteError(w http.ResponseWriter, status int, message string, optionalError ...error) {
	var data any
	if len(optionalError) > 0 && optionalError[0] != nil {
		data = optionalError[0].Error()
	}

	writeJSON(w, status, response{
		Status:  false,
		Message: message,
		Data:    data,
	})
}
