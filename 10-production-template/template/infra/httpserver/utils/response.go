package utils

import (
	"encoding/json"
	"errors"
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

func WriteValidationError(w http.ResponseWriter, err error) {
	var verr *ValidationError
	if !errors.As(err, &verr) {
		WriteError(w, http.StatusBadRequest, "invalid request")
		return
	}

	switch verr.Kind {
	case "json":
		WriteError(w, http.StatusBadRequest, "invalid json body")

	case "validation":
		writeJSON(w, http.StatusBadRequest, response{
			Status:  false,
			Message: "validation failed",
			Data: map[string]any{
				"fields": verr.Fields,
			},
		})

	default:
		WriteError(w, http.StatusBadRequest, "invalid request")
	}
}
